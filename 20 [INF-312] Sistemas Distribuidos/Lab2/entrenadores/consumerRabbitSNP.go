package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"l2/types"

	"github.com/streadway/amqp"
)

func conectarRabbit(rabbit_ip, rabbit_port, rabbit_user, rabbit_pass string) (*amqp.Connection, *amqp.Channel, error) {
	fmt.Println("[ConectarRabbit] Conectando a RabbitMQ...")
	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbit_user, rabbit_pass, rabbit_ip, rabbit_port)

	var conn *amqp.Connection
	var ch *amqp.Channel
	var err error

	for i := 1; i <= 5; i++ {
		fmt.Printf("    [ * ] Intentando conectar a RabbitMQ (intento %d/5)...\n", i)
		conn, err = amqp.Dial(amqpURL)
		if err == nil {
			ch, err = conn.Channel()
			if err == nil {
				fmt.Println("    [ OK ] Conexión a RabbitMQ establecida correctamente")
				return conn, ch, nil
			}
			conn.Close()
		}
		fmt.Printf("    [ ! ] Error conectando a RabbitMQ.")
		if i < 5 {
			fmt.Println(" Esperando 5 segundos antes de reintentar...")
			time.Sleep(5 * time.Second)
		}
	}
	return nil, nil, fmt.Errorf("    [ ! ] Error conectando a RabbitMQ después de 5 intentos: %v", err)
}

func logNotifyWithTimestamp(format string, args ...interface{}) {
	now := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf(format, args...)
	logMsg := fmt.Sprintf("[%s] %s\n", now, msg)
	f, err := os.OpenFile("notify_logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		defer f.Close()
		f.WriteString(logMsg)
	}
}

func escucharSNPNotifyQueue(ch *amqp.Channel, lista_entrenadores *[]types.Entrenador, lista_historial *[]types.RegistroCombate) {
	fmt.Printf("[EscucharSNP] Escuchando cola SNPNotifyQueue...\n")
	msgs, err := ch.Consume(
		"SNPNotifyQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("    [ ! ] Error iniciando consumo: %v\n", err)
		return
	}
	for d := range msgs {
		var notificacion types.NotificacionGenerica
		err := json.Unmarshal(d.Body, &notificacion)
		if err != nil {
			fmt.Printf("[Cola: SNPNotifyQueue] Mensaje recibido:\nError al deserializar mensaje: %v\n    %s\n", err, d.Body)
			continue
		}

		switch notificacion.TipoMensaje {
		case "registro_combate":
			var registro types.RegistroCombate
			err := json.Unmarshal([]byte(notificacion.Payload), &registro)
			if err != nil {
				fmt.Printf("[Cola: SNPNotifyQueue] Mensaje recibido:\nError al deserializar RegistroCombate: %v\n", err)
				continue
			}
			msg := fmt.Sprintf("Entrenador %s tuvo %s en torneo %s", registro.Nombre, registro.Resultado, registro.IdTorneo)
			fmt.Printf("[Cola: SNPNotifyQueue] Mensaje recibido:\n    %s\n", msg)
			logNotifyWithTimestamp("%s|registro_combate - %s", "", msg)
			*lista_historial = append(*lista_historial, registro)
		case "resultado_combate":
			var resultado types.ResultadoCombate
			err := json.Unmarshal([]byte(notificacion.Payload), &resultado)
			if err != nil {
				fmt.Printf("[Cola: SNPNotifyQueue] Mensaje recibido:\nError al deserializar ResultadoCombate: %v\n", err)
				continue
			}
			msg := fmt.Sprintf("Torneo %s finalizado. %s vs %s, gana %s", resultado.IdTorneo, resultado.NombreEntrenador1, resultado.NombreEntrenador2, resultado.NombreGanador)
			fmt.Printf("[Cola: SNPNotifyQueue] Mensaje recibido:\n    %s\n", msg)
			logNotifyWithTimestamp("%s|resultado_combate - %s", resultado.Fecha, msg)
		case "actualizacion_ranking":
			var ranking types.RankingActualizado
			err := json.Unmarshal([]byte(notificacion.Payload), &ranking)
			if err != nil {
				fmt.Printf("[Cola: SNPNotifyQueue] Mensaje recibido:\nError al deserializar RankingActualizado: %v\n", err)
				continue
			}
			actualizarRankingEntrenador(lista_entrenadores, ranking)
			msg := fmt.Sprintf("Entrenador %s ranking actualizado a %d", ranking.NombreEntrenador, ranking.NuevoRanking)
			logNotifyWithTimestamp("%s|actualizacion_ranking - %s", ranking.Fecha, msg)
		case "nuevo_torneo":
			var torneo types.NuevoTorneo
			err := json.Unmarshal([]byte(notificacion.Payload), &torneo)
			if err != nil {
				fmt.Printf("[Cola: SNPNotifyQueue] Mensaje recibido:\nError al deserializar NuevoTorneo: %v\n", err)
				continue
			}
			msg := fmt.Sprintf("Nuevo torneo disponible en %s. ID: %s", torneo.Region, torneo.IdTorneo)
			fmt.Printf("[Cola: SNPNotifyQueue] Mensaje recibido:\n    %s\n", msg)
			logNotifyWithTimestamp("%s|nuevo_torneo - %s", torneo.Fecha, msg)
		case "actualizacion_penalizacion":
			var penalizacion types.PenalizacionEntrenador
			err := json.Unmarshal([]byte(notificacion.Payload), &penalizacion)
			if err != nil {
				fmt.Printf("[Cola: SNPNotifyQueue] Mensaje recibido:\nError al deserializar PenalizacionEntrenador: %v\n", err)
				continue
			}
			actualizarPenalizacionEntrenador(lista_entrenadores, penalizacion)
		}
	}
}

