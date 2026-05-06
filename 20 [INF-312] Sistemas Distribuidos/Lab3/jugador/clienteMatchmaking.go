package main

import (
	"context"
	"fmt"
	"l3/types"
	"time"

	pb "jugador/MatchmakingProto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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

// Conecta al servidor de Matchmaking y retorna la conexión y el cliente
func conectarServer(ip string, port string) (*grpc.ClientConn, pb.MatchmakingServiceClient, error) {
	logWithTimestamp("[--!--]    Conectando a Servidor en %s: %s", ip, port)
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", ip, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, nil, fmt.Errorf("[ERROR]    No se pudo conectar al Servidor : %v", err)
	}

	client := pb.NewMatchmakingServiceClient(conn)

	return conn, client, nil
}

func QueuePlayer(client pb.MatchmakingServiceClient, player string, reloj types.VectorClockMap) (int32, types.VectorClockMap, error) {
	logWithTimestamp("[QueuePlayer] Solicitando cola para jugador %s", player)
	logWithTimestamp("[QueuePlayer] Reloj enviado: %v", reloj)

	solicitud := &pb.PlayerInfoRequest{
		PlayerId:    player,
		VectorClock: toProtoVectorClock(reloj),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := client.QueuePlayer(ctx, solicitud)
	if err != nil {
		logWithTimestamp("[QueuePlayer] [ERROR] Error al conectar con Matchmaking")
		return 1, reloj, err
	}

	relojRecibido := fromProtoVectorClock(res.VectorClock.Entries)
	logWithTimestamp("[QueuePlayer] Respuesta recibida - Código: %d", res.StatusCode)
	logWithTimestamp("[QueuePlayer] Reloj anterior: %v, Reloj recibido: %v", reloj, relojRecibido)
	logWithTimestamp("[QueuePlayer] Jugador %s procesado en cola", player)

	return res.StatusCode, relojRecibido, nil
}

func CancelQueue(client pb.MatchmakingServiceClient, player string, reloj types.VectorClockMap) (*pb.CancelQueueResponse, error) {
	logWithTimestamp("[CancelQueue] Solicitando cancelar cola para jugador %s", player)
	logWithTimestamp("[CancelQueue] Reloj enviado: %v", reloj)

	solicitud := &pb.CancelQueueRequest{
		PlayerId:    player,
		VectorClock: toProtoVectorClock(reloj),
	}
	res, err := client.CancelQueue(context.Background(), solicitud)
	if err != nil {
		logWithTimestamp("[CancelQueue] [ERROR] Error al conectar con Matchmaking para cancelar cola")
		return res, err
	}

	relojRecibido := fromProtoVectorClock(res.VectorClock.Entries)
	logWithTimestamp("[CancelQueue] Respuesta recibida - Código: %d, Mensaje: %s", res.StatusCode, res.Message)
	logWithTimestamp("[CancelQueue] Reloj anterior: %v, Reloj recibido: %v", reloj, relojRecibido)

	return res, nil
}

func GetPlayerStatus(client pb.MatchmakingServiceClient, player string, reloj types.VectorClockMap) (types.PlayerStatusResponse, error) {

	solicitud := &pb.PlayerStatusRequest{
		PlayerId:    player,
		VectorClock: toProtoVectorClock(reloj),
	}
	res, err := client.GetPlayerStatus(context.Background(), solicitud)

	if err != nil {
		logWithTimestamp("[GetPlayerStatus] [ERROR] Error al conectar con Matchmaking")
		vclock := make(types.VectorClockMap)
		if res != nil {
			vclock = fromProtoVectorClock(res.VectorClock.Entries)
		}
		return types.PlayerStatusResponse{
			Status:        types.PlayerStatus(res.Status),
			MatchId:       res.MatchId,
			ServerAddress: res.ServerAddress,
			VectorClock:   vclock,
		}, err
	}

	relojRecibido := fromProtoVectorClock(res.VectorClock.Entries)

	return types.PlayerStatusResponse{
		Status:        types.PlayerStatus(res.Status),
		MatchId:       res.MatchId,
		ServerAddress: res.ServerAddress,
		VectorClock:   relojRecibido,
	}, nil
}
