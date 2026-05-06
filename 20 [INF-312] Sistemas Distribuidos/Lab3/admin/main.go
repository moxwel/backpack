package main

import (
	"fmt"
	"l3/types"
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
	/* === ENVIRONMENT === */

	match_ip := os.Getenv("MATCH_IP")
	if match_ip == "" {
		match_ip = "matchmaker"
	}
	match_port := os.Getenv("MATCH_PORT")
	if match_port == "" {
		match_port = "50051"
	}

	id := os.Getenv("ID")
	if id == "" {
		id = "admin"
	}

	/* =================== */

	clock := types.VectorClockMap{}

	clock[id] = 0

	matchConn, clienteMatch, err := conectarServer(match_ip, match_port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer matchConn.Close()

	var opcion int
	var jugadores []types.PlayerQueueEntry
	var servidores []types.ServerState

	for {
		fmt.Printf("=======BIENVENIDO ADMINISTRADOR=======\nSeleccione una acción a realizar:\n1:Obtener el estado de los servidores\n2:Modificar el estado de un servidor\n3: Salir\n")
		_, err := fmt.Scanf("%d", &opcion)
		if err != nil {
			fmt.Println("Error al leer la opción:", err)
			continue
		}
		switch opcion {
		case 1:
			logWithTimestamp("[AdminMain] Solicitando estado del sistema")
			logWithTimestamp("[AdminMain] Reloj actual: %v", clock)

			status, err := AdminGetSystemStatus(clienteMatch, clock)
			if err != nil {
				logWithTimestamp("[AdminMain] [ERROR] Error al obtener estado del sistema: %v", err)
				incrementarReloj(clock, id)
				continue
			}

			relojAnterior := clock
			newClock := status.VectorClock
			clock = fusionarMayores(clock, newClock)
			incrementarReloj(clock, id)

			logWithTimestamp("[AdminMain] Reloj anterior: %v, Reloj recibido: %v, Reloj fusionado: %v", relojAnterior, newClock, clock)

			servidores = status.ServerList
			jugadores = status.PlayerQueue
			logWithTimestamp("[AdminMain] Estado obtenido - %d servidores, %d jugadores", len(servidores), len(jugadores))
			printServers(servidores)
			printJugadores(jugadores)
		case 2:
			logWithTimestamp("[AdminMain] Iniciando actualización de estado de servidor")
			logWithTimestamp("[AdminMain] Reloj actual: %v", clock)

			status, err := AdminGetSystemStatus(clienteMatch, clock)

			if err != nil {
				logWithTimestamp("[AdminMain] [ERROR] Error al obtener estado para actualización: %v", err)
				incrementarReloj(clock, id)
				continue
			}

			relojAnterior := clock
			newClock := status.VectorClock
			clock = fusionarMayores(clock, newClock)
			logWithTimestamp("[AdminMain] Reloj anterior: %v, Reloj recibido: %v, Reloj fusionado: %v", relojAnterior, newClock, clock)

			servidores = status.ServerList
			fmt.Printf("Seleccione un servidor a actualizar:\n")
			printServers(servidores)
			var servidor string
			var newStatus int
			_, err = fmt.Scanf("%s", &servidor)
			if err != nil {
				logWithTimestamp("[AdminMain] [ERROR] Error al leer servidor: %v", err)
				continue
			}

			//verificar si el servidor existe
			found := false
			for _, s := range servidores {
				if s.Id == servidor {
					found = true
					break
				}
			}
			if !found {
				logWithTimestamp("[AdminMain] Servidor %s no encontrado", servidor)
				fmt.Println("Servidor no encontrado. Operacion cancelada.")
				continue
			}
			fmt.Printf("Seleccione el nuevo estatus:\n")
			fmt.Printf("0: DISPONIBLE\n1: OCUPADO\n2: CAIDO\n3:DESCONOCIDO\n4: CANCELAR\n")
			_, err = fmt.Scanf("%d", &newStatus)
			if err != nil {
				logWithTimestamp("[AdminMain] [ERROR] Error al leer nuevo estado: %v", err)
				continue
			}
			switch newStatus {
			case 0, 1, 2, 3:
				logWithTimestamp("[AdminMain] Actualizando servidor %s a estado %d", servidor, newStatus)
				logWithTimestamp("[AdminMain] Reloj antes de actualización: %v", clock)

				status, err := AdminUpdateServerState(clienteMatch, servidor, types.ServerStatus(newStatus), clock)
				if err != nil {
					logWithTimestamp("[AdminMain] [ERROR] Error al actualizar estado del servidor: %v", err)
					incrementarReloj(clock, id)
					continue
				}
				relojAnterior = clock
				newClock = status.VectorClock
				clock = fusionarMayores(clock, newClock)
				incrementarReloj(clock, id)

				logWithTimestamp("[AdminMain] Servidor actualizado - Código: %d, Mensaje: %s", status.StatusCode, status.Message)
				logWithTimestamp("[AdminMain] Reloj anterior: %v, Reloj recibido: %v, Reloj fusionado: %v", relojAnterior, newClock, clock)
			case 4:
				logWithTimestamp("[AdminMain] Operación de actualización cancelada por usuario")
				fmt.Println("Operación cancelada.")
				continue
			default:
				logWithTimestamp("[AdminMain] Opción inválida seleccionada: %d", newStatus)
				fmt.Println("Opción no válida. Intente nuevamente.")
			}
		case 3:
			logWithTimestamp("[AdminMain] Cerrando aplicación de administrador")
			fmt.Println("Saliendo del programa...")
			return
		default:
			logWithTimestamp("[AdminMain] Opción inválida en menú principal: %d", opcion)
			fmt.Println("Opción no válida. Intente nuevamente.")
		}
	}
}

func printServers(servers []types.ServerState) {
	for _, server := range servers {
		fmt.Printf("=====Servidor %s=====\n", server.Id)
		fmt.Printf("    Status: %d\n", server.Status)
		fmt.Printf("    Address: %s\n", server.Address)
		if server.CurrentMatchId == "" {
			fmt.Printf("    Current Match: ninguno\n")
		} else {
			fmt.Printf("    Current Match: %s\n", server.CurrentMatchId)
		}
		fmt.Printf("=====     =====\n")
	}
}

func printJugadores(players []types.PlayerQueueEntry) {
	for _, player := range players {
		fmt.Printf("=====Player %s=====\n", player.PlayerId)
		fmt.Printf("     Time in Queue: %d\n", player.TimeInQueue)
		fmt.Printf("=====     =====\n")
	}
}
