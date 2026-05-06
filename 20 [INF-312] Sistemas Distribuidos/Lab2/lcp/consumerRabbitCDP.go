package main

import (
	"encoding/json"
	"fmt"
	"l2/types"
	"time"

	"github.com/streadway/amqp"
)

func escucharCDPResultQueue(chConsum *amqp.Channel, chProd *amqp.Channel, listaEntrenadores *[]types.Entrenador, listaTorneos *[]types.Torneo, listaPenalizaciones *[]types.Penalizaciones) {
	fmt.Println("[escucharCDPResultQueue] Escuchando la cola CDPResultQueue...")

	msgs, err := chConsum.Consume(
		"CDPResultQueue", // queue
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		fmt.Printf("    [ ! ] Error al consumir mensajes de la cola CDPResultQueue: %v\n", err)
		return
	}

	for msg := range msgs {
		fmt.Printf("[Cola: CDPResultQueue] Mensaje recibido de CDPResultQueue: %s\n", msg.Body)
		var resultado types.ResultadoCombate
		if err := json.Unmarshal(msg.Body, &resultado); err != nil {
			fmt.Printf("    [ ! ] Error deserializando mensaje: %v\n", err)
			continue
		}
		fmt.Printf("    [ OK ] ResultadoCombate deserializado: %+v\n", resultado)

		// Actualizar torneo
		finalizarTorneoPorID(listaTorneos, resultado.IdTorneo)

		if resultado.Duracion <= 5 {
			err := penalizarEntrenador(resultado.IdGanador, listaPenalizaciones, listaEntrenadores, chProd)
			if err != nil {
				fmt.Printf("    [ ! ] Error al aplicar penalización al entrenador %s: %v\n", resultado.IdGanador, err)
			} else {
				fmt.Printf("    [ OK ] Penalización aplicada al entrenador %s por tiempo de combate insuficiente\n", resultado.IdGanador)
			}
			continue
		}

		// Actualizar ranking del entrenador
		var nombreGanador string
		var nombrePerdedor string
		idGanador := resultado.IdGanador
		nombreGanador = resultado.NombreGanador
		var idPerdedor string
		if resultado.IdGanador == resultado.IdEntrenador1 {
			idPerdedor = resultado.IdEntrenador2
			nombrePerdedor = resultado.NombreEntrenador2
		} else {
			idPerdedor = resultado.IdEntrenador1
			nombrePerdedor = resultado.NombreEntrenador1
		}
		nuevoRankGanador, _ := cambiarRankingEntrenadorPorID(listaEntrenadores, idGanador, 30)
		nuevoRankPerdedor, _ := cambiarRankingEntrenadorPorID(listaEntrenadores, idPerdedor, -10)

		// Enviar mensajes a la cola de resultados

		// Enviar registro de combate para el ganador
		registroGanador := types.RegistroCombate{
			IdEntrenador:   idGanador,
			Nombre:         nombreGanador,
			IdTorneo:       resultado.IdTorneo,
			Resultado:      "victoria",
			RankingAntes:   nuevoRankGanador - 30,
			RankingDespues: nuevoRankGanador,
		}
		registroGanadorJSON, err := json.Marshal(registroGanador)
		if err == nil {
			errEnvio := SendLCPEventsQueue(chProd, "registro_combate", string(registroGanadorJSON))
			if errEnvio != nil {
				fmt.Printf("    [ ! ] Error al enviar registro de combate del ganador: %v\n", errEnvio)
			}
		}

		// Enviar registro de combate para el perdedor
		registroPerdedor := types.RegistroCombate{
			IdEntrenador:   idPerdedor,
			Nombre:         nombrePerdedor,
			IdTorneo:       resultado.IdTorneo,
			Resultado:      "derrota",
			RankingAntes:   nuevoRankPerdedor + 10,
			RankingDespues: nuevoRankPerdedor,
		}
		registroPerdedorJSON, err := json.Marshal(registroPerdedor)
		if err == nil {
			errEnvio := SendLCPEventsQueue(chProd, "registro_combate", string(registroPerdedorJSON))
			if errEnvio != nil {
				fmt.Printf("    [ ! ] Error al enviar registro de combate del perdedor: %v\n", errEnvio)
			}
		}

		// Actualizacion de ranking para el ganador
		mensaje_ranking_ganador := types.RankingActualizado{
			IdEntrenador:     idGanador,
			NombreEntrenador: nombreGanador,
			NuevoRanking:     nuevoRankGanador,
			Fecha:            resultado.Fecha,
			TipoMensaje:      "actualizacion_ranking",
		}
		mensaje_ranking_ganador_json, err := json.Marshal(mensaje_ranking_ganador)
		if err != nil {
			fmt.Printf("    [ ! ] Error al serializar mensaje de ranking del ganador: %v\n", err)
			continue
		}
		err = SendLCPEventsQueue(chProd, "actualizacion_ranking", string(mensaje_ranking_ganador_json))
		if err != nil {
			fmt.Printf("    [ ! ] Error al enviar mensaje de ranking actualizado del ganador: %v\n", err)
		}
		fmt.Printf("    [ OK ] Mensaje de ranking actualizado del ganador enviado a la cola de resultados: %s\n", mensaje_ranking_ganador_json)

		// Actualizacion de ranking para el perdedor
		mensaje_ranking_perdedor := types.RankingActualizado{
			IdEntrenador:     idPerdedor,
			NombreEntrenador: nombrePerdedor,
			NuevoRanking:     nuevoRankPerdedor,
			Fecha:            resultado.Fecha,
			TipoMensaje:      "actualizacion_ranking",
		}
		mensaje_ranking_perdedor_json, err := json.Marshal(mensaje_ranking_perdedor)
		if err != nil {
			fmt.Printf("    [ ! ] Error al serializar mensaje de ranking del perdedor: %v\n", err)
			continue
		}
		err = SendLCPEventsQueue(chProd, "actualizacion_ranking", string(mensaje_ranking_perdedor_json))
		if err != nil {
			fmt.Printf("    [ ! ] Error al enviar mensaje de ranking actualizado del perdedor: %v\n", err)
		}
		fmt.Printf("    [ OK ] Mensaje de ranking actualizado del perdedor enviado a la cola de resultados: %s\n", mensaje_ranking_perdedor_json)
	}
}

