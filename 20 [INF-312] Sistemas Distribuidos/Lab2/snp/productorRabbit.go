package main

import (
	"encoding/json"
	"fmt"
	"time"

	"l2/types"

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

func declararSNPNotifyQueue(ch *amqp.Channel) error {
	logWithTimestamp("[declararSNPNotifyQueue] Declarando cola SNPNotifyQueue en RabbitMQ...")

	_, err := ch.QueueDeclare(
		"SNPNotifyQueue", // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,
	)
	if err != nil {
		return fmt.Errorf("    [ ! ] Error declarando cola SNPNotifyQueue: %v", err)
	}

	logWithTimestamp("    [ OK ] Cola SNPNotifyQueue declarada correctamente en RabbitMQ")
	return nil
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

func declararGymResultQueue(ch *amqp.Channel) error {
	fmt.Println("[declararGymResultQueue] Declarando cola GymResultQueue en RabbitMQ...")
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
	fmt.Println("    [ OK ] Cola GymResultQueue declarada correctamente en RabbitMQ")
	return nil
}

func EnviarNotificacion(ch *amqp.Channel, msg types.NotificacionGenerica) error {
	logWithTimestamp("[EnviarNotificacion] Enviando notificación a SNPNotifyQueue...")
	mensajeJSON, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("    [ ! ] Error serializando notificación: %v", err)
	}
	logWithTimestamp("    Notificación: %s", string(mensajeJSON))

	returns := make(chan amqp.Return, 1)
	ch.NotifyReturn(returns)

	amqpMsg := amqp.Publishing{
		ContentType: "application/json",
		Body:        mensajeJSON,
	}

	var publishErr error
	for i := 1; i <= 3; i++ {
		publishErr = ch.Publish(
			"",               // exchange
			"SNPNotifyQueue", // routing key
			true,             // mandatory: para detectar si no se enruta
			false,            // immediate
			amqpMsg,
		)
		if publishErr == nil {
			select {
			case ret := <-returns:
				logWithTimestamp("    [ ! ] Mensaje no enrutado (intento %d/3): %s", i, string(ret.Body))
				publishErr = fmt.Errorf("mensaje no enrutado")
			default:
				logWithTimestamp("    [ OK ] Notificación enviada correctamente a SNPNotifyQueue")
				return nil
			}
		} else {
			logWithTimestamp("    [ ! ] Error enviando notificación (intento %d/3): %v", i, publishErr)
		}
		time.Sleep(1 * time.Second)
	}

	if publishErr != nil {
		logWithTimestamp("    [ ! ] No se pudo enviar la notificación tras 3 intentos")
	}
	return publishErr
}
