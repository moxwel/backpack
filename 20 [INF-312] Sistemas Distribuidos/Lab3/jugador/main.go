package main

import (
	"fmt"
	pb "jugador/MatchmakingProto"
	"l3/types"
	"os"
	"strings"
	"sync"
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
		id = "jugador"
	}

	/* =================== */

	var mu sync.Mutex

	clock := types.VectorClockMap{}

	clock[id] = 0

	matchConn, clienteMatch, err := conectarServer(match_ip, match_port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer matchConn.Close()
	var opcion int
	var jugador string
	ids := []string{}
	enCola := false
	jugadorEnCola := ""

	go pollPlayerStatus(&enCola, clienteMatch, &jugadorEnCola, &clock, id, &mu)

	for {
		fmt.Printf("=======BIENVENIDO JUGADOR=======\n")

		// decir si hay un jugador en cola
		mu.Lock()
		if enCola {
			fmt.Printf("Jugador en cola: %s\n", jugadorEnCola)
		} else {
			fmt.Printf("No hay jugadores en cola.\n")
		}
		mu.Unlock()

		fmt.Printf("Seleccione una acción a realizar:\n1:Entrar en cola\n2:Obtener estado\n3:Salir de cola\n4: Salir\n")
		_, err := fmt.Scanf("%d", &opcion)
		if err != nil {
			fmt.Println("Error al leer la opción:\n", err)
			continue
		}
		switch opcion {
		case 1:
			mu.Lock()
			if enCola {
				mu.Unlock()
				fmt.Println("Ya hay un jugador en cola. No se puede entrar en cola nuevamente.")
				continue
			}
			mu.Unlock()

			fmt.Printf("Seleccione perfil jugador (o escriba uno nuevo):\n")
			printIds(ids)
			_, err := fmt.Scanf("%s", &jugador)
			if err != nil {
				fmt.Println("Error al leer el jugador:\n", err)
				continue
			}

			mu.Lock()
			clockCopia := clock
			mu.Unlock()

			code, newClock, err := QueuePlayer(clienteMatch, jugador, clockCopia)

			mu.Lock()
			if err != nil {
				logWithTimestamp("[ERROR]    Error al entrar en cola: %v", err)
				incrementarReloj(clock, id)
			} else if code == 0 {
				fmt.Println("Jugador añadido a la cola con éxito.")
				enCola = true
				jugadorEnCola = jugador
				ids = append(ids, jugador)
				clock = fusionarMayores(clock, newClock)
				incrementarReloj(clock, id)
			} else {
				fmt.Printf("Error al entrar en cola: %d\n", code)
				logWithTimestamp("[ERROR]    Error al entrar en cola: %d", code)
				clock = fusionarMayores(clock, newClock)
				incrementarReloj(clock, id)
			}
			mu.Unlock()
		case 2:
			mu.Lock()
			if !enCola {
				mu.Unlock()
				fmt.Println("No hay jugador en cola para consultar.")
				continue
			}
			jugadorAConsultar := jugadorEnCola
			clockCopia := clock
			mu.Unlock()

			logWithTimestamp("[GetPlayerStatus] Consultando estado de jugador %s", jugadorAConsultar)
			logWithTimestamp("[GetPlayerStatus] Reloj enviado: %v", clockCopia)

			new, err := GetPlayerStatus(clienteMatch, jugadorAConsultar, clockCopia)
			if err != nil {
				logWithTimestamp("[ERROR]    Error al obtener status de jugador %s", jugadorAConsultar)
			} else {
				jugadorStatus := types.Player{
					Status:        new.Status,
					MatchId:       new.MatchId,
					ServerAddress: new.ServerAddress,
				}

				printJugador(jugadorAConsultar, jugadorStatus)

				mu.Lock()
				clock = fusionarMayores(clock, new.VectorClock)
				incrementarReloj(clock, id)
				mu.Unlock()
				logWithTimestamp("[GetPlayerStatus] Reloj anterior: %v, Reloj nuevo: %v", clockCopia, clock)
			}
		case 3:
			mu.Lock()
			if !enCola {
				mu.Unlock()
				fmt.Println("No hay jugador en cola para sacar.")
				continue
			}
			jugadorASacar := jugadorEnCola
			clockCopia := clock
			mu.Unlock()

			fmt.Printf("Saliendo de cola de jugador %s\n", jugadorASacar)

			res, err := CancelQueue(clienteMatch, jugadorASacar, clockCopia)

			mu.Lock()
			if err != nil {
				logWithTimestamp("[ERROR]    Error al salir de cola: %v", err)
				incrementarReloj(clock, id)
			} else {
				enCola = false
				jugadorEnCola = ""
				clock = fusionarMayores(clock, fromProtoVectorClock(res.VectorClock.Entries))
				incrementarReloj(clock, id)
			}
			mu.Unlock()

		case 4:
			fmt.Printf("Saliendo del programa...\n")
			return
		default:
			fmt.Println("Opción no válida. Intente nuevamente.")
		}
	}
}

