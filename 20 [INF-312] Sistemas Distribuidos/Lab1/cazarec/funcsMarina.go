package main

import (
	"context"
	"fmt"
	"time"

	pb "cazarec/servicioMensajes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func marConectar(ip string, port string) (*grpc.ClientConn, pb.MarinaServiceClient) {
	fmt.Println("Conectando a Marina en " + ip + ":" + port)

	conexion, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", ip, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	cliente := pb.NewMarinaServiceClient(conexion)

	if err != nil {
		panic(err)
	}

	return conexion, cliente
}

func marBasic(cliente pb.MarinaServiceClient, mensaje string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	respuesta, err := cliente.Basic(ctx, &pb.BasicRequest{Msg: mensaje})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Respuesta del Gobierno Mundial:", respuesta.GetMsg())
}

func marEntregaPirata(cliente pb.MarinaServiceClient, balance int32, pirata *pb.Pirata, reputacion int32) (string, int32) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	respuesta, err := cliente.EntregarPirata(ctx, &pb.EntregaRequest{Pirata: pirata, Reputacion: reputacion, Balance: balance})

	if err != nil {
		fmt.Println(err)
	}

	return respuesta.GetEstado(), respuesta.GetPago()
}
