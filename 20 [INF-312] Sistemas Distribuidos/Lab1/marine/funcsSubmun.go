package main

import (
	"context"
	"fmt"
	"time"

	pb "marine/servicioMensajes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func subConectar(ip string, port string) (*grpc.ClientConn, pb.SubMundoServiceClient) {
	fmt.Println("Conectando a Submundo en " + ip + ":" + port)

	conexion, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", ip, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	cliente := pb.NewSubMundoServiceClient(conexion)

	if err != nil {
		panic(err)
	}

	return conexion, cliente
}

func subBasic(cliente pb.SubMundoServiceClient, mensaje string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	respuesta, err := cliente.Basic(ctx, &pb.BasicRequest{Msg: mensaje})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Respuesta del Gobierno Mundial:", respuesta.GetMsg())
}

func subActivarRedadas(cliente pb.SubMundoServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := cliente.ActivarRedadas(ctx, &pb.Empty{})

	if err != nil {
		fmt.Println(err)
	}
}
