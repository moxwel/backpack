package main

import (
	"fmt"
	"net"
)

func main() {
	// Obtener direccion UDP
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("Error al resolver la dirección:", err)
		return
	}
	fmt.Println("Dirección UDP resuelta:", addr)

	// Crear conexion UDP
	fmt.Println("Conectando al servidor UDP:", addr, "...")
	conn_udp, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error al conectarse al servidor:", err)
		return
	}
	fmt.Println("Conexion exitosa.")
	defer conn_udp.Close()

	// Enviar mensaje de solicitud de partida
	fmt.Println("Enviando solicitud de partida...")
	msg := []byte("Iniciar partida.")
	_, err = conn_udp.Write(msg)
	if err != nil {
		fmt.Println("Error al enviar el mensaje:", err)
		return
	}

	// Recibir respuesta del servidor
	fmt.Println("Esperando respuesta del servidor...")
	buffer := make([]byte, 1024)
	n, addr, err := conn_udp.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error al leer los datos:", err)
		return
	}
	fmt.Println("Mensaje recibido de:", addr, ":", string(buffer[:n]))

	if n == 0 {
		fmt.Println("No se recibió ningún mensaje.")
		return
	}

	// Si se recibió un 0, el servidor no está disponible
	if buffer[0] == 0 {
		fmt.Println("El servidor no está disponible.")
		return
	}

	// Si se recibió un 1, el servidor está disponible y
	// comenzara a recibir los datos del servidor.

	dispo := int(buffer[0]) // Se guarda disponibilidad del servidor
	fmt.Println("Servidor disponible:", dispo == 1)

	// Recibir datos del servidor
	// Puerto
	n, addr, err = conn_udp.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error al leer los datos:", err)
		return
	}
	tcp_port := string(buffer[:n])
	fmt.Println("Puerto TCP:", tcp_port)

	// IP
	n, addr, err = conn_udp.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error al leer los datos:", err)
		return
	}
	server_ip := string(buffer[:n])
	fmt.Println("IP del servidor:", server_ip)

	// Tablero1
	n, addr, err = conn_udp.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error al leer los datos:", err)
		return
	}
	tablero := make([]byte, n)
	n_copy := copy(tablero, buffer)
	fmt.Println("Tablero copiado:", n_copy, "bytes.")

	fmt.Println("Tablero:")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(tablero[i*2+j], " ")
		}
		fmt.Println()
	}

	// Crear conexion TCP
	addr_tcp, err := net.ResolveTCPAddr("tcp", server_ip+tcp_port)
	if err != nil {
		fmt.Println("Error al resolver la dirección:", err)
		return
	}
	fmt.Println("Dirección TCP resuelta:", addr_tcp)

	// Conectar al servidor TCP
	fmt.Println("Conectando al servidor TCP:", addr_tcp, "...")
	conn_tcp, err := net.DialTCP("tcp", nil, addr_tcp)
	if err != nil {
		fmt.Println("Error al conectarse al servidor:", err)
		return
	}
	fmt.Println("Conexion exitosa.")
	defer conn_tcp.Close()
	tablero_sel := []string{"A", "B", "C", "D"}
	var continuar byte = 1
	for continuar == 1 {
		fmt.Println("Tu tablero:")
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				fmt.Print(tablero[i*2+j], " ")
			}
			fmt.Println()
		}
		//escoger jugada
		fmt.Println("Ingrese su jugada: (A,B,C,D)")
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				fmt.Print(tablero_sel[i*2+j], " ")
			}
			fmt.Println()
		}
		var jugada string
		fmt.Scan(&jugada)

		var numero int
		switch jugada {
		case "A", "a":
			numero = 0
		case "B", "b":
			numero = 1
		case "C", "c":
			numero = 2
		case "D", "d":
			numero = 3
		default:
			fmt.Println("Jugada inválida. Fin del juego.")
			return
		}

		//enviar jugada al servidor
		fmt.Println("Enviando jugada al servidor...")
		_, err = conn_tcp.Write([]byte{byte(numero)})
		if err != nil {
			fmt.Println("Error al enviar el mensaje:", err)
			return
		}
		//recibir continuar del servidor para verificar si se sigue jugando
		fmt.Println("Esperando respuesta del servidor...")
		n, err = conn_tcp.Read(buffer)
		if err != nil {
			fmt.Println("Error al leer los datos:", err)
			return
		}
		continuar = buffer[0]
		fmt.Println("Recibido:", continuar)

	}

	// Recibir ganador
	fmt.Println("Termino el juego.")
	fmt.Println("Esperando resultados del servidor...")
	n, err = conn_tcp.Read(buffer)
	if err != nil {
		fmt.Println("Error al leer los datos:", err)
		return
	}
	ganador := buffer[0]
	if ganador == 1 {
		fmt.Println("Ganaste!")
	} else {
		fmt.Println("Perdiste!")
	}
	fmt.Println("Fin del juego.")
}
