package main

import (
	"cdp/PokemonProto"
	"encoding/json"
	"fmt"
	"l2/types"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func logWithTimestamp(format string, args ...interface{}) {
	now := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf(format, args...)
	logMsg := fmt.Sprintf("[%s] %s\n", now, msg)
	fmt.Print(logMsg)
	f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		defer f.Close()
		f.WriteString(logMsg)
	}
}

// Consume mensajes de la cola GymResultQueue y deserializa los datos
func ConsumeGymResultQueue(ch *amqp.Channel) error {
	logWithTimestamp("[ConsumeGymResultQueue] Esperando mensajes en GymResultQueue...")

	msgs, err := ch.Consume(
		"GymResultQueue", // queue
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		logWithTimestamp("[ ! ] Error al consumir la cola GymResultQueue: %v", err)
		return fmt.Errorf("[ ! ] Error al consumir la cola GymResultQueue: %v", err)
	}

	go func() {
		for d := range msgs {
			logWithTimestamp("[ * ] Mensaje recibido: %s", d.Body)
			var msg types.CombateGymMessage
			if err := json.Unmarshal(d.Body, &msg); err != nil {
				logWithTimestamp("[ ! ] Error deserializando mensaje: %v", err)
				continue
			}
			logWithTimestamp("[ OK ] Mensaje deserializado: %+v", msg)
			// Aquí puedes procesar el mensaje según sea necesario
		}
	}()

	return nil
}

var combatesProcesados []string // IDs de combates ya procesados

