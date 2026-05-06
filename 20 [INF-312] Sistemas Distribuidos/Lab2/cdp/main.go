package main

import (
	"fmt"
	"os"
)

func main() {
	//=== VARIABLES DE ENTORNO ===//
	lcp_ip := os.Getenv("LCP_IP")
	if lcp_ip == "" {
		lcp_ip = "127.0.0.1"
	}
	lcp_port := os.Getenv("LCP_PORT")
	if lcp_port == "" {
		lcp_port = "50051"
	}

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

	rabbitConn, rabbitCh, err := conectarRabbit(rabbit_ip, rabbit_port, rabbit_user, rabbit_pass)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rabbitConn.Close()
	defer rabbitCh.Close()

	err = declararCDPResultQueue(rabbitCh)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = declararCDPErrorsQueue(rabbitCh)
	if err != nil {
		fmt.Println(err)
		return
	}

	consumerCh, err := rabbitConn.Channel()
	if err != nil {
		fmt.Println("Error creando canal para consumidor:", err)
		return
	}
	defer consumerCh.Close()

	lcpConn, clienteLCP, err := conectarLCP(lcp_ip, lcp_port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer lcpConn.Close()

	// Escuchar la cola 'GymResultQueue'
	llaves := map[string]string{
		"kanto":  os.Getenv("AES_KEY_KANTO"),
		"johto":  os.Getenv("AES_KEY_JOHTO"),
		"sinnoh": os.Getenv("AES_KEY_SINNOH"),
	}
	go escucharGymResultQueue(consumerCh, clienteLCP, llaves)

	// TODO: Descifrifar el mensaje recibido y verificar datos

	// TODO: Usar funciones gRPC de LCP para verificar entrenadores (lcpExisteEntrenador, lcpObtenerEntrenador)

	// TODO: Enviar resultados a la cola 'CDPResultQueue' y errores a la cola 'CDPErrorsQueue'

	// Mantener el programa corriendo para que la goroutine siga escuchando
	select {}
}
