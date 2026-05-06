package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"os"

	pb "marine/servicioMensajes"

	"google.golang.org/grpc"
)

type servidor struct {
	pb.UnimplementedMarinaServiceServer
	redadas      bool
	sub_ip       string
	sub_port     string
	listaPiratas []*pb.Pirata
	gob_ip       string
	gob_port     string
}

func (s *servidor) Basic(ctx context.Context, req *pb.BasicRequest) (*pb.BasicResponse, error) {
	fmt.Println("Recibiendo mensaje:", req.GetMsg())
	msgRecibido := req.GetMsg()
	mensaje := "Hola, " + msgRecibido + "! Soy la Marina."
	return &pb.BasicResponse{Msg: mensaje}, nil
}

func (s *servidor) EntregarPirata(ctx context.Context, req *pb.EntregaRequest) (*pb.EntregaResponse, error) {
	fmt.Println("==============\nRecibiendo un pirata...")

	// Obtener info del mensaje
	balance := req.GetBalance()
	reputacion := req.GetReputacion()
	pirata := req.GetPirata()
	peligrosidad := req.GetPirata().GetPeligrosidad()

	fmt.Printf("    ID: %d, Nombre: %s, Recompensa: %d, Peligrosidad: %s, Estado: %s\n", pirata.GetId(), pirata.GetNombre(), pirata.GetRecompensa(), pirata.GetPeligrosidad(), pirata.GetEstado())

	// Conectar a Gobierno Mundial
	fmt.Println("Voy a ver la lista de piratas del Gobierno Mundial...")
	conexion, cliente := gobConectar(s.gob_ip, s.gob_port)
	defer conexion.Close()

	// Obtener lista de piratas
	s.listaPiratas = gobListaPiratasTodos(cliente)

	// Verificar si el pirata está en la lista
	encontrado := false
	for _, p := range s.listaPiratas {
		if p.Id == pirata.Id {
			if pirata.Estado == "Capturado" {
				encontrado = true
				break
			}
		}
	}
	if encontrado {
		fmt.Println("    El pirata ya fue capturado por otro cazarecompensas! Rechazado.\n==============")
		return &pb.EntregaResponse{Estado: "Rechazado", Pago: 0}, nil
	}

	//Verificar si el pirata escapa
	if peligrosidad == "Baja" {
		random := rand.Intn(100)
		if random < 15 {
			fmt.Println("El pirata se ha escapado!\n==============")
			return &pb.EntregaResponse{Estado: "Perdido", Pago: 0}, nil
		}
	} else if peligrosidad == "Media" {
		random := rand.Intn(100)
		if random < 25 {
			fmt.Println("El pirata se ha escapado!\n==============")
			return &pb.EntregaResponse{Estado: "Perdido", Pago: 0}, nil
		}
	} else {
		random := rand.Intn(100)
		if random < 45 {
			fmt.Println("El pirata se ha escapado!\n==============")
			return &pb.EntregaResponse{Estado: "Perdido", Pago: 0}, nil
		}
	}

	// Si hay mucha recompensa por el pirata, puede perderse por un ataque del submundo
	if pirata.Recompensa > 200000000 {
		random := rand.Intn(100)
		if random < 35 {
			fmt.Println("Nos ha atacado el submundo!\n    El pirata se ha escapado!\n==============")
			return &pb.EntregaResponse{Estado: "Perdido", Pago: 0}, nil
		}
	}

	// Si tiene mala reputación lo rechaza
	if reputacion < 50 {
		fmt.Println("El cazarecompensas tiene muy mala reputación! Rechazado.\n==============")
		return &pb.EntregaResponse{Estado: "Rechazado", Pago: 0}, nil
	}

	pago := int32(pirata.Recompensa)

	// Si hay mucha actividad ilegal, el pago se reduce a la mitad
	if balance < -3 {
		fmt.Println("El cazarecompensas ha hecho mucha actividad ilegal! El pago se reduce a la mitad.")
		pago = int32(pirata.Recompensa / 2)
	}

	fmt.Printf("Entrega completa. Pagando %d berries.\n==============\n", pago)
	return &pb.EntregaResponse{Estado: "Capturado", Pago: pago}, nil
}

func (s *servidor) ActividadIlegal(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	fmt.Println("==============\nEl Gobierno Mundial me ha pedido activar/desactivar redadas!")

	conexion, cliente := subConectar(s.sub_ip, s.sub_port)
	defer conexion.Close()

	subActivarRedadas(cliente)
	s.redadas = !s.redadas

	fmt.Printf("    Redadas activadas: %t\n==============\n", s.redadas)
	return &pb.Empty{}, nil
}

func main() {
	//=== SETEAR VARIABLES DE ENTORNO DESDE DOCKER-COMPOSE
	ipAddr := "0.0.0.0"
	serverIp := os.Getenv("IP_ADDR")
	if serverIp == "" {
		serverIp = ipAddr
	}

	portNumbr := "50052"
	serverPort := os.Getenv("PORT_NUMBR")
	if serverPort == "" {
		serverPort = portNumbr
	}

	sub_ip := os.Getenv("SUB_IP_ADDR")
	if sub_ip == "" {
		sub_ip = "127.0.0.1"
	}
	sub_port := os.Getenv("SUB_PORT_NUMBR")
	if sub_port == "" {
		sub_port = "50051"
	}

	gob_port := os.Getenv("GOB_PORT_NUMBR")
	if gob_port == "" {
		gob_port = "50053"
	}
	gob_ip := os.Getenv("GOB_IP_ADDR")
	if gob_ip == "" {
		gob_ip = "127.0.0.1"
	}

	//=== FIN VARIABLES ENTORNO

	lis, err := net.Listen("tcp", serverIp+":"+serverPort)
	if err != nil {
		panic(err)
	}

	server := &servidor{
		redadas:      false,
		sub_ip:       sub_ip,
		sub_port:     sub_port,
		gob_ip:       gob_ip,
		gob_port:     gob_port,
		listaPiratas: []*pb.Pirata{},
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMarinaServiceServer(grpcServer, server)

	fmt.Println("Servidor gRPC escuchando en " + serverIp + ":" + serverPort)
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
