package main

import (
	"context"
	"fmt"
	"time"

	pb "marine/servicioMensajes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func gobConectar(ip string, port string) (*grpc.ClientConn, pb.GobiernoServiceClient) {
	fmt.Println("Conectando a Gobierno Mundial en " + ip + ":" + port)

	conexion, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", ip, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	cliente := pb.NewGobiernoServiceClient(conexion)

	if err != nil {
		panic(err)
	}

	return conexion, cliente
}

func gobListaPiratas(cliente pb.GobiernoServiceClient) []*pb.Pirata {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	respuesta, err := cliente.ObtenerListaPiratas(ctx, &pb.Empty{})

	if err != nil {
		fmt.Println(err)
	}
	return respuesta.GetPiratas()
}

func gobListaPiratasTodos(cliente pb.GobiernoServiceClient) []*pb.Pirata {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	respuesta, err := cliente.ObtenerListaPiratasTodos(ctx, &pb.Empty{})

	if err != nil {
		fmt.Println(err)
	}
	return respuesta.GetPiratas()
}
