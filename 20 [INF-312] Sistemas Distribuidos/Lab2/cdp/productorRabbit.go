package main

import (
	"encoding/json"
	"fmt"
	"l2/types"
	"time"

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

func declararCDPResultQueue(ch *amqp.Channel) error {
	fmt.Println("[declararCDPResultQueue] Declarando cola CDPResultQueue en RabbitMQ...")

	_, err := ch.QueueDeclare(
		"CDPResultQueue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("    [ ! ] Error declarando cola CDPResultQueue: %v", err)
	}

	fmt.Println("    [ OK ] Cola CDPResultQueue declarada correctamente en RabbitMQ")
	return nil
}

func declararCDPErrorsQueue(ch *amqp.Channel) error {
	fmt.Println("[declararCDPErrorsQueue] Declarando cola CDPErrorsQueue en RabbitMQ...")

	_, err := ch.QueueDeclare(
		"CDPErrorsQueue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("    [ ! ] Error declarando cola CDPErrorsQueue: %v", err)
	}

	fmt.Println("    [ OK ] Cola CDPErrorsQueue declarada correctamente en RabbitMQ")
	return nil
}

// TODO: Añadir función para enviar mensajes a la cola 'CDPResultQueue' y 'CDPErrorsQueue'

func SendCDPResultQueue(ch *amqp.Channel, message types.ResultadoCombate) error {
	fmt.Println("[SendCDPResultQueue] Enviando mensaje a la cola CDPResultQueue...")

	mensajeBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("    [ ! ] Error serializando mensaje para CDPResultQueue: %v", err)
	}

	for i := 1; i <= 3; i++ {
		err = ch.Publish(
			"",               // exchange
			"CDPResultQueue", // routing key
			false,            // mandatory
			false,            // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        mensajeBytes,
			},
		)
		if err == nil {
			fmt.Println("    [ OK ] Mensaje enviado a la cola CDPResultQueue")
			return nil
		}
		fmt.Printf("    [ ! ] Error enviando mensaje a la cola CDPResultQueue (intento %d/3): %v\n", i, err)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("    [ ! ] No se pudo enviar el mensaje a la cola CDPResultQueue tras 3 intentos")
	return err
}

func SendCDPErrorsQueue(ch *amqp.Channel, message types.FalloCombate) error {
	fmt.Println("[SendCDPErrorsQueue] Enviando mensaje a la cola CDPErrorsQueue...")

	mensajeBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("    [ ! ] Error serializando mensaje para CDPErrorsQueue: %v", err)
	}

	for i := 1; i <= 3; i++ {
		err = ch.Publish(
			"",               // exchange
			"CDPErrorsQueue", // routing key
			false,            // mandatory
			false,            // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        mensajeBytes,
			},
		)
		if err == nil {
			fmt.Println("    [ OK ] Mensaje enviado a la cola CDPErrorsQueue")
			return nil
		}
		fmt.Printf("    [ ! ] Error enviando mensaje a la cola CDPErrorsQueue (intento %d/3): %v\n", i, err)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("    [ ! ] No se pudo enviar el mensaje a la cola CDPErrorsQueue tras 3 intentos")
	return err
}
