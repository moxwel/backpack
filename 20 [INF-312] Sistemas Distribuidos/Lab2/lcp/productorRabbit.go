package main

import (
	"encoding/json"
	"fmt"
	"time"

	"l2/types"

	"github.com/streadway/amqp"
)

func conectarRabbit(rabbit_ip, rabbit_port, rabbit_user, rabbit_pass string) (*amqp.Connection, *amqp.Channel, error) {
	fmt.Println("[conectarRabbit] Conectando a RabbitMQ...")
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

func declararLCPEventsQueue(ch *amqp.Channel) error {
	fmt.Println("[declararLCPEventsQueue] Declarando cola LCPEventsQueue en RabbitMQ...")

	_, err := ch.QueueDeclare(
		"LCPEventsQueue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("    [ ! ] Error declarando cola LCPEventsQueue: %v", err)
	}

	fmt.Println("    [ OK ] Cola LCPEventsQueue declarada correctamente en RabbitMQ")
	return nil
}

func SendLCPEventsQueue(ch *amqp.Channel, tipo_mensaje string, payload string) error {
	fmt.Println("[SendLCPEventsQueue] Enviando notificación a la cola LCPEventsQueue...")

	notificacion := types.NotificacionGenerica{
		TipoMensaje: tipo_mensaje,
		Payload:     payload,
	}

	mensajeBytes, err := json.Marshal(notificacion)
	if err != nil {
		return fmt.Errorf("    [ ! ] Error serializando notificación para LCPEventsQueue: %v", err)
	}

	err = ch.Publish(
		"",               // exchange
		"LCPEventsQueue", // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        mensajeBytes,
		},
	)
	if err != nil {
		return fmt.Errorf("    [ ! ] Error enviando mensaje a la cola LCPEventsQueue: %v", err)
	}

	fmt.Println("    [ OK ] Notificación enviada a la cola LCPEventsQueue")
	return nil
}
