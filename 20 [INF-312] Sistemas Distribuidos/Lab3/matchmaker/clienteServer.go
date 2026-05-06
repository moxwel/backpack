package main

import (
	"fmt"
	"l3/types"
	pb "matchmaker/MatchmakingProto"

	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func conectarServer(ip string, port string) (*grpc.ClientConn, pb.GameServerServiceClient, error) {
	logWithTimestamp("[--!--]    Conectando a Servidor en %s: %s", ip, port)
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", ip, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, nil, fmt.Errorf("[ERROR]    No se pudo conectar al Servidor : %v", err)
	}

	client := pb.NewGameServerServiceClient(conn)

	return conn, client, nil
}

func AssignMatch(ctx context.Context, serverClient pb.GameServerServiceClient, serverId string, matchId string, player1 string, player2 string, reloj types.VectorClockMap) (types.VectorClockMap, error) {
	// Crear contexto con timeout de 10 segundos
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	logWithTimestamp("[AssignMatch] Asignando partida '%s' entre '%s' y '%s' en servidor '%s'", matchId, player1, player2, serverId)
	logWithTimestamp("[AssignMatch] Reloj enviado: %v", reloj)

	protoVectorClock := toProtoVectorClock(reloj)
	match := &pb.AssignMatchRequest{
		Id:          matchId,
		PlayerId1:   player1,
		PlayerId2:   player2,
		VectorClock: protoVectorClock,
	}

	res, err := serverClient.AssignMatch(timeoutCtx, match)
	if err != nil {
		logWithTimestamp("[AssignMatch] [ERROR] Error al enviar mensaje: %v", err)
		return reloj, err
	}

	clock := fromProtoVectorClock(res.VectorClock.Entries)
	relojFusionado := fusionarMayores(reloj, clock)

	logWithTimestamp("[AssignMatch] Respuesta recibida - Código: %d, Mensaje: %s", res.StatusCode, res.Message)
	logWithTimestamp("[AssignMatch] Reloj anterior: %v, Reloj recibido: %v, Reloj fusionado: %v", reloj, clock, relojFusionado)

	return relojFusionado, err
}

func PingServer(ctx context.Context, serverClient pb.GameServerServiceClient, serverId string, server *ServidorMatchmaking, reloj types.VectorClockMap) (types.VectorClockMap, error) {
	// Crear contexto con timeout de 10 segundos
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	logWithTimestamp("[PingServer] Enviando ping a servidor %s", serverId)
	logWithTimestamp("[PingServer] Reloj enviado: %v", reloj)

	protoVectorClock := toProtoVectorClock(reloj)
	res, err := serverClient.PingServer(timeoutCtx, &pb.PingServerRequest{
		VectorClock: protoVectorClock,
	})
	if err != nil || res.StatusCode != 0 {
		logWithTimestamp("[PingServer] Servidor %s marcado como caído", serverId)
		s := server.servidores[serverId]
		s.Status = types.CAIDO
		server.servidores[serverId] = s
		return reloj, err
	}

	clock := fromProtoVectorClock(res.VectorClock.Entries)
	relojFusionado := fusionarMayores(reloj, clock)

	logWithTimestamp("[PingServer] Servidor %s funcionando correctamente", serverId)
	logWithTimestamp("[PingServer] Reloj anterior: %v, Reloj recibido: %v, Reloj fusionado: %v", reloj, clock, relojFusionado)

	return relojFusionado, err
}

func UpdateGameServerState(ctx context.Context, serverClient pb.GameServerServiceClient, serverId string, newStatus types.ServerStatus, reloj types.VectorClockMap) (types.VectorClockMap, error) {
	// Crear contexto con timeout de 10 segundos
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	logWithTimestamp("[UpdateGameServerState] Actualizando estado de servidor %s a %d", serverId, newStatus)
	logWithTimestamp("[UpdateGameServerState] Reloj enviado: %v", reloj)

	protoVectorClock := toProtoVectorClock(reloj)
	req := &pb.UpdateGameServerStateRequest{
		ServerId:    serverId,
		NewStatus:   pb.ServerStatus(newStatus),
		VectorClock: protoVectorClock,
	}

	res, err := serverClient.UpdateGameServerState(timeoutCtx, req)
	if err != nil {
		logWithTimestamp("[UpdateGameServerState] [ERROR] Error al actualizar estado del servidor %s: %v", serverId, err)
		return reloj, err
	}

	clock := fromProtoVectorClock(res.VectorClock.Entries)
	relojFusionado := fusionarMayores(reloj, clock)

	logWithTimestamp("[UpdateGameServerState] Estado del servidor %s actualizado - Código: %d, Mensaje: %s", serverId, res.StatusCode, res.Message)
	logWithTimestamp("[UpdateGameServerState] Reloj anterior: %v, Reloj recibido: %v, Reloj fusionado: %v", reloj, clock, relojFusionado)

	return relojFusionado, err
}
