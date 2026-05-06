package main

import (
	"context"
	"fmt"
	"l3/types"
	"os"
	"time"

	pb "servidor/MatchmakingProto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Conecta al servidor de Matchmaking y retorna la conexión y el cliente
func conectarMatchmaking(ip string, port string) (*grpc.ClientConn, pb.MatchmakingServiceClient, error) {
	fmt.Println("[conectarMatchmaking] Conectando a Matchmaking en " + ip + ":" + port)
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", ip, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("[conectarMatchmaking] [ ! ] No se pudo conectar al servidor Matchmaking: %v", err)
	}

	client := pb.NewMatchmakingServiceClient(conn)
	return conn, client, nil
}

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

// Llama a UpdateServerStatus en el servidor de Matchmaking
func matchmakingUpdateServerStatus(client pb.MatchmakingServiceClient, server_id string, server_address string, status types.ServerStatus, reloj types.VectorClockMap) (codigo int, mensaje string, clock types.VectorClockMap, err error) {
	logWithTimestamp("[matchmakingUpdateServerStatus] Llamando a UpdateServerStatus")
	logWithTimestamp("[matchmakingUpdateServerStatus] ServerId: %s, ServerAddress: %s, Status: %d", server_id, server_address, status)
	logWithTimestamp("[matchmakingUpdateServerStatus] Reloj enviado: %v", reloj)

	protoVectorClock := toProtoVectorClock(reloj)

	peticion := &pb.ServerStatusUpdateRequest{
		ServerId:      server_id,
		ServerAddress: server_address,
		NewStatus:     pb.ServerStatus(status),
		VectorClock:   protoVectorClock,
	}

	response, err := client.UpdateServerStatus(context.Background(), peticion)
	if err != nil {
		logWithTimestamp("[matchmakingUpdateServerStatus] [ERROR] Error al llamar a UpdateServerStatus: %v", err)
		return 1, "", nil, fmt.Errorf("error al llamar a UpdateServerStatus: %v", err)
	}

	resp_codigo := int(response.GetStatusCode())
	resp_mensaje := response.GetMessage()
	clock = fromProtoVectorClock(response.GetVectorClock().GetEntries())

	logWithTimestamp("[matchmakingUpdateServerStatus] Respuesta recibida - Código: %d, Mensaje: %s", resp_codigo, resp_mensaje)
	logWithTimestamp("[matchmakingUpdateServerStatus] Reloj anterior: %v, Reloj recibido: %v", reloj, clock)

	if resp_codigo != 0 {
		logWithTimestamp("[matchmakingUpdateServerStatus] [ERROR] Error en UpdateServerStatus: Codigo %d, Mensaje: %s", resp_codigo, resp_mensaje)
		return resp_codigo, resp_mensaje, clock, fmt.Errorf("codigo de respuesta: %d, Mensaje: %s", resp_codigo, resp_mensaje)
	}

	//TODO: reintentar si el servidor no responde

	return 0, "Estado del servidor actualizado correctamente", clock, nil
}
