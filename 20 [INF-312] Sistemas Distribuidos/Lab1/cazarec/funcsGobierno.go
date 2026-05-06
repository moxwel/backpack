package main

import (
	"context"
	"fmt"
	"time"

	pb "cazarec/servicioMensajes"

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

func gobBasic(cliente pb.GobiernoServiceClient, mensaje string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	respuesta, err := cliente.Basic(ctx, &pb.BasicRequest{Msg: mensaje})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Respuesta del Gobierno Mundial:", respuesta.GetMsg())
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

func gobActualizarReputacion(cliente pb.GobiernoServiceClient, id_pirata int32, estado string, subMundo bool) int32 {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	respuesta, err := cliente.ResultadoEntrega(ctx, &pb.ActualizarReputacionRequest{IdPirata: id_pirata, Estado: estado, Submundo: subMundo})
	if err != nil {
		fmt.Println(err)
	}

	return respuesta.GetReputacion()
}

func gobMarcarCaptura(cliente pb.GobiernoServiceClient, id_pirata int32) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := cliente.MarcarCaptura(ctx, &pb.PirataCapturado{IdPirata: id_pirata})

	if err != nil {
		fmt.Println(err)
	}
}
