package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	pb "gobmun/servicioMensajes"

	"google.golang.org/grpc"
)

type servidor struct {
	pb.UnimplementedGobiernoServiceServer
	listaPiratas     []*pb.Pirata
	balance          int
	mar_ip           string
	mar_port         string
	redadas          bool
	piratasRestantes int
}

func cargarPiratas(arch string) ([]*pb.Pirata, int, error) {
	file, err := os.Open(arch)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	piratas_capturados := 0
	var listaPiratas []*pb.Pirata
	scanner := bufio.NewScanner(file)
	scanner.Scan() // Salta la primera línea
	for scanner.Scan() {
		linea := scanner.Text()
		partes := strings.Split(linea, ",")
		// fmt.Println("Partes:", partes)
		// fmt.Println("Cantidad de partes:", len(partes))
		if len(partes) != 5 {
			continue
		}
		id, _ := strconv.Atoi(partes[0])
		recompensa, _ := strconv.Atoi(partes[2])
		if partes[4] == "Capturado" {
			piratas_capturados++
		}
		pirata := &pb.Pirata{
			Id:           int32(id),
			Nombre:       partes[1],
			Recompensa:   int32(recompensa),
			Peligrosidad: partes[3],
			Estado:       partes[4],
		}
		listaPiratas = append(listaPiratas, pirata)
	}
	//printear lista de piratas
	fmt.Println("    Lista de piratas cargada.")
	for _, pirata := range listaPiratas {
		fmt.Printf("        ID: %d, Nombre: %s, Recompensa: %d, Peligrosidad: %s, Estado: %s\n", pirata.GetId(), pirata.GetNombre(), pirata.GetRecompensa(), pirata.GetPeligrosidad(), pirata.GetEstado())
	}
	if err := scanner.Err(); err != nil {
		return nil, 0, err
	}
	return listaPiratas, piratas_capturados, nil
}

func pirataCapturado(id_pirata int32, listaPiratas []*pb.Pirata) {
	fmt.Printf("Pirata #%d marcado como Capturado.\n", id_pirata)
	for _, pirata := range listaPiratas {
		if pirata.GetId() == id_pirata {
			pirata.Estado = "Capturado"
			break
		}
	}
}

func pirataEnBusqueda(id_pirata int32, listaPiratas []*pb.Pirata) {
	fmt.Printf("Pirata #%d marcado como Buscado.\n", id_pirata)
	for _, pirata := range listaPiratas {
		if pirata.GetId() == id_pirata {
			pirata.Estado = "Buscado"
			break
		}
	}
}

func pirataEnCaptura(id_pirata int32, listaPiratas []*pb.Pirata) {
	fmt.Printf("Pirata #%d marcado como En Captura.\n", id_pirata)
	for _, pirata := range listaPiratas {
		if pirata.GetId() == id_pirata {
			pirata.Estado = "En Captura"
			break
		}
	}
}

func interruptorRedadas(mar_ip string, mar_port string) {
	conexion, cliente := marConectar(mar_ip, mar_port)
	defer conexion.Close()
	marActividadIlegal(cliente)
}

func (s *servidor) Basic(ctx context.Context, req *pb.BasicRequest) (*pb.BasicResponse, error) {
	fmt.Println("Recibiendo mensaje:", req.GetMsg())
	msgRecibido := req.GetMsg()
	mensaje := "Hola, " + msgRecibido + "! Soy el Gobierno Mundial."
	return &pb.BasicResponse{Msg: mensaje}, nil
}

func (s *servidor) ObtenerListaPiratas(ctx context.Context, req *pb.Empty) (*pb.ListaPiratas, error) {
	// fmt.Println("Recibiendo solicitud de lista de piratas")

	// Filtrar la lista de piratas para obtener solo los que están buscados
	var listaBuscados []*pb.Pirata
	for _, pirata := range s.listaPiratas {
		if pirata.Estado == "Buscado" {
			listaBuscados = append(listaBuscados, pirata)
		}
	}

	return &pb.ListaPiratas{Piratas: listaBuscados}, nil
}

func (s *servidor) ObtenerListaPiratasTodos(ctx context.Context, req *pb.Empty) (*pb.ListaPiratas, error) {
	// fmt.Println("Recibiendo solicitud de lista de todos los piratas")
	return &pb.ListaPiratas{Piratas: s.listaPiratas}, nil
}

func (s *servidor) ObtenerNpiratasRestantes(ctx context.Context, req *pb.Empty) (*pb.NpiratasRestantes, error) {
	// fmt.Println("Recibiendo solicitud de cantidad de piratas restantes")
	return &pb.NpiratasRestantes{Cantidad: int32(s.piratasRestantes)}, nil
}