func printIds(players []string) {
	fmt.Printf("=======JUGADORES REGISTRADOS=======\n")
	for i, player := range players {
		fmt.Printf("Jugador %d: %s\n", i, player)
	}
	fmt.Printf("===================================\n")
}

func printJugador(playerId string, player types.Player) {
	fmt.Printf("=======Jugador: %s=======\n", playerId)
	switch player.Status {
	case types.IN_MATCH:
		fmt.Printf("Status: IN_MATCH\nMatchID: %s\n Server: %s\n", player.MatchId, player.ServerAddress)
	case types.IN_QUEUE:
		fmt.Printf("Status: IN_QUEUE\n")
	case types.IDLE:
		fmt.Printf("Status: IDLE\n")
	case types.UNKNOWN:
		fmt.Printf("Status: UNKNOW\n")
	}
	fmt.Printf("=========================\n")
}

func pollPlayerStatus(enCola *bool, clienteMatch pb.MatchmakingServiceClient, jugadorEnCola *string, clock *types.VectorClockMap, id string, mu *sync.Mutex) {
	for {
		time.Sleep(1 * time.Second)
		// Si no hay jugador en cola, no hacemos nada
		mu.Lock()
		if !*enCola {
			mu.Unlock()
			continue
		}
		currentJugador := *jugadorEnCola
		currentClock := *clock
		mu.Unlock()

		new, err := GetPlayerStatus(clienteMatch, currentJugador, currentClock)
		if err != nil {
			continue
		}

		jugadorStatus := types.Player{
			Status:        new.Status,
			MatchId:       new.MatchId,
			ServerAddress: new.ServerAddress,
		}

		mu.Lock()
		// Comprobar si seguimos siendo responsables de este jugador
		if !*enCola || *jugadorEnCola != currentJugador {
			mu.Unlock()
			continue
		}

		if jugadorStatus.Status == types.IN_MATCH {
			// el jugador está en partida, obtener el resultado de la partida
			logWithTimestamp("[pollPlayerStatus] Jugador %s en partida, obteniendo resultado...", currentJugador)
			ipPort := jugadorStatus.ServerAddress
			matchId := jugadorStatus.MatchId
			clockForMatchResult := *clock
			mu.Unlock()

			arr := strings.Split(ipPort, ":")
			if len(arr) < 2 {
				logWithTimestamp("[pollPlayerStatus] [ERROR] Dirección de servidor de partida inválida: %s", ipPort)
				continue
			}
			ip, port := arr[0], arr[1]

			conn, client, err := conectarGameServer(ip, port)
			if err != nil {
				logWithTimestamp("[pollPlayerStatus] [ERROR] Error al conectar con Game Server: %v", err)
				continue
			}

			logWithTimestamp("[pollPlayerStatus] Conectado a Game Server en %s:%s para obtener resultado de partida %s", ip, port, matchId)
			res, err := GetMatchResult(client, matchId, clockForMatchResult)
			conn.Close()
			if err != nil {
				logWithTimestamp("[pollPlayerStatus] [ERROR] Error al obtener resultado de partida: %v", err)
				continue
			}
			fmt.Printf("\nResultado de partida: \n")
			fmt.Printf("Ganador: %s\n", res.WinnerId)
			fmt.Printf("Perdedor: %s\n", res.LoserId)
			fmt.Printf("=========================\n")

			// Actualizar estado compartido
			mu.Lock()
			if *enCola && *jugadorEnCola == currentJugador {
				// Actualizar el reloj vectorial con el resultado de la partida
				relojAnterior := *clock
				relojRecibido := fromProtoVectorClock(res.VectorClock.Entries)
				*clock = fusionarMayores(*clock, relojRecibido)
				incrementarReloj(*clock, id)
				logWithTimestamp("[pollPlayerStatus] Reloj anterior: %v, Reloj recibido: %v, Reloj fusionado: %v", relojAnterior, relojRecibido, *clock)

				// marcar jugador como libre
				*enCola = false
				*jugadorEnCola = ""
			}
			mu.Unlock()
		} else {
			// El jugador no está en partida (sigue en cola), solo actualizar reloj
			relojRecibido := new.VectorClock
			*clock = fusionarMayores(*clock, relojRecibido)
			mu.Unlock()
		}
	}
}
