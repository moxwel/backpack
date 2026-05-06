package main

import (
	"context"
	"fmt"
	pb "jugador/MatchmakingProto"
	"l3/types"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func conectarGameServer(ip string, port string) (*grpc.ClientConn, pb.GameServerServiceClient, error) {
	logWithTimestamp("[--!--]    Conectando a Game Server en %s:%s", ip, port)
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", ip, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, nil, fmt.Errorf("[ERROR]    No se pudo conectar al Game Server: %v", err)
	}

	client := pb.NewGameServerServiceClient(conn)
	return conn, client, nil
}

func GetMatchResult(client pb.GameServerServiceClient, matchId string, reloj types.VectorClockMap) (*pb.GetMatchResultResponse, error) {
	logWithTimestamp("[GetMatchResult] Solicitando resultado de partida %s", matchId)
	logWithTimestamp("[GetMatchResult] Reloj enviado: %v", reloj)

	solicitud := &pb.GetMatchResultRequest{
		MatchId:     matchId,
		VectorClock: toProtoVectorClock(reloj),
	}

	//context con timeout de 1 minuto
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	res, err := client.GetMatchResult(ctx, solicitud)
	if err != nil {
		logWithTimestamp("[GetMatchResult] [ERROR] Error al conectar con GameServer para obtener resultado")
		return res, err
	}

	relojRecibido := fromProtoVectorClock(res.VectorClock.Entries)
	logWithTimestamp("[GetMatchResult] Resultado obtenido - Ganador: %s, Perdedor: %s", res.WinnerId, res.LoserId)
	logWithTimestamp("[GetMatchResult] Reloj anterior: %v, Reloj recibido: %v", reloj, relojRecibido)

	return res, nil
}
