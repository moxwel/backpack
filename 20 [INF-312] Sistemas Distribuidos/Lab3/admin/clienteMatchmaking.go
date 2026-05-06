package main

import (
	pb "admin/MatchmakingProto"
	"context"
	"fmt"
	"l3/types"

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

func AdminGetSystemStatus(serverClient pb.MatchmakingServiceClient, reloj types.VectorClockMap) (types.SystemStatusResponse, error) {

	logWithTimestamp("[AdminGetSystemStatus] Solicitando estado del sistema")
	logWithTimestamp("[AdminGetSystemStatus] Reloj enviado: %v", reloj)

	protoVectorClock := toProtoVectorClock(reloj)

	res, err := serverClient.AdminGetSystemStatus(context.Background(), &pb.AdminRequest{
		VectorClock: protoVectorClock,
	})

	serverList := []types.ServerState{}
	playerQueue := []types.PlayerQueueEntry{}

	if err != nil {
		logWithTimestamp("[AdminGetSystemStatus] [ERROR] Error al obtener status: %v", err)
		return types.SystemStatusResponse{
			ServerList:  serverList,
			PlayerQueue: playerQueue,
			VectorClock: reloj,
		}, nil
	}

	relojRecibido := fromProtoVectorClock(res.VectorClock.Entries)
	logWithTimestamp("[AdminGetSystemStatus] Estado del sistema obtenido")
	logWithTimestamp("[AdminGetSystemStatus] Reloj anterior: %v, Reloj recibido: %v", reloj, relojRecibido)

	for _, serv := range res.ServerList {
		act := types.ServerState{
			Id:             serv.Id,
			Status:         types.ServerStatus(serv.Status),
			Address:        serv.Address,
			CurrentMatchId: serv.CurrentMatchId,
		}
		serverList = append(serverList, act)
	}

	for _, pl := range res.PlayerQueue {
		act := types.PlayerQueueEntry{
			PlayerId:    pl.PlayerId,
			TimeInQueue: pl.TimeInQueue,
		}
		playerQueue = append(playerQueue, act)
	}

	return types.SystemStatusResponse{
		ServerList:  serverList,
		PlayerQueue: playerQueue,
		VectorClock: relojRecibido,
	}, nil

}

func AdminUpdateServerState(serverClient pb.MatchmakingServiceClient, serverId string, newStatus types.ServerStatus, reloj types.VectorClockMap) (types.AdminUpdateResponse, error) {

	logWithTimestamp("[AdminUpdateServerState] Actualizando estado de servidor %s a %d", serverId, newStatus)
	logWithTimestamp("[AdminUpdateServerState] Reloj enviado: %v", reloj)

	protoVectorClock := toProtoVectorClock(reloj)

	req := &pb.AdminServerUpdateRequest{
		ServerId:    serverId,
		NewStatus:   pb.ServerStatus(newStatus),
		VectorClock: protoVectorClock,
	}

	res, err := serverClient.AdminUpdateServerState(context.Background(), req)
	if err != nil {
		logWithTimestamp("[AdminUpdateServerState] [ERROR] Error al conectar con el servidor de Matchmaking")
		return types.AdminUpdateResponse{}, nil
	}

	relojRecibido := fromProtoVectorClock(res.VectorClock.Entries)
	logWithTimestamp("[AdminUpdateServerState] Respuesta recibida - Código: %d, Mensaje: %s", res.StatusCode, res.Message)
	logWithTimestamp("[AdminUpdateServerState] Reloj anterior: %v, Reloj recibido: %v", reloj, relojRecibido)

	return types.AdminUpdateResponse{
		StatusCode:  res.StatusCode,
		Message:     res.Message,
		VectorClock: relojRecibido,
	}, nil
}
