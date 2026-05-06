package main

import (
	"context"
	"fmt"
	"l3/types"
	"strings"
	"sync"
	"time"

	pb "matchmaker/MatchmakingProto"

	"google.golang.org/grpc"
)

type ServidorMatchmaking struct {
	pb.UnimplementedMatchmakingServiceServer
	id          string                                // Id del servidor
	servidores  map[string]types.Server               // Mapeo por Id
	jugadores   map[string]types.Player               // Mapeo por Id
	queuePlayer []string                              //Lista de espera de jugadores
	clients     map[string]pb.GameServerServiceClient //Mapeo por Id de los clientes
	conns       map[string]*grpc.ClientConn           // Mapeo por Id de las conexiones
	reloj       types.VectorClockMap                  // Reloj vectorial para el servidor
	mutex       sync.Mutex                            // Semáforo para operaciones de cola
}

func toProtoVectorClock(vc types.VectorClockMap) *pb.VectorClock {
	entries := []*pb.VectorClockEntry{}
	for id, timestamp := range vc {
		entries = append(entries, &pb.VectorClockEntry{
			ServerId:  id,
			Timestamp: timestamp,
		})
	}
	return &pb.VectorClock{
		Entries: entries,
	}
}

func fromProtoVectorClock(entries []*pb.VectorClockEntry) types.VectorClockMap {
	vc := types.VectorClockMap{}
	for _, entry := range entries {
		vc[entry.ServerId] = entry.Timestamp
	}
	return vc
}

func incrementarReloj(vc types.VectorClockMap, nodeId string) {
	vc[nodeId]++
}

func fusionarMayores(r1, r2 types.VectorClockMap) types.VectorClockMap {
	result := make(types.VectorClockMap)
	for id := range r1 {
		if r1[id] > r2[id] {
			result[id] = r1[id]
		} else {
			result[id] = r2[id]
		}
	}
	for id := range r2 {
		if _, ok := result[id]; !ok {
			result[id] = r2[id]
		}
	}
	return result
}

func (s *ServidorMatchmaking) QueuePlayer(ctx context.Context, player *pb.PlayerInfoRequest) (*pb.QueuePlayerResponse, error) {

	logWithTimestamp("[QueuePlayer] Nueva solicitud de cola de jugador: %s", player.PlayerId)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	relojEntrante := fromProtoVectorClock(player.VectorClock.Entries)
	relojAnterior := s.reloj
	logWithTimestamp("[QueuePlayer] Reloj anterior: %v, Reloj entrante: %v", relojAnterior, relojEntrante)

	s.reloj = fusionarMayores(s.reloj, relojEntrante)
	incrementarReloj(s.reloj, s.id)
	logWithTimestamp("[QueuePlayer] Reloj fusionado: %v", s.reloj)

	relojProto := toProtoVectorClock(s.reloj)

	if jugador, exists := s.jugadores[player.PlayerId]; !exists {
		logWithTimestamp("[QueuePlayer] Jugador no existe, creando jugador...")
		jugador = types.Player{
			Status:        types.IN_QUEUE,
			TimeInQueue:   0,
			QueueSince:    time.Now(),
			ServerAddress: "",
			MatchId:       "",
		}
		s.jugadores[player.PlayerId] = jugador
		logWithTimestamp("[QueuePlayer] Nuevo jugador registrado: %s", player.PlayerId)
	} else {

		if jugador.Status != types.IDLE {
			logWithTimestamp("[QueuePlayer] Jugador ocupado, no se puede poner en cola")
			return &pb.QueuePlayerResponse{
				StatusCode:  1,
				Mensaje:     "Jugador ocupado",
				VectorClock: relojProto,
			}, nil
		}

		jugador.Status = types.IN_QUEUE
		jugador.QueueSince = time.Now()
		s.jugadores[player.PlayerId] = jugador
	}

	s.queuePlayer = append(s.queuePlayer, player.PlayerId)

	logWithTimestamp("[QueuePlayer] Jugador %s añadido a cola exitosamente", player.PlayerId)

	return &pb.QueuePlayerResponse{
		StatusCode:  0,
		Mensaje:     "Jugador en cola",
		VectorClock: relojProto,
	}, nil
}

