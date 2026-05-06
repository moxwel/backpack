package main

import (
	"context"
	"fmt"
	"time"

	pb "gobmun/servicioMensajes"

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

func marActividadIlegal(cliente pb.MarinaServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := cliente.ActividadIlegal(ctx, &pb.Empty{})

	if err != nil {
		fmt.Println(err)
	}

}
