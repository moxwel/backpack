package main

import (
	"fmt"
	pb "lcp/PokemonProto"

	"context"
	"l2/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func conectarGym(ip string, port string) (*grpc.ClientConn, pb.ServicioGymClient, error) {
	fmt.Println("[conectarGym] Conectando a Gimnasio en " + ip + ":" + port)

	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", ip, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, nil, fmt.Errorf("    [ ! ] No se pudo conectar al servidor Gimnasio: %v", err)
	}

	client := pb.NewServicioGymClient(conn)

	return conn, client, nil
}

func gymAsignarCombate(gymClient pb.ServicioGymClient, entrenador1 *types.Entrenador, entrenador2 *types.Entrenador, torneoId string, region string) error {
	combate := &pb.CombateGym{
		Id: torneoId,
		Entrenador_1: &pb.Entrenador{
			Id:         entrenador1.Id,
			Nombre:     entrenador1.Nombre,
			Region:     entrenador1.Region,
			Ranking:    entrenador1.Ranking,
			Estado:     pb.EstadoEntrenador(entrenador1.Estado),
			Suspencion: entrenador1.Suspencion,
		},
		Entrenador_2: &pb.Entrenador{
			Id:         entrenador2.Id,
			Nombre:     entrenador2.Nombre,
			Region:     entrenador2.Region,
			Ranking:    entrenador2.Ranking,
			Estado:     pb.EstadoEntrenador(entrenador2.Estado),
			Suspencion: entrenador2.Suspencion,
		},
		Region: region,
	}
	_, err := gymClient.AsignarCombate(context.Background(), combate)
	return err
}