func (s *ServidorMatchmaking) CancelQueue(ctx context.Context, req *pb.CancelQueueRequest) (*pb.CancelQueueResponse, error) {
	logWithTimestamp("[CancelQueue] Nueva solicitud para cancelar cola: %s", req.PlayerId)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	relojEntrante := fromProtoVectorClock(req.VectorClock.Entries)
	relojAnterior := s.reloj
	logWithTimestamp("[CancelQueue] Reloj anterior: %v, Reloj entrante: %v", relojAnterior, relojEntrante)

	s.reloj = fusionarMayores(s.reloj, relojEntrante)
	incrementarReloj(s.reloj, s.id)
	logWithTimestamp("[CancelQueue] Reloj fusionado: %v", s.reloj)

	relojProto := toProtoVectorClock(s.reloj)

	findIndex := -1
	for i, id := range s.queuePlayer {
		if id == req.PlayerId {
			findIndex = i
			break
		}
	}

	if findIndex != -1 {
		logWithTimestamp("[CancelQueue] Jugador %s encontrado en cola, removiéndolo", req.PlayerId)
		// Sacar jugador de la cola
		s.queuePlayer = append(s.queuePlayer[:findIndex], s.queuePlayer[findIndex+1:]...)

		// Actualizar estado del jugador
		if jugador, exists := s.jugadores[req.PlayerId]; exists {
			jugador.Status = types.IDLE
			jugador.TimeInQueue = 0
			s.jugadores[req.PlayerId] = jugador
		}

		return &pb.CancelQueueResponse{
			StatusCode:  0,
			Message:     "Jugador fuera de cola",
			VectorClock: relojProto,
		}, nil
	}

	logWithTimestamp("[CancelQueue] Jugador %s no encontrado en cola", req.PlayerId)
	return &pb.CancelQueueResponse{
		StatusCode:  1,
		Message:     "Jugador no estaba en cola",
		VectorClock: relojProto,
	}, nil
}

func (s *ServidorMatchmaking) GetPlayerStatus(ctx context.Context, player *pb.PlayerStatusRequest) (*pb.PlayerStatusResponse, error) {
	logWithTimestamp("[GetPlayerStatus] Nueva solicitud de estado de jugador: %s", player.PlayerId)

	relojEntrante := fromProtoVectorClock(player.VectorClock.Entries)
	relojAnterior := s.reloj
	logWithTimestamp("[GetPlayerStatus] Reloj anterior: %v, Reloj entrante: %v", relojAnterior, relojEntrante)

	s.reloj = fusionarMayores(s.reloj, relojEntrante)
	logWithTimestamp("[GetPlayerStatus] Reloj fusionado: %v", s.reloj)

	relojProto := toProtoVectorClock(s.reloj)

	if jugador, exists := s.jugadores[player.PlayerId]; exists {
		if jugador.Status == types.IN_MATCH {
			logWithTimestamp("[GetPlayerStatus] Jugador en partida - MatchId: %s, ServerAddress: %s", jugador.MatchId, jugador.ServerAddress)
			return &pb.PlayerStatusResponse{
				Status:        pb.PlayerStatus(types.IN_MATCH),
				MatchId:       jugador.MatchId,
				ServerAddress: jugador.ServerAddress,
				VectorClock:   relojProto,
			}, nil
		} else if jugador.Status == types.IN_QUEUE {
			logWithTimestamp("[GetPlayerStatus] Jugador en cola")
			return &pb.PlayerStatusResponse{
				Status:        pb.PlayerStatus(types.IN_QUEUE),
				MatchId:       "",
				ServerAddress: "",
				VectorClock:   relojProto,
			}, nil
		} else {
			logWithTimestamp("[GetPlayerStatus] Jugador libre")
			return &pb.PlayerStatusResponse{
				Status:        pb.PlayerStatus(types.IDLE),
				MatchId:       "",
				ServerAddress: "",
				VectorClock:   relojProto,
			}, nil
		}

	} else {
		logWithTimestamp("[GetPlayerStatus] El jugador no existe")
		return &pb.PlayerStatusResponse{
			Status:        pb.PlayerStatus(types.UNKNOWN),
			MatchId:       "",
			ServerAddress: "",
			VectorClock:   relojProto,
		}, nil
	}

}

