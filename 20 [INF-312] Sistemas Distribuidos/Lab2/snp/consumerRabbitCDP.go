package main

import (
	"encoding/json"

	"l2/types"

	"github.com/streadway/amqp"
)

// TODO: Añadir función para escuchar la cola 'CDPErrorsQueue'

func escucharCDPErrorsQueue(lis_ch, prod_ch *amqp.Channel) {
	logWithTimestamp("[EscucharCDPResultQueue] Escuchando la cola CDPResultQueue...")

	msgs, err := lis_ch.Consume(
		"CDPErrorsQueue", // queue
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		logWithTimestamp("    [ ! ] Error al consumir mensajes de la cola CDPResultQueue: %v", err)
		return
	}

	for msg := range msgs {
		logWithTimestamp("[Cola: CDPErrorsQueue] Mensaje recibido de CDPErrorsQueue: %s\n", msg.Body)

		var fallo types.FalloCombate
		err := json.Unmarshal(msg.Body, &fallo)
		if err != nil {
			logWithTimestamp("    [ ! ] Error deserializando mensaje FalloCombate: %v", err)
			continue
		}

		payloadBytes, err := json.Marshal(fallo)
		if err != nil {
			logWithTimestamp("    [ ! ] Error serializando payload FalloCombate: %v", err)
			continue
		}

		notificacion := types.NotificacionGenerica{
			TipoMensaje: "FalloCombate",
			Payload:     string(payloadBytes),
		}

		err = EnviarNotificacion(prod_ch, notificacion)
		if err != nil {
			logWithTimestamp("    [ ! ] Error reenviando mensaje a SNPNotifyQueue: %v", err)
		}
	}
}
