package main

import (
	"fmt"
	"math/rand"
	"net"
)

func main() {
	// Crear tableros
	tablero1 := [4]int{0, 0, 0, 0}
	tablero2 := [4]int{0, 0, 0, 0}

	rand1 := rand.Intn(4)
	rand2 := rand.Intn(4)

	tablero1[rand1] = 1
	tablero2[rand2] = 1

	fmt.Println("Tablero1:")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(tablero1[i*2+j], " ")
		}
		fmt.Println()
	}
	fmt.Println("Tablero2:")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(tablero2[i*2+j], " ")
		}
		fmt.Println()
	}

	// Obtener direccion UDP
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("Error al resolver la dirección:", err)
		return
	}
	fmt.Println("Dirección UDP resuelta:", addr)

	// Crear servidor UDP
	conn_udp, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
	var disponible byte = 1 // Servidor disponible
	buffer := make([]byte, 1024)
	fmt.Println("Servidor escuchando en UDP 8080...")
	defer conn_udp.Close()

	// Esperar petición de partida
	n, addr, err := conn_udp.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error al leer los datos:", err)
		return
	}
	fmt.Println("Mensaje recibido de:", addr, ":", string(buffer[:n]))
	fmt.Println("Respondiendo...")

	if n == 0 {
		fmt.Println("No se recibió ningún mensaje.")
		return
	}

	// Si el servidor no esta disponible, enviar 0
	if disponible == 0 {
		_, err = conn_udp.WriteToUDP([]byte{disponible}, addr)
		if err != nil {
			fmt.Println("Error al responder al cliente:", err)
			return
		}
	}

	// Responder al cliente con la disponibilidad del servidor
	_, err = conn_udp.WriteToUDP([]byte{disponible}, addr)
	if err != nil {
		fmt.Println("Error al responder al cliente:", err)
		return
	}

	disponible = 0 // Ahora el servidor no está disponible

	// Enviar datos al cliente: Puerto, IP, tablero1
	tcp_port := ":8081"
	_, err = conn_udp.WriteToUDP([]byte(tcp_port), addr)

	server_ip := addr.IP.String()
	_, err = conn_udp.WriteToUDP([]byte(server_ip), addr)

	// Enviar tablero1
	_, err = conn_udp.WriteToUDP([]byte{byte(tablero1[0]), byte(tablero1[1]), byte(tablero1[2]), byte(tablero1[3])}, addr)

	//server tcp para el juego

	// Crear servidor TCP
	escuchador, err := net.Listen("tcp", tcp_port)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
	fmt.Println("Servidor escuchando en TCP 8081...")
	defer escuchador.Close()

	continuar := 1
	ganador := 0

	// recibir movimiento del cliente
	conexion, err := escuchador.Accept()
	if err != nil {
		fmt.Println("Error al aceptar la conexión:", err)
		return
	}
	fmt.Println("Cliente conectado desde:", conexion.RemoteAddr())
	for continuar == 1 {
		// ver movimiento del cliente
		n, err := conexion.Read(buffer)
		if err != nil {
			fmt.Println("Error al leer los datos:", err)
			return
		}
		ataque := int(buffer[:n][0])
		fmt.Println("Mensaje recibido del cliente:", ataque)
		if tablero2[ataque] == 1 {
			tablero2[ataque] = 0
			fmt.Println("Ataque exitoso del cliente")
			continuar = 0
			ganador = 1
		}
		// jugada del servidor
		if continuar == 1 {
			ataque = rand.Intn(4)
			fmt.Println("Ataque del servidor:", ataque)
			if tablero1[ataque] == 1 {
				tablero1[ataque] = 0
				fmt.Println("Ataque exitoso del servidor")
				continuar = 0
				ganador = 2
			}
		} else {
			fmt.Println("Servidor derrotado.")
		}
		// enviar continuar al cliente
		fmt.Println("Enviando continuar (", continuar, ") al cliente...")
		_, err = conexion.Write([]byte{byte(continuar)})
		if err != nil {
			fmt.Println("Error al responder al cliente:", err)
			return
		}
	}
	// enviar ganador al cliente
	fmt.Println("Enviando ganador (", ganador, ") al cliente...")
	_, err = conexion.Write([]byte{byte(ganador)})
	if err != nil {
		fmt.Println("Error al responder al cliente:", err)
		return
	}
	fmt.Println("Fin del juego.")
}