func cambiarRankingEntrenadorPorID(listaEntrenadores *[]types.Entrenador, id string, cambioRanking int32) (int32, error) {
	for i, entrenador := range *listaEntrenadores {
		if entrenador.Id == id {
			nuevoRanking := (*listaEntrenadores)[i].Ranking + cambioRanking
			if nuevoRanking < 0 {
				nuevoRanking = 0
			}
			(*listaEntrenadores)[i].Ranking = nuevoRanking
			fmt.Printf("    [ OK ] Ranking del entrenador con ID %s actualizado a %d\n", id, nuevoRanking)
			return nuevoRanking, nil
		}
	}
	return 0, fmt.Errorf("entrenador con ID %s no encontrado", id)
}

func finalizarTorneoPorID(listaTorneos *[]types.Torneo, id string) error {
	for i, torneo := range *listaTorneos {
		if torneo.Id == id {
			(*listaTorneos)[i].Estado = types.TORNEO_FINALIZADO
			fmt.Printf("    [ OK ] Torneo con ID %s finalizado correctamente\n", id)
			return nil
		}
	}
	return fmt.Errorf("torneo con ID %s no encontrado", id)
}

func penalizarEntrenador(id string, listaPenalizaciones *[]types.Penalizaciones, listaEntrenadores *[]types.Entrenador, ch *amqp.Channel) error {
	var index_ent = -1
	var index_pen = -1

	for i, entrenador := range *listaEntrenadores {
		if entrenador.Id == id {
			index_ent = i
			break
		}
	}

	for i, penalizacion := range *listaPenalizaciones {
		if penalizacion.IdEntrenador == id {
			index_pen = i
		}
	}

	if index_pen != -1 {
		if (*listaPenalizaciones)[index_pen].Penalizacion == 3 {
			(*listaEntrenadores)[index_ent].Estado = types.ENTRENADOR_INACTIVO
			(*listaEntrenadores)[index_ent].Suspencion = -1 // SUSPENCION PERMANENTEEEE
			(*listaPenalizaciones)[index_pen].Penalizacion = 4
			penalizacion := types.PenalizacionEntrenador{
				IdEntrenador:     id,
				NombreEntrenador: (*listaEntrenadores)[index_ent].Nombre,
				Penalizacion:     -1,
				Fecha:            time.Now().Format("2006-01-02"),
				TipoMensaje:      "actualizacion_penalizacion",
			}

			penalizacionJSON, err := json.Marshal(penalizacion)
			if err != nil {
				return fmt.Errorf("error al serializar penalización: %v", err)
			}

			err = SendLCPEventsQueue(ch, "actualizacion_penalizacion", string(penalizacionJSON))
			if err != nil {
				return fmt.Errorf("error al enviar penalización: %v", err)
			}
			fmt.Printf("    [ OK ] Entrenador con ID %s penalizado permanentemente\n", id)
			return nil

		} else {
			(*listaPenalizaciones)[index_pen].Penalizacion++
			(*listaEntrenadores)[index_ent].Estado = types.ENTRENADOR_SUSPENDIDO
			(*listaEntrenadores)[index_ent].Suspencion = 3
			penalizacion := types.PenalizacionEntrenador{
				IdEntrenador:     id,
				NombreEntrenador: (*listaEntrenadores)[index_ent].Nombre,
				Penalizacion:     3,
				Fecha:            time.Now().Format("2006-01-02"),
				TipoMensaje:      "actualizacion_penalizacion",
			}
			penalizacionJSON, err := json.Marshal(penalizacion)
			if err != nil {
				return fmt.Errorf("error al serializar penalización: %v", err)
			}
			err = SendLCPEventsQueue(ch, "actualizacion_penalizacion", string(penalizacionJSON))
			if err != nil {
				return fmt.Errorf("error al enviar penalización: %v", err)
			}
			fmt.Printf("    [ OK ] Entrenador con ID %s penalizado, nueva penalización: %d\n", id, (*listaPenalizaciones)[index_pen].Penalizacion)
			return nil
		}
	}

	return fmt.Errorf("no se encontró penalización para el entrenador con ID %s", id)
}