func actualizarRankingEntrenador(lista_entrenadores *[]types.Entrenador, ranking types.RankingActualizado) {
	var rankingAntiguo int32 = -1
	for i := range *lista_entrenadores {
		if (*lista_entrenadores)[i].Id == ranking.IdEntrenador {
			rankingAntiguo = (*lista_entrenadores)[i].Ranking
			(*lista_entrenadores)[i].Ranking = ranking.NuevoRanking
			break
		}
	}
	if rankingAntiguo != -1 {
		fmt.Printf("[Cola: SNPNotifyQueue] Mensaje recibido:\n    Entrenador %s cambia ranking de %d a %d\n", ranking.NombreEntrenador, rankingAntiguo, ranking.NuevoRanking)
	} else {
		fmt.Printf("[Cola: SNPNotifyQueue] Mensaje recibido:\n    Entrenador %s ranking actualizado a %d\n", ranking.NombreEntrenador, ranking.NuevoRanking)
	}
}

func actualizarPenalizacionEntrenador(lista_entrenadores *[]types.Entrenador, penalizacion types.PenalizacionEntrenador) {
	for i := range *lista_entrenadores {
		if (*lista_entrenadores)[i].Id == penalizacion.IdEntrenador {
			if penalizacion.Penalizacion < 0 {
				(*lista_entrenadores)[i].Suspencion = penalizacion.Penalizacion
				(*lista_entrenadores)[i].Estado = types.ENTRENADOR_INACTIVO
				msg := fmt.Sprintf("Entrenador %s penalizado permanentemente. Estado cambiado a inactivo.", penalizacion.NombreEntrenador)
				fmt.Printf("[Cola: SNPNotifyQueue] Mensaje recibido:\n    %s\n", msg)
				logNotifyWithTimestamp("%s|actualizacion_penalizacion - %s", penalizacion.Fecha, msg)
			} else {
				(*lista_entrenadores)[i].Suspencion += penalizacion.Penalizacion
				(*lista_entrenadores)[i].Estado = types.ENTRENADOR_SUSPENDIDO
				msg := fmt.Sprintf("Entrenador %s penalizado con %d turnos. Estado cambiado a suspendido.", penalizacion.NombreEntrenador, penalizacion.Penalizacion)
				fmt.Printf("[Cola: SNPNotifyQueue] Mensaje recibido:\n    %s\n", msg)
				logNotifyWithTimestamp("%s|actualizacion_penalizacion - %s", penalizacion.Fecha, msg)
			}
			break
		}
	}
}