func (s *servidor) ResultadoEntrega(ctx context.Context, req *pb.ActualizarReputacionRequest) (*pb.ActualizarReputacionResponse, error) {
	fmt.Println("==============\nRecibiendo reporte de entrega de pirata.")
	// Obtener info del mensaje
	estado := req.GetEstado()
	id_pirata := req.GetIdPirata()
	sub_mun := req.GetSubmundo()
	// Manejar la solicitud
	reputacion := 0

	fmt.Printf("    ID pirata: %d, Estado: %s, Entregado a submundo: %t\n", id_pirata, estado, sub_mun)

	// Si la captura fue exitosa, se actualiza la reputación
	if estado == "Capturado" {
		fmt.Println("Se reportó captura exitosa. +10 Reputacion.")
		if sub_mun {
			fmt.Println("    Sin embargo, el pirata fue entregado al submundo...")
			s.balance--
		} else {
			s.balance++
		}
		reputacion += 10
		pirataCapturado(id_pirata, s.listaPiratas)
		s.piratasRestantes--
	} else if estado == "Perdido" {
		fmt.Println("Se reportó pirata perdido. -10 Reputacion.")
		pirataEnBusqueda(id_pirata, s.listaPiratas)
		reputacion = -10
	} else if estado == "Confiscado" {
		fmt.Println("Se reportó pirata confiscado por la marina.")
		pirataCapturado(id_pirata, s.listaPiratas)
		s.piratasRestantes--
		s.balance++
	} else if estado == "Rechazado" {
		fmt.Println("Se reportó pirata rechazado por la marina.")
		pirataEnBusqueda(id_pirata, s.listaPiratas)
	}

	if s.piratasRestantes <= 0 {
		fmt.Println("Todos los piratas han sido capturados!!!")
		s.piratasRestantes = 0
	}

	// Si hay actividad ilegal, se crean las redadas
	if s.balance <= -2 && !s.redadas {
		fmt.Println("Se detectó actividad ilegal. Avisando a la marina para activar redadas.")
		interruptorRedadas(s.mar_ip, s.mar_port)
		s.redadas = true
	} else if s.balance == 0 && s.redadas {
		fmt.Println("Ya no hay actividad ilegal. Avisando a la marina para desactivar redadas.")
		s.redadas = false
		interruptorRedadas(s.mar_ip, s.mar_port)
	}

	fmt.Println("Estado final:")
	fmt.Printf("    Balance de actividad ilegal: %d\n", s.balance)
	fmt.Printf("    Cantidad de piratas restantes: %d\n", s.piratasRestantes)
	fmt.Printf("    Redadas activas: %t\n==============\n", s.redadas)

	return &pb.ActualizarReputacionResponse{Reputacion: int32(reputacion)}, nil
}

func (s *servidor) MarcarCaptura(ctx context.Context, req *pb.PirataCapturado) (*pb.Empty, error) {
	fmt.Println("==============\nUn cazarecompensa esta capturando un pirata.")

	id_pirata := req.GetIdPirata()
	pirataEnCaptura(id_pirata, s.listaPiratas)

	fmt.Printf("=============\n")
	return &pb.Empty{}, nil
}

func main() {

	//=== SETEAR VARIABLES DE ENTORNO DESDE DOCKER-COMPOSE
	ipAddr := "0.0.0.0"
	serverIp := os.Getenv("IP_ADDR")
	if serverIp == "" {
		serverIp = ipAddr
	}
	portNumbr := "50053"
	serverPort := os.Getenv("PORT_NUMBR")
	if serverPort == "" {
		serverPort = portNumbr
	}

	mar_ip := os.Getenv("MAR_IP_ADDR")
	if mar_ip == "" {
		mar_ip = "127.0.0.1"
	}
	mar_port := os.Getenv("MAR_PORT_NUMBR")
	if mar_port == "" {
		mar_port = "50052"
	}

	pir_file := os.Getenv("PIRATAS_FILE")
	if pir_file == "" {
		pir_file = "piratas.csv"
	}
	//=== FIN VARIABLES ENTORNO

	//=== Cargar piratas
	fmt.Println("Cargando piratas desde el archivo:", pir_file)
	listaPiratas, piratasCapturados, err := cargarPiratas(pir_file)
	if err != nil {
		fmt.Println("    Error al cargar la lista de piratas:", err)
		listaPiratas = []*pb.Pirata{}
	}
	piratasPorBuscar := len(listaPiratas) - piratasCapturados

	lis, err := net.Listen("tcp", serverIp+":"+serverPort)
	if err != nil {
		panic(err)
	}

	server := &servidor{
		listaPiratas:     listaPiratas,
		balance:          0,
		mar_ip:           mar_ip,
		mar_port:         mar_port,
		redadas:          false,
		piratasRestantes: piratasPorBuscar,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGobiernoServiceServer(grpcServer, server)

	fmt.Println("Servidor gRPC escuchando en " + serverIp + ":" + serverPort)
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