// Goroutine para escuchar la cola GymResultQueue, deserializar y descifrar el mensaje
func escucharGymResultQueue(ch *amqp.Channel, client PokemonProto.ServicioLCPClient, regionKeys map[string]string) {
	logWithTimestamp("[escucharGymResultQueue] Escuchando cola GymResultQueue...\n")

	msgs, err := ch.Consume(
		"GymResultQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logWithTimestamp("[ ! ] Error iniciando consumo: %v", err)
		return
	}
	for d := range msgs {
		logWithTimestamp("[Cola: GymResultQueue] Mensaje recibido:\n    %s", d.Body)
		var msg types.CombateGymMessage
		if err := json.Unmarshal(d.Body, &msg); err != nil {
			logWithTimestamp("[ ! ] Error deserializando mensaje: %v", err)
			continue
		}
		logWithTimestamp("[ * ] Mensaje deserializado: %+v", msg)
		// Descifrar el mensaje usando la clave correspondiente a la región
		llave, ok := regionKeys[msg.Region]
		if !ok {
			logWithTimestamp("[ ! ] Región desconocida para descifrado: %s", msg.Region)
			continue
		}
		resultado, err := DesencriptarResultadoCombate(msg.MsgCifrado, llave)
		if err != nil {
			logWithTimestamp("[ ! ] Error descifrando mensaje: %v", err)
			continue
		}

		// Validar estructura de ResultadoCombate
		fechaValida := true
		if _, err := time.Parse("2006-01-02", resultado.Fecha); err != nil {
			fechaValida = false
		}
		if resultado.IdTorneo == "" || resultado.IdEntrenador1 == "" || resultado.IdEntrenador2 == "" || resultado.IdGanador == "" || !fechaValida {
			logWithTimestamp("[ ! ] ResultadoCombate inválido: %+v", resultado)
			fallo := types.FalloCombate{
				IdTorneo:          resultado.IdTorneo,
				IdEntrenador1:     resultado.IdEntrenador1,
				NombreEntrenador1: resultado.NombreEntrenador1,
				IdEntrenador2:     resultado.IdEntrenador2,
				NombreEntrenador2: resultado.NombreEntrenador2,
				Fecha:             resultado.Fecha,
				TipoMensaje:       "fallo_combate",
			}
			_ = SendCDPErrorsQueue(ch, fallo)
			continue
		}

		// Verificar duplicados
		duplicado := false
		for _, id := range combatesProcesados {
			if id == resultado.IdTorneo {
				duplicado = true
				break
			}
		}
		if duplicado {
			logWithTimestamp("[ ! ] Combate duplicado detectado (ID: %s), notificando como fallo", resultado.IdTorneo)
			fallo := types.FalloCombate{
				IdTorneo:          resultado.IdTorneo,
				IdEntrenador1:     resultado.IdEntrenador1,
				NombreEntrenador1: resultado.NombreEntrenador1,
				IdEntrenador2:     resultado.IdEntrenador2,
				NombreEntrenador2: resultado.NombreEntrenador2,
				Fecha:             resultado.Fecha,
				TipoMensaje:       "fallo_combate",
			}
			_ = SendCDPErrorsQueue(ch, fallo)
			continue
		}

		logWithTimestamp("[ OK ] Resultado descifrado: %+v", resultado)

		id_e1 := resultado.IdEntrenador1
		id_e2 := resultado.IdEntrenador2

		existe1, _ := lcpExisteEntrenador(client, id_e1)
		existe2, _ := lcpExisteEntrenador(client, id_e2)
		if !existe1 || !existe2 {
			logWithTimestamp("[ ! ] Uno o ambos entrenadores no existen en LCP: %s, %s", id_e1, id_e2)
			fallo := types.FalloCombate{
				IdTorneo:          resultado.IdTorneo,
				IdEntrenador1:     resultado.IdEntrenador1,
				NombreEntrenador1: resultado.NombreEntrenador1,
				IdEntrenador2:     resultado.IdEntrenador2,
				NombreEntrenador2: resultado.NombreEntrenador2,
				Fecha:             resultado.Fecha,
				TipoMensaje:       "fallo_combate",
			}
			err = SendCDPErrorsQueue(ch, fallo)
			if err != nil {
				logWithTimestamp("[ ! ] Error al enviar error: %v", err)
			} else {
				logWithTimestamp("[ OK ] Error enviado al CDP: %+v", fallo)
			}
			logWithTimestamp("[ * ] Resultado no enviado al CDP porque uno o ambos entrenadores no existen.")
			continue
		}

		entrenador1, _ := lcpObtenerEntrenador(client, id_e1)
		entrenador2, _ := lcpObtenerEntrenador(client, id_e2)

		activos := entrenador1.Estado == types.ENTRENADOR_ACTIVO && entrenador2.Estado == types.ENTRENADOR_ACTIVO
		if !activos {
			logWithTimestamp("[ ! ] Uno o ambos entrenadores no están activos: %s, %s", id_e1, id_e2)
			fallo := types.FalloCombate{
				IdTorneo:          resultado.IdTorneo,
				IdEntrenador1:     resultado.IdEntrenador1,
				NombreEntrenador1: resultado.NombreEntrenador1,
				IdEntrenador2:     resultado.IdEntrenador2,
				NombreEntrenador2: resultado.NombreEntrenador2,
				Fecha:             resultado.Fecha,
				TipoMensaje:       "fallo_combate",
			}
			err = SendCDPErrorsQueue(ch, fallo)
			if err != nil {
				logWithTimestamp("[ ! ] Error al enviar error: %v", err)
			} else {
				logWithTimestamp("[ OK ] Error enviado al CDP: %+v", fallo)
			}
			logWithTimestamp("[ * ] Resultado no enviado al CDP porque uno o ambos entrenadores no están activos.")
		} else {
			// Validar que el ganador sea uno de los entrenadores y el nombre coincida
			ganadorValido := false
			if resultado.IdGanador == entrenador1.Id && resultado.NombreGanador == entrenador1.Nombre {
				ganadorValido = true
			}
			if resultado.IdGanador == entrenador2.Id && resultado.NombreGanador == entrenador2.Nombre {
				ganadorValido = true
			}
			if !ganadorValido {
				logWithTimestamp("[ ! ] El ganador no coincide con los entrenadores recibidos. ID/NOMBRE ganador: %s/%s", resultado.IdGanador, resultado.NombreGanador)
				fallo := types.FalloCombate{
					IdTorneo:          resultado.IdTorneo,
					IdEntrenador1:     resultado.IdEntrenador1,
					NombreEntrenador1: resultado.NombreEntrenador1,
					IdEntrenador2:     resultado.IdEntrenador2,
					NombreEntrenador2: resultado.NombreEntrenador2,
					Fecha:             resultado.Fecha,
					TipoMensaje:       "fallo_combate",
				}
				err = SendCDPErrorsQueue(ch, fallo)
				if err != nil {
					logWithTimestamp("[ ! ] Error al enviar error: %v", err)
				} else {
					logWithTimestamp("[ OK ] Error enviado al CDP: %+v", fallo)
				}
				logWithTimestamp("[ * ] Resultado no enviado al CDP porque el ganador no es válido.")
			} else {
				result := types.ResultadoCombate{
					IdTorneo:          resultado.IdTorneo,
					IdEntrenador1:     resultado.IdEntrenador1,
					NombreEntrenador1: resultado.NombreEntrenador1,
					IdEntrenador2:     resultado.IdEntrenador2,
					NombreEntrenador2: resultado.NombreEntrenador2,
					IdGanador:         resultado.IdGanador,
					NombreGanador:     resultado.NombreGanador,
					Fecha:             resultado.Fecha,
					TipoMensaje:       "resultado_combate",
					Duracion:          resultado.Duracion,
				}
				err = SendCDPResultQueue(ch, result)
				if err != nil {
					logWithTimestamp("[ ! ] Error al enviar resultado: %v", err)
				} else {
					logWithTimestamp("[ OK ] Resultado enviado al CDP: %+v", result)
					combatesProcesados = append(combatesProcesados, result.IdTorneo)
				}
			}
		}
		logWithTimestamp("[LOG] Evento procesado para combate ID: %s", resultado.IdTorneo)
	}
}
