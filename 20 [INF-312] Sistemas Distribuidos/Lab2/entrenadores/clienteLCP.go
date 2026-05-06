package main

import (
	"context"
	pb "entrenadores/PokemonProto"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"l2/types"
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

func lcpRegistrarEntrenador(cliente pb.ServicioLCPClient, nuevo_entr types.Entrenador) error {
	fmt.Println("[lcpRegistrarEntrenador] Registrando entrenador en LCP.\n    Datos:", nuevo_entr)

	inscripcion := &pb.Entrenador{
		Id:         nuevo_entr.Id,
		Nombre:     nuevo_entr.Nombre,
		Region:     nuevo_entr.Region,
		Ranking:    nuevo_entr.Ranking,
		Estado:     pb.EstadoEntrenador(nuevo_entr.Estado),
		Suspencion: nuevo_entr.Suspencion,
	}

	fmt.Println("    [...] Ejecutando gRPC con los datos:", inscripcion)

	respuesta, err := cliente.RegistrarEntrenador(context.Background(), inscripcion)
	if err != nil {
		return fmt.Errorf("    [ ! ] Error al ejecutar gRPC: %v", err)
	}

	if respuesta.Codigo != 0 {
		return fmt.Errorf("    [ ! ] Error al inscribir al entrenador: %s", respuesta.Mensaje)
	}

	fmt.Println("    [ END ] Respuesta del servidor LCP:", respuesta.Codigo, respuesta.Mensaje)
	return nil
}

func lcpObtenerTorneos(cliente pb.ServicioLCPClient) ([]types.Torneo, error) {
	// fmt.Println("[lcpObtenerTorneos] Solicitando torneos disponibles al LCP...")
	resp, err := cliente.ObtenerTorneos(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("[ ! ] Error al obtener torneos: %v", err)
	}

	torneos := make([]types.Torneo, 0, len(resp.Torneos))
	for _, t := range resp.Torneos {
		torneos = append(torneos, types.Torneo{
			Id:     t.Id,
			Region: t.Region,
			Estado: types.EstadoTorneo(t.Estado),
		})
	}
	return torneos, nil
}

func lcpObtenerTorneosRegion(cliente pb.ServicioLCPClient, region string) ([]types.Torneo, error) {
	// fmt.Println("[lcpObtenerTorneosRegion] Solicitando torneos de la región", region)
	resp, err := cliente.ObtenerTorneosRegion(context.Background(), &pb.RegionString{Region: region})
	if err != nil {
		return nil, fmt.Errorf("[ ! ] Error al obtener torneos de la región %s: %v", region, err)
	}

	torneos := make([]types.Torneo, 0, len(resp.Torneos))
	for _, t := range resp.Torneos {
		if t.Region == region && t.Estado == pb.EstadoTorneo(pb.EstadoTorneo_TORNEO_ACTIVO) {
			torneos = append(torneos, types.Torneo{
				Id:     t.Id,
				Region: t.Region,
				Estado: types.EstadoTorneo(t.Estado),
			})
		}
	}
	return torneos, nil
}

func lcpInscribirTorneo(cliente pb.ServicioLCPClient, entrenador *types.Entrenador, torneoID string) error {
	inscripcion := &pb.InscripcionTorneo{
		EntrenadorId: entrenador.Id,
		TorneoId:     torneoID,
	}
	respuesta, err := cliente.InscribirEnTorneo(context.Background(), inscripcion)
	if err != nil {
		return fmt.Errorf("[ ! ] Error al inscribir en torneo: %v", err)
	}
	fmt.Printf("[lcpInscribirTorneo] Respuesta: %d - %s\n", respuesta.Codigo, respuesta.Mensaje)
	if entrenador.Estado == types.ENTRENADOR_SUSPENDIDO && entrenador.Suspencion > 0 {
		entrenador.Suspencion -= 1
		fmt.Printf("[lcpInscribirTorneo] Entrenador suspendido, se reduce suspensión a %d.\n", entrenador.Suspencion)
		if entrenador.Suspencion == 0 {
			entrenador.Estado = types.ENTRENADOR_ACTIVO
			fmt.Println("[lcpInscribirTorneo] Suspensión finalizada. Entrenador reactivado.")
		}
	}
	if respuesta.Codigo != 0 {
		return fmt.Errorf(respuesta.Mensaje)
	}
	return nil
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