func (s *ServidorMatchmaking) UpdateServerStatus(ctx context.Context, server *pb.ServerStatusUpdateRequest) (*pb.ServerStatusUpdateResponse, error) {
	newStatus := types.ServerStatus(server.NewStatus)
	serverId := server.ServerId
	address := server.ServerAddress
	logWithTimestamp("[UpdateServerStatus] Nueva solicitud de actualización de estado - ServerId: %s, Nuevo estado: %d", serverId, newStatus)

	relojEntrante := fromProtoVectorClock(server.VectorClock.Entries)
	relojAnterior := s.reloj
	logWithTimestamp("[UpdateServerStatus] Reloj anterior: %v, Reloj entrante: %v", relojAnterior, relojEntrante)

	s.reloj = fusionarMayores(s.reloj, relojEntrante)
	incrementarReloj(s.reloj, s.id)
	logWithTimestamp("[UpdateServerStatus] Reloj fusionado: %v", s.reloj)

	relojProto := toProtoVectorClock(s.reloj)

	if servidorActual, exists := s.servidores[serverId]; exists {
		servidorActual.Status = newStatus
		if newStatus == types.DISPONIBLE {
			logWithTimestamp("[UpdateServerStatus] Actualizando estado del servidor %s a DISPONIBLE", serverId)
			j1 := servidorActual.Player1
			j2 := servidorActual.Player2

			if j1 != "" && j2 != "" {
				jugador1 := s.jugadores[j1]
				jugador2 := s.jugadores[j2]

				jugador1.MatchId = ""
				jugador1.Status = types.IDLE

				jugador2.MatchId = ""
				jugador2.Status = types.IDLE

				s.jugadores[j1] = jugador1
				s.jugadores[j2] = jugador2

				logWithTimestamp("[UpdateServerStatus] Jugadores %s y %s liberados de partida", j1, j2)
			}

			servidorActual.CurrentMatchId = ""
			servidorActual.Player1 = ""
			servidorActual.Player2 = ""
		}
		s.servidores[serverId] = servidorActual
		logWithTimestamp("[UpdateServerStatus] Servidor %s actualizado exitosamente", serverId)
		return &pb.ServerStatusUpdateResponse{
			StatusCode:  0,
			Message:     "Estado actualizado con éxito",
			VectorClock: relojProto,
		}, nil
	} else {
		logWithTimestamp("[UpdateServerStatus] Servidor no existe, creando servidor %s...", serverId)
		fmt.Printf("Dirección del servidor: %s\n", address)
		sub := strings.Split(address, ":")
		ip := sub[0]
		port := sub[1]

		if ip == "0.0.0.0" || ip == "localhost" {
			ip = serverId
			logWithTimestamp("[UpdateServerStatus] IP Docker ajustada: %s", ip)
		}

		conn, client, err := conectarServer(ip, port)
		if err != nil {
			logWithTimestamp("[UpdateServerStatus] [ERROR] Error al conectar con el servidor")
			return &pb.ServerStatusUpdateResponse{
				StatusCode:  1,
				Message:     "Error al crear cliente del servidor",
				VectorClock: relojProto,
			}, nil
		}
		newServer := types.Server{
			Address:        address,
			Status:         types.ServerStatus(newStatus),
			CurrentMatchId: "",
			Player1:        "",
			Player2:        "",
		}

		s.servidores[serverId] = newServer
		s.clients[serverId] = client
		s.conns[serverId] = conn
		logWithTimestamp("[UpdateServerStatus] Servidor %s creado exitosamente", serverId)
		return &pb.ServerStatusUpdateResponse{
			StatusCode:  0,
			Message:     "Server Registrado",
			VectorClock: relojProto,
		}, nil
	}

}

