package main

import (
	"context"
	"fmt"
	"l3/types"
	"math/rand"
	"sync"
	"time"

	pb "servidor/MatchmakingProto"
)

type ServidorJuegos struct {
	pb.UnimplementedGameServerServiceServer
	server_id        string
	server_address   string
	current_match_id string
	current_player_1 string
	current_player_2 string
	status           types.ServerStatus
	matchmakerClient pb.MatchmakingServiceClient
	clock            types.VectorClockMap
	resultsMutex     sync.Mutex
	resultsCond      *sync.Cond
	matchResults     map[string]types.MatchResult
}

func simularPartida(s *ServidorJuegos, match_id, player1, player2 string) {
	duracion := 10 + (rand.Intn(11)) // 10 a 20 segundos

	logWithTimestamp("[simularPartida] Iniciando simulación de partida '%s' entre '%s' y '%s' en servidor '%s'. Duración: %d segundos", match_id, player1, player2, s.server_id, duracion)

	time.Sleep(time.Duration(duracion) * time.Second)

	winner, loser := player1, player2
	if rand.Intn(2) == 1 {
		winner, loser = player2, player1
	}

	logWithTimestamp("[simularPartida] Partida '%s' finalizada - Ganador: '%s', Perdedor: '%s'", match_id, winner, loser)

	s.resultsMutex.Lock()
	if s.resultsCond == nil {
		s.resultsCond = sync.NewCond(&s.resultsMutex)
	}
	s.matchResults[match_id] = types.MatchResult{
		Winner: winner,
		Loser:  loser,
	}
	s.resultsCond.Broadcast()
	logWithTimestamp("[simularPartida] Resultado de partida '%s' almacenado y notificado", match_id)
	s.resultsMutex.Unlock()

	s.current_match_id = ""
	s.current_player_1 = ""
	s.current_player_2 = ""
	s.status = types.DISPONIBLE

	logWithTimestamp("[simularPartida] Notificando a Matchmaker el cambio de estado a DISPONIBLE")
	logWithTimestamp("[simularPartida] Reloj antes de notificación: %v", s.clock)

	_, _, reloj, err := matchmakingUpdateServerStatus(s.matchmakerClient, s.server_id, s.server_address, s.status, s.clock)
	if err != nil {
		logWithTimestamp("[simularPartida] [ERROR] Error al actualizar estado del servidor: %v", err)
	} else {
		relojAnterior := s.clock
		s.clock = fusionarMayores(s.clock, reloj)
		logWithTimestamp("[simularPartida] Estado actualizado exitosamente")
		logWithTimestamp("[simularPartida] Reloj anterior: %v, Reloj recibido: %v, Reloj fusionado: %v", relojAnterior, reloj, s.clock)
	}
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

func (s *ServidorJuegos) AssignMatch(ctx context.Context, req *pb.AssignMatchRequest) (*pb.AssignMatchResponse, error) {
	match_id := req.GetId()
	player_id_1 := req.GetPlayerId1()
	player_id_2 := req.GetPlayerId2()

	logWithTimestamp("[AssignMatch] Solicitud de asignación recibida - MatchID: %s, Jugador1: %s, Jugador2: %s", match_id, player_id_1, player_id_2)

	relojEntrante := fromProtoVectorClock(req.VectorClock.Entries)
	relojAnterior := s.clock
	logWithTimestamp("[AssignMatch] Reloj anterior: %v, Reloj entrante: %v", relojAnterior, relojEntrante)

	s.clock = fusionarMayores(s.clock, relojEntrante)
	incrementarReloj(s.clock, s.server_id)
	logWithTimestamp("[AssignMatch] Reloj fusionado: %v", s.clock)

	relojProto := toProtoVectorClock(s.clock)

	// Validaciones
	if s.status == types.OCUPADO {
		logWithTimestamp("[AssignMatch] Rechazo: servidor ocupado")
		return &pb.AssignMatchResponse{
			StatusCode:  1,
			Message:     "El servidor no está disponible para asignar partidas",
			VectorClock: relojProto,
		}, nil
	} else if s.status == types.CAIDO {
		logWithTimestamp("[AssignMatch] Rechazo: servidor caído")
		time.Sleep(15 * time.Second)
		return &pb.AssignMatchResponse{
			StatusCode:  1,
			Message:     "El servidor no está disponible para asignar partidas",
			VectorClock: relojProto,
		}, nil
	}
	if match_id == "" {
		logWithTimestamp("[AssignMatch] Rechazo: ID de partida vacío")
		return &pb.AssignMatchResponse{
			StatusCode:  1,
			Message:     "ID de partida no proporcionado",
			VectorClock: relojProto,
		}, nil
	}
	if player_id_1 == "" || player_id_2 == "" {
		logWithTimestamp("[AssignMatch] Rechazo: IDs de jugadores vacíos")
		return &pb.AssignMatchResponse{
			StatusCode:  1,
			Message:     "IDs de jugadores no proporcionados",
			VectorClock: relojProto,
		}, nil
	}
	if player_id_1 == player_id_2 {
		logWithTimestamp("[AssignMatch] Rechazo: jugadores idénticos")
		return &pb.AssignMatchResponse{
			StatusCode:  1,
			Message:     "Los jugadores no pueden ser el mismo",
			VectorClock: relojProto,
		}, nil
	}

	// Asignar partida
	logWithTimestamp("[AssignMatch] Asignando partida '%s' entre '%s' y '%s'", match_id, player_id_1, player_id_2)

	s.current_match_id = match_id
	s.status = types.OCUPADO
	s.current_player_1 = player_id_1
	s.current_player_2 = player_id_2

	logWithTimestamp("[AssignMatch] Notificando cambio de estado a OCUPADO")
	logWithTimestamp("[AssignMatch] Reloj antes de notificación: %v", s.clock)

	cod_resp, msg, reloj, err := matchmakingUpdateServerStatus(s.matchmakerClient, s.server_id, s.server_address, s.status, s.clock)

	if err == nil {
		relojAnteriorNotif := s.clock
		s.clock = fusionarMayores(s.clock, reloj)
		logWithTimestamp("[AssignMatch] Estado notificado exitosamente")
		logWithTimestamp("[AssignMatch] Reloj anterior: %v, Reloj recibido: %v, Reloj fusionado: %v", relojAnteriorNotif, reloj, s.clock)
	}

	go simularPartida(s, match_id, player_id_1, player_id_2)

	incrementarReloj(s.clock, s.server_id)
	logWithTimestamp("[AssignMatch] Simulación iniciada, reloj incrementado: %v", s.clock)

	// Verificar respuesta
	if err != nil {
		logWithTimestamp("[AssignMatch] [ERROR] Error en UpdateServerStatus: %v", err)
		return &pb.AssignMatchResponse{
			StatusCode:  1,
			Message:     "Error al ejecutar UpdateServerStatus: " + err.Error(),
			VectorClock: relojProto,
		}, nil
	}
	if cod_resp != 0 {
		logWithTimestamp("[AssignMatch] [ERROR] Código de respuesta no exitoso: %d", cod_resp)
		return &pb.AssignMatchResponse{
			StatusCode:  int32(cod_resp),
			Message:     msg,
			VectorClock: relojProto,
		}, nil
	}

	logWithTimestamp("[AssignMatch] Partida asignada exitosamente")
	return &pb.AssignMatchResponse{
		StatusCode:  0,
		Message:     "Partida asignada correctamente",
		VectorClock: relojProto,
	}, nil
}

func (s *ServidorJuegos) PingServer(ctx context.Context, req *pb.PingServerRequest) (*pb.PingServerResponse, error) {
	logWithTimestamp("[PingServer] Ping recibido")

	if s.status == types.CAIDO {
		logWithTimestamp("[PingServer] Respondiendo con error - servidor marcado como caído")
		time.Sleep(15 * time.Second)
		return &pb.PingServerResponse{
			StatusCode:  1,
			Message:     "Servidor está caído, no se puede responder al ping",
			VectorClock: toProtoVectorClock(s.clock),
		}, nil
	}

	relojEntrante := fromProtoVectorClock(req.VectorClock.Entries)
	relojAnterior := s.clock
	logWithTimestamp("[PingServer] Reloj anterior: %v, Reloj entrante: %v", relojAnterior, relojEntrante)

	s.clock = fusionarMayores(s.clock, relojEntrante)
	incrementarReloj(s.clock, s.server_id)
	logWithTimestamp("[PingServer] Reloj fusionado: %v", s.clock)

	relojProto := toProtoVectorClock(s.clock)

	mensaje := fmt.Sprintf("Servidor '%s' está activo. Estado: %d", s.server_id, s.status)
	logWithTimestamp("[PingServer] Respondiendo exitosamente - %s", mensaje)

	return &pb.PingServerResponse{
		StatusCode:  0,
		Message:     mensaje,
		VectorClock: relojProto,
	}, nil
}

func (s *ServidorJuegos) UpdateGameServerState(ctx context.Context, req *pb.UpdateGameServerStateRequest) (*pb.UpdateGameServerStateResponse, error) {
	newStatus := req.GetNewStatus()
	serverId := req.GetServerId()

	logWithTimestamp("[UpdateGameServerState] Solicitud de actualización de estado recibida - Nuevo estado: %d", newStatus)

	relojEntrante := fromProtoVectorClock(req.VectorClock.Entries)
	relojAnterior := s.clock
	logWithTimestamp("[UpdateGameServerState] Reloj anterior: %v, Reloj entrante: %v", relojAnterior, relojEntrante)

	s.clock = fusionarMayores(s.clock, relojEntrante)
	incrementarReloj(s.clock, s.server_id)
	logWithTimestamp("[UpdateGameServerState] Reloj fusionado: %v", s.clock)

	relojProto := toProtoVectorClock(s.clock)

	// Validar que el serverId coincida
	if serverId != s.server_id {
		logWithTimestamp("[UpdateGameServerState] [ERROR] ID de servidor no coincide - Esperado: %s, Recibido: %s", s.server_id, serverId)
		return &pb.UpdateGameServerStateResponse{
			StatusCode:  1,
			Message:     "ID de servidor no coincide",
			VectorClock: relojProto,
		}, nil
	}

	// Actualizar el estado del servidor
	oldStatus := s.status
	s.status = types.ServerStatus(newStatus)

	logWithTimestamp("[UpdateGameServerState] Estado actualizado de %d a %d", oldStatus, s.status)

	// Si se está reactivando un servidor caído, limpiar estado de partida actual
	if oldStatus == types.CAIDO && s.status == types.DISPONIBLE {
		s.current_match_id = ""
		s.current_player_1 = ""
		s.current_player_2 = ""
		logWithTimestamp("[UpdateGameServerState] Servidor reactivado - estado de partida limpiado")
	}

	return &pb.UpdateGameServerStateResponse{
		StatusCode:  0,
		Message:     "Estado actualizado correctamente",
		VectorClock: relojProto,
	}, nil
}

func (s *ServidorJuegos) GetMatchResult(ctx context.Context, req *pb.GetMatchResultRequest) (*pb.GetMatchResultResponse, error) {
	matchId := req.GetMatchId()
	logWithTimestamp("[GetMatchResult] Solicitud de resultado para partida '%s' - esperando finalización...", matchId)

	relojEntrante := fromProtoVectorClock(req.VectorClock.Entries)
	relojAnterior := s.clock
	logWithTimestamp("[GetMatchResult] Reloj anterior: %v, Reloj entrante: %v", relojAnterior, relojEntrante)

	s.clock = fusionarMayores(s.clock, relojEntrante)
	incrementarReloj(s.clock, s.server_id)
	logWithTimestamp("[GetMatchResult] Reloj fusionado: %v", s.clock)

	relojProto := toProtoVectorClock(s.clock)

	s.resultsMutex.Lock()
	if s.resultsCond == nil {
		s.resultsCond = sync.NewCond(&s.resultsMutex)
	}

	result, found := s.matchResults[matchId]
	for !found {
		logWithTimestamp("[GetMatchResult] Resultado no disponible aún, esperando...")
		s.resultsCond.Wait()
		result, found = s.matchResults[matchId]
	}
	s.resultsMutex.Unlock()

	logWithTimestamp("[GetMatchResult] Resultado obtenido para partida '%s' - Ganador: %s, Perdedor: %s", matchId, result.Winner, result.Loser)

	return &pb.GetMatchResultResponse{
		StatusCode:  0,
		Message:     "Resultado obtenido.",
		WinnerId:    result.Winner,
		LoserId:     result.Loser,
		VectorClock: relojProto,
	}, nil
}
