package main

import (
	"fmt"
	"os"
	"time"
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

func main() {
	//=== VARIABLES DE ENTORNO ===//
	rabbit_ip := os.Getenv("RABBIT_IP")
	if rabbit_ip == "" {
		rabbit_ip = "127.0.0.1"
	}
	rabbit_port := os.Getenv("RABBIT_PORT")
	if rabbit_port == "" {
		rabbit_port = "5672"
	}
	rabbit_user := os.Getenv("RABBIT_USER")
	if rabbit_user == "" {
		rabbit_user = "guest"
	}
	rabbit_pass := os.Getenv("RABBIT_PASS")
	if rabbit_pass == "" {
		rabbit_pass = "guest"
	}
	//============

	rabbitConn, rabbitChProducer, err := conectarRabbit(rabbit_ip, rabbit_port, rabbit_user, rabbit_pass)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rabbitConn.Close()
	defer rabbitChProducer.Close()

	// Crear un canal separado para escuchar
	_, rabbitChListener, err := conectarRabbit(rabbit_ip, rabbit_port, rabbit_user, rabbit_pass)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rabbitChListener.Close()

	err = declararSNPNotifyQueue(rabbitChProducer)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = declararLCPEventsQueue(rabbitChProducer)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = declararCDPErrorsQueue(rabbitChProducer)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = declararCDPResultQueue(rabbitChProducer)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = declararGymResultQueue(rabbitChProducer)
	if err != nil {
		fmt.Println(err)
		return
	}
	// TODO: Consumir mensajes de la cola 'LCPEventsQueue' y 'CDPErrorsQueue'

	go escucharLCPEventsQueue(rabbitChListener, rabbitChProducer)
	go escucharCDPErrorsQueue(rabbitChListener, rabbitChProducer)
	// TODO: Emitir mensajes a la cola 'SNPNotifyQueue'
	fmt.Println("[SNP Consumer] Escuchando colas de RabbitMQ...")
	select {} // Mantener el programa en ejecución
}