func (s *ServidorMatchmaking) AdminUpdateServerState(ctx context.Context, adminRequest *pb.AdminServerUpdateRequest) (*pb.AdminUpdateResponse, error) {

	relojEntrante := fromProtoVectorClock(adminRequest.VectorClock.Entries)
	relojAnterior := s.reloj
	logWithTimestamp("[AdminUpdateServerState] Reloj anterior: %v, Reloj entrante: %v", relojAnterior, relojEntrante)

	s.reloj = fusionarMayores(s.reloj, relojEntrante)
	incrementarReloj(s.reloj, s.id)
	logWithTimestamp("[AdminUpdateServerState] Reloj fusionado: %v", s.reloj)

	relojProto := toProtoVectorClock(s.reloj)

	newStatus := adminRequest.NewStatus
	serverId := adminRequest.ServerId
	var estado string
	switch newStatus {
	case 0:
		estado = "DISPONIBLE"
	case 1:
		estado = "OCUPADO"
	case 2:
		estado = "CAIDO"
	case 4:
		estado = "DESCONOCIDO"
	}
	logWithTimestamp("[AdminUpdateServerState] ADMIN: solicitud de actualización de estado de servidor %s a estado %s", serverId, estado)

	if servidorActual, exists := s.servidores[serverId]; exists {
		// Intentar actualizar el estado en el servidor de juegos
		if client, clientExists := s.clients[serverId]; clientExists {
			logWithTimestamp("[AdminUpdateServerState] Enviando solicitud de actualización de estado al servidor %s", serverId)

			newReloj, err := UpdateGameServerState(ctx, client, serverId, types.ServerStatus(newStatus), s.reloj)
			if err != nil {
				logWithTimestamp("[AdminUpdateServerState] [ERROR] Error al actualizar estado en servidor de juegos %s: %v", serverId, err)
				servidorActual.Status = types.ServerStatus(newStatus)
				s.servidores[serverId] = servidorActual
				return &pb.AdminUpdateResponse{
					StatusCode:  1,
					Message:     "Estado actualizado localmente pero falló comunicación con servidor: " + err.Error(),
					VectorClock: toProtoVectorClock(newReloj),
				}, nil
			}

			s.reloj = fusionarMayores(s.reloj, newReloj)
			logWithTimestamp("[AdminUpdateServerState] Estado actualizado exitosamente en servidor de juegos %s", serverId)
		} else {
			logWithTimestamp("[AdminUpdateServerState] No hay cliente disponible para servidor %s, actualizando solo localmente", serverId)
		}

		servidorActual.Status = types.ServerStatus(newStatus)
		s.servidores[serverId] = servidorActual
		logWithTimestamp("[AdminUpdateServerState] Estado de servidor %s actualizado exitosamente", serverId)
		return &pb.AdminUpdateResponse{
			StatusCode:  0,
			Message:     "Estado actualizado con éxito",
			VectorClock: relojProto,
		}, nil
	} else {
		logWithTimestamp("[AdminUpdateServerState] El servidor %s no existe", serverId)
		return &pb.AdminUpdateResponse{
			StatusCode:  1,
			Message:     "Server no encontrado",
			VectorClock: relojProto,
		}, nil
	}
}

func (s *ServidorMatchmaking) AdminGetSystemStatus(ctx context.Context, admin *pb.AdminRequest) (*pb.SystemStatusResponse, error) {

	logWithTimestamp("[AdminGetSystemStatus] ADMIN: solicitud de estado del sistema")

	relojEntrante := fromProtoVectorClock(admin.VectorClock.Entries)
	relojAnterior := s.reloj
	logWithTimestamp("[AdminGetSystemStatus] Reloj anterior: %v, Reloj entrante: %v", relojAnterior, relojEntrante)

	s.reloj = fusionarMayores(s.reloj, relojEntrante)
	incrementarReloj(s.reloj, s.id)
	logWithTimestamp("[AdminGetSystemStatus] Reloj fusionado: %v", s.reloj)

	relojProto := toProtoVectorClock(s.reloj)

	serverList := []*pb.ServerState{}

	for sID, server := range s.servidores {
		server := &pb.ServerState{
			Id:             sID,
			Status:         pb.ServerStatus(server.Status),
			Address:        server.Address,
			CurrentMatchId: server.CurrentMatchId,
		}
		serverList = append(serverList, server)
	}

	playerList := []*pb.PlayerQueueEntry{}

	for pID, player := range s.jugadores {
		//calcular tiempo en cola
		timeInQueue := int32(0)
		if player.Status == types.IN_QUEUE {
			timeInQueue = int32(time.Since(player.QueueSince).Seconds())
		} else {
			// Si el jugador está en partida, no tiene sentido mostrar tiempo en cola
			timeInQueue = 0
		}

		player := &pb.PlayerQueueEntry{
			PlayerId:    pID,
			TimeInQueue: timeInQueue,
		}
		playerList = append(playerList, player)
	}

	logWithTimestamp("[AdminGetSystemStatus] Estado del sistema obtenido - %d servidores, %d jugadores", len(serverList), len(playerList))

	return &pb.SystemStatusResponse{
		ServerList:  serverList,
		PlayerQueue: playerList,
		VectorClock: relojProto,
	}, nil

}
