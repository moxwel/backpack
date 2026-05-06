package main

// TODO: Añadir función para escuchar la cola 'LCPEventsQueue'

import (
	"encoding/json"

	"l2/types"

	"github.com/streadway/amqp"
)

func escucharLCPEventsQueue(lis_ch, prod_ch *amqp.Channel) {
	logWithTimestamp("[EscucharLCPEventsQueue] Escuchando la cola LCPEventsQueue...")

	msgs, err := lis_ch.Consume(
		"LCPEventsQueue", // queue
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		logWithTimestamp("    [ ! ] Error al consumir mensajes de la cola LCPEventsQueue: %v", err)
		return
	}

	for msg := range msgs {
		logWithTimestamp("[Cola: LCPEventsQueue] Mensaje recibido de LCPEventsQueue: %s\n", msg.Body)
		var notificacion types.NotificacionGenerica
		if err := json.Unmarshal(msg.Body, &notificacion); err != nil {
			logWithTimestamp("    [ ! ] Error deserializando mensaje: %v", err)
			continue
		}
		logWithTimestamp("    [ OK ] Notificación deserializada: %+v", notificacion)
		// Aquí podrías hacer lógica adicional según notificacion.TipoMensaje si lo necesitas
		err = EnviarNotificacion(prod_ch, notificacion)
		if err != nil {
			logWithTimestamp("    [ ! ] Error enviando notificación a SNPNotifyQueue: %v", err)
		} else {
			logWithTimestamp("    [ OK ] Notificación reenviada a SNPNotifyQueue")
		}
	}
}
