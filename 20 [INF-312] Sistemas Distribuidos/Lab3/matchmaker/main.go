package main

import (
	"context"
	"fmt"
	"l3/types"
	"net"
	"os"
	"sync"
	"time"

	pb "matchmaker/MatchmakingProto"

	"github.com/google/uuid"
	"google.golang.org/grpc"
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

	server_ip := os.Getenv("IP")
	if server_ip == "" {
		server_ip = "0.0.0.0"
	}
	server_port := os.Getenv("PORT")
	if server_port == "" {
		server_port = "50051"
	}
	id := os.Getenv("ID")
	if id == "" {
		id = "matchmaker"
	}

	/* =================== */

	logWithTimestamp("[--!--]    Iniciando Servidor MATCHMAKING")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server_ip, server_port))
	if err != nil {
		panic("[ERROR]    No se pudo iniciar el listener: " + err.Error())
	}

	server := &ServidorMatchmaking{
		id:          id,
		servidores:  map[string]types.Server{},
		jugadores:   map[string]types.Player{},
		queuePlayer: []string{},
		clients:     map[string]pb.GameServerServiceClient{},
		conns:       map[string]*grpc.ClientConn{},
		reloj:       types.VectorClockMap{id: 0},
		mutex:       sync.Mutex{},
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMatchmakingServiceServer(grpcServer, server)

	go monitorearServidores(server)
	go asignarMatchmaking(server)

	logWithTimestamp("[--!--]    Servidor Matchmaking escuchando en: %s:%s ", server_ip, server_port)
	if err := grpcServer.Serve(listener); err != nil {
		panic("[ERROR]    Fallo al servir: " + err.Error())
	}

	for _, conn := range server.conns {
		conn.Close()
	}
}

func asignarMatchmaking(server *ServidorMatchmaking) {
	for {
		logWithTimestamp("[asignarMatchmaking] Verificando cola de jugadores y servidores disponibles")
		// Contar servidores disponibles
		disponibles := 0
		for _, servidor := range server.servidores {
			if servidor.Status == 0 {
				disponibles++
			}
		}
		logWithTimestamp("[asignarMatchmaking] Estado actual - Jugadores en cola: %d, Servidores disponibles: %d", len(server.queuePlayer), disponibles)
		fmt.Printf("Jugadores en cola: %d\nServidores disponibles: %d\n", len(server.queuePlayer), disponibles)

		if len(server.queuePlayer) >= 2 {
			for serverId, servidor := range server.servidores {
				if servidor.Status == 0 {

					logWithTimestamp("[asignarMatchmaking] Servidor disponible encontrado: %s", serverId)

					logWithTimestamp("[asignarMatchmaking] Iniciando asignación de partida en servidor %s", serverId)
					logWithTimestamp("[asignarMatchmaking] Jugador1: %s, Jugador2: %s", server.queuePlayer[0], server.queuePlayer[1])

					player1 := server.queuePlayer[0]
					player2 := server.queuePlayer[1]

					matchId := uuid.New().String()

					logWithTimestamp("[asignarMatchmaking] MatchID generado: %s", matchId)

					servidor.Player1 = player1
					servidor.Player2 = player2
					servidor.CurrentMatchId = matchId

					jugador1 := server.jugadores[player1]
					jugador1.Status = types.IN_MATCH
					jugador1.MatchId = matchId
					jugador1.ServerAddress = servidor.Address

					jugador2 := server.jugadores[player2]
					jugador2.Status = types.IN_MATCH
					jugador2.MatchId = matchId
					jugador2.ServerAddress = servidor.Address

					relojAntes := server.reloj
					incrementarReloj(server.reloj, server.id)
					logWithTimestamp("[asignarMatchmaking] Reloj antes de asignación: %v, Reloj después: %v", relojAntes, server.reloj)

					// Actualizar el estado del servidor ("read your writes")
					server.servidores[serverId] = servidor

					ctx := context.Background()
					logWithTimestamp("[asignarMatchmaking] Enviando solicitud AssignMatch al servidor %s", serverId)
					logWithTimestamp("[asignarMatchmaking] Reloj enviado: %v", server.reloj)

					newReloj, err := AssignMatch(ctx, server.clients[serverId], serverId, matchId, player1, player2, server.reloj)
					if err != nil {
						logWithTimestamp("[asignarMatchmaking] [ERROR] Error al asignar partida: %v", err)
						break
					}

					relojAnterior := server.reloj
					server.reloj = fusionarMayores(server.reloj, newReloj)
					logWithTimestamp("[asignarMatchmaking] Partida asignada exitosamente")
					logWithTimestamp("[asignarMatchmaking] Reloj anterior: %v, Reloj recibido: %v, Reloj fusionado: %v", relojAnterior, newReloj, server.reloj)

					server.jugadores[player1] = jugador1
					server.jugadores[player2] = jugador2

					server.queuePlayer = server.queuePlayer[2:]

					logWithTimestamp("[asignarMatchmaking] Partida completamente configurada: %s vs %s en servidor %s con MatchID %s", player1, player2, serverId, matchId)

					break
				}
			}
		}
		time.Sleep(5 * time.Second)
	}
}

func monitorearServidores(server *ServidorMatchmaking) {
	logWithTimestamp("[monitorearServidores] Iniciando monitoreo de servidores")
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		logWithTimestamp("[monitorearServidores] Iniciando ronda de verificación de servidores")
		for serverId, servidor := range server.servidores {
			if servidor.Status == 0 {
				logWithTimestamp("[monitorearServidores] Enviando ping a servidor: %s", serverId)
				logWithTimestamp("[monitorearServidores] Reloj antes del ping: %v", server.reloj)

				ctx := context.Background()
				newReloj, err := PingServer(ctx, server.clients[serverId], serverId, server, server.reloj)
				if err != nil {
					logWithTimestamp("[monitorearServidores] Servidor %s no responde, marcando como caído. Error: %v", serverId, err)
					servidor.Status = types.CAIDO
					server.servidores[serverId] = servidor
				} else {
					relojAnterior := server.reloj
					server.reloj = fusionarMayores(server.reloj, newReloj)
					logWithTimestamp("[monitorearServidores] Ping exitoso a servidor %s", serverId)
					logWithTimestamp("[monitorearServidores] Reloj anterior: %v, Reloj recibido: %v, Reloj fusionado: %v", relojAnterior, newReloj, server.reloj)
				}
			}
		}
		logWithTimestamp("[monitorearServidores] Ronda de verificación completada")
		time.Sleep(10 * time.Second)
	}
}
