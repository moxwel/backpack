package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"os"

	pb "submun/servicioMensajes"

	"google.golang.org/grpc"
)

type servidor struct {
	pb.UnimplementedSubMundoServiceServer
	redadas bool
}

func (s *servidor) Basic(ctx context.Context, req *pb.BasicRequest) (*pb.BasicResponse, error) {
	fmt.Println("Recibiendo mensaje:", req.GetMsg())
	msgRecibido := req.GetMsg()
	mensaje := "Hola, " + msgRecibido + "! Soy el SubMundo."
	return &pb.BasicResponse{Msg: mensaje}, nil
}

func (s *servidor) ActivarRedadas(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	fmt.Println("==============\nAhi viene la Marina!")
	s.redadas = !s.redadas
	fmt.Printf("    Redadas activadas: %t\n==============\n", s.redadas)
	return &pb.Empty{}, nil
}

func (s *servidor) EntregarPirata(ctx context.Context, req *pb.EntregaRequest) (*pb.EntregaResponse, error) {
	fmt.Println("==============\nRecibiendo un pirata...")

	// Obtener info del mensaje
	pirata_recompensa := req.GetPirata().GetRecompensa()
	pirata_peligro := req.GetPirata().GetPeligrosidad()
	pirata := req.GetPirata()
	fmt.Printf("    ID: %d, Nombre: %s, Recompensa: %d, Peligrosidad: %s, Estado: %s\n", pirata.GetId(), pirata.GetNombre(), pirata.GetRecompensa(), pirata.GetPeligrosidad(), pirata.GetEstado())

	//Verificar si el pirata escapa
	p_escape := rand.Intn(100)
	if pirata_peligro == "Baja" {
		if p_escape < 15 {
			fmt.Println("El pirata se ha escapado!\n==============")
			return &pb.EntregaResponse{Estado: "Perdido", Pago: 0}, nil
		}
	} else if pirata_peligro == "Media" {
		if p_escape < 25 {
			fmt.Println("El pirata se ha escapado!\n==============")
			return &pb.EntregaResponse{Estado: "Perdido", Pago: 0}, nil
		}
	} else {
		if p_escape < 45 {
			fmt.Println("El pirata se ha escapado!\n==============")
			return &pb.EntregaResponse{Estado: "Perdido", Pago: 0}, nil
		}
	}

	//Verificar si el pirata es capturado por la redada
	p_captura := rand.Intn(100)
	if s.redadas {
		if p_captura < 30 {
			fmt.Println("El pirata ha sido confiscado por la marina!\n==============")
			return &pb.EntregaResponse{Estado: "Confiscado", Pago: 0}, nil
		}
	}

	//Probabilidad de que el submundo realice estafa
	p_estafa := rand.Intn(100)
	if p_estafa < 35 {
		fmt.Println("El submundo ha decidido estafar al cazarecompensas!\n==============")
		return &pb.EntregaResponse{Estado: "Capturado", Pago: 0}, nil
	}

	//Calcular el pago
	bono_pago := rand.Float64()*0.5 + 1
	pago_final := int32(float64(pirata_recompensa) * bono_pago)
	fmt.Printf("Entrega completa. Pagando %d berries con bono de x%.2f. Monto total: %d\n==============\n", pirata_recompensa, bono_pago, pago_final)

	return &pb.EntregaResponse{Estado: "Capturado", Pago: pago_final}, nil
}

func main() {
	//=== SETEAR VARIABLES DE ENTORNO DESDE DOCKER-COMPOSE
	ipAddr := "0.0.0.0"
	serverIp := os.Getenv("IP_ADDR")
	if serverIp == "" {
		serverIp = ipAddr
	}

	portNumbr := "50051"
	serverPort := os.Getenv("PORT_NUMBR")
	if serverPort == "" {
		serverPort = portNumbr
	}
	//=== FIN VARIABLES ENTORNO

	lis, err := net.Listen("tcp", serverIp+":"+serverPort)
	if err != nil {
		panic(err)
	}

	server := &servidor{
		redadas: false,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSubMundoServiceServer(grpcServer, server)

	fmt.Println("Servidor gRPC escuchando en " + serverIp + ":" + serverPort)
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
