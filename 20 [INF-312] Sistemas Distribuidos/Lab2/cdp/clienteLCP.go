package main

import (
	pb "cdp/PokemonProto"
	"context"
	"fmt"
	"l2/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func conectarLCP(ip string, port string) (*grpc.ClientConn, pb.ServicioLCPClient, error) {
	fmt.Println("[conectarLCP] Conectando a Liga Central Pokemon en " + ip + ":" + port)

	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", ip, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, nil, fmt.Errorf("    [ ! ] No se pudo conectar al servidor LCP: %v", err)
	}

	client := pb.NewServicioLCPClient(conn)

	return conn, client, nil
}

func lcpExisteEntrenador(client pb.ServicioLCPClient, idEntrenador string) (bool, error) {
	fmt.Println("[lcpExisteEntrenador] Comprobando si existe el entrenador con ID:", idEntrenador)

	req := &pb.EntrenadorID{
		Id: idEntrenador,
	}

	respuesta, err := client.ExisteEntrenador(context.Background(), req)
	if err != nil {
		return false, fmt.Errorf("    [ ! ] Error al comprobar si existe el entrenador: %v", err)
	}

	fmt.Println("    [ END ] Respuesta del servidor LCP:", respuesta.Codigo, respuesta.Mensaje)

	existe := false
	if respuesta.Codigo == 0 {
		fmt.Println("    [ OK ] Entrenador existe")
		existe = true
	} else {
		fmt.Println("    [ ! ] Entrenador no existe")
	}

	return existe, nil
}

func lcpObtenerEntrenador(client pb.ServicioLCPClient, idEntrenador string) (types.Entrenador, error) {
	fmt.Println("[lcpObtenerEntrenador] Obteniendo datos del entrenador con ID:", idEntrenador)

	req := &pb.EntrenadorID{
		Id: idEntrenador,
	}

	respuesta, err := client.ObtenerEntrenador(context.Background(), req)
	if err != nil {
		return types.Entrenador{}, fmt.Errorf("    [ ! ] Error al obtener los datos del entrenador: %v", err)
	}

	entrenador := types.Entrenador{
		Id:         respuesta.Id,
		Nombre:     respuesta.Nombre,
		Region:     respuesta.Region,
		Ranking:    respuesta.Ranking,
		Estado:     types.EstadoEntrenador(respuesta.Estado),
		Suspencion: respuesta.Suspencion,
	}

	fmt.Println("    [ OK ] Entrenador obtenido:", entrenador)

	return entrenador, nil
}
