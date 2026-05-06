package main

import (
	"encoding/json"
	"fmt"
	"l2/types"
	"time"

	"github.com/streadway/amqp"
)

func conectarRabbit(rabbit_ip, rabbit_port, rabbit_user, rabbit_pass string) (*amqp.Connection, *amqp.Channel, error) {
	logWithTimestamp("[conectarRabbit] Conectando a RabbitMQ...")
	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbit_user, rabbit_pass, rabbit_ip, rabbit_port)

	var conn *amqp.Connection
	var ch *amqp.Channel
	var err error

	for i := 1; i <= 5; i++ {
		logWithTimestamp("    [ * ] Intentando conectar a RabbitMQ (intento %d/5)...", i)
		conn, err = amqp.Dial(amqpURL)
		if err == nil {
			ch, err = conn.Channel()
			if err == nil {
				logWithTimestamp("    [ OK ] Conexión a RabbitMQ establecida correctamente")
				return conn, ch, nil
			}
			conn.Close()
		}
		logWithTimestamp("    [ ! ] Error conectando a RabbitMQ.")
		if i < 5 {
			logWithTimestamp(" Esperando 5 segundos antes de reintentar...")
			time.Sleep(5 * time.Second)
		}
	}
	return nil, nil, fmt.Errorf("    [ ! ] Error conectando a RabbitMQ después de 5 intentos: %v", err)
}

func declararGymResultQueue(ch *amqp.Channel) error {
	logWithTimestamp("[declararGymResultQueue] Declarando cola GymResultQueue en RabbitMQ...")

	_, err := ch.QueueDeclare(
		"GymResultQueue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("    [ ! ] Error declarando cola GymResultQueue: %v", err)
	}

	logWithTimestamp("    [ OK ] Cola GymResultQueue declarada correctamente en RabbitMQ")
	return nil
}

func SendGymResultQueue(ch *amqp.Channel, message types.CombateGymMessage) error {
	logWithTimestamp("[SendGymResultQueue] Enviando mensaje a GymResultQueue...")
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("    [ ! ] Error serializando mensaje a JSON: %v", err)
	}
	messageStr := string(messageJSON)
	logWithTimestamp("    [ * ] Mensaje a enviar: %s", messageStr)

	err = ch.Publish(
		"",               // exchange
		"GymResultQueue", // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(messageStr),
		},
	)
	if err != nil {
		return fmt.Errorf("    [ ! ] Error enviando mensaje a GymResultQueue: %v", err)
	}

	logWithTimestamp("    [ OK ] Mensaje enviado a GymResultQueue correctamente")
	return nil
}
