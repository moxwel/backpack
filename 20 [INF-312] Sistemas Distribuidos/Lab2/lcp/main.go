package main

import (
	pb "lcp/PokemonProto"

	"l2/types"

	"google.golang.org/grpc"

	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/streadway/amqp"

	"encoding/json"
)

func main() {
	//=== VARIABLES DE ENTORNO ===//
	server_ip := os.Getenv("IP")
	if server_ip == "" {
		server_ip = "0.0.0.0"
	}
	server_port := os.Getenv("PORT")
	if server_port == "" {
		server_port = "50051"
	}

	gym_ip := os.Getenv("GYM_IP")
	if gym_ip == "" {
		gym_ip = "127.0.0.1"
	}
	gym_port := os.Getenv("GYM_PORT")
	if gym_port == "" {
		gym_port = "50052"
	}

	rabbit_ip := os.Getenv("RABBIT_IP")
	if rabbit_ip == "" {
		rabbit_ip = "127.0.0.1"
	}
	rabbit_port := os.Getenv("RABBIT_PORT")
	if rabbit_port == "" {
		rabbit_port = "5672"
	}
	rabbit_user := os.Getenv("RABBIT_USER")
	if rabbit_user == "" {
		rabbit_user = "guest"
	}
	rabbit_pass := os.Getenv("RABBIT_PASS")
	if rabbit_pass == "" {
		rabbit_pass = "guest"
	}
	//============

	//=== Rabbit y cliente gRPC para Gym ===//
	rabbitConn, rabbitCh, err := conectarRabbit(rabbit_ip, rabbit_port, rabbit_user, rabbit_pass)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rabbitConn.Close()
	defer rabbitCh.Close()

	err = declararLCPEventsQueue(rabbitCh)
	if err != nil {
		fmt.Println(err)
		return
	}

	rabbitChConsumer, err := rabbitConn.Channel()
	if err != nil {
		fmt.Println("Error al crear canal de consumidor:", err)
		return
	}
	defer rabbitChConsumer.Close()

	gymConn, clienteGym, err := conectarGym(gym_ip, gym_port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer gymConn.Close()
	//======================

	fmt.Println("[LCP] Iniciando servidor...")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server_ip, server_port))
	if err != nil {
		panic("    [ ! ] No se pudo iniciar el listener: " + err.Error())
	}

	// === Penalizaciones compartidas ===
	listaPenalizaciones := []types.Penalizaciones{}
	listaEntrenadores := []types.Entrenador{}

	server := &ServidorLCP{
		listaEntrenadores:   &listaEntrenadores,
		listaTorneos:        []types.Torneo{},
		listaInscripciones:  []types.InscripcionTorneo{},
		gymClient:           clienteGym,
		listaPenalizaciones: &listaPenalizaciones,
	}

	go iniciarGeneradorTorneos(server, rabbitCh)
	go escucharCDPResultQueue(rabbitChConsumer, rabbitCh, server.listaEntrenadores, &server.listaTorneos, &listaPenalizaciones)

	grpcServer := grpc.NewServer()
	pb.RegisterServicioLCPServer(grpcServer, server)

	fmt.Println("[LCP] Servidor escuchando en", fmt.Sprintf("%s:%s", server_ip, server_port))
	if err := grpcServer.Serve(listener); err != nil {
		panic("    [ ! ] Fallo al servir: " + err.Error())
	}

	fmt.Println(clienteGym)
}

func iniciarGeneradorTorneos(server *ServidorLCP, chProd *amqp.Channel) {
	regiones := []string{"kanto", "johto", "sinnoh"}
	for {
		time.Sleep(10 * time.Second)
		activos := 0
		for _, t := range server.listaTorneos {
			if t.Estado == 0 { // 0 = TORNEO_ACTIVO
				activos++
			}
		}
		if activos < 5 {
			nuevaRegion := regiones[rand.Intn(len(regiones))]
			nuevoTorneo := types.Torneo{
				Id:     uuid.New().String(),
				Region: nuevaRegion,
				Estado: 0, // TORNEO_ACTIVO
			}
			server.listaTorneos = append(server.listaTorneos, nuevoTorneo)
			fmt.Println("[iniciarGeneradorTorneos] Nuevo torneo agregado:", nuevoTorneo)

			// Publicar el nuevo torneo en RabbitMQ
			mensaje_nuevo_torneo := types.NuevoTorneo{
				IdTorneo:    nuevoTorneo.Id,
				Region:      nuevoTorneo.Region,
				Estado:      true,
				Fecha:       time.Now().Format("2006-01-02"),
				TipoMensaje: "nuevo_torneo",
			}
			mensaje_nuevo_torneo_json, err := json.Marshal(mensaje_nuevo_torneo)
			if err != nil {
				fmt.Println("    [ ! ] Error al serializar el nuevo torneo:", err)
				continue
			}
			err = SendLCPEventsQueue(chProd, "nuevo_torneo", string(mensaje_nuevo_torneo_json))
			if err != nil {
				fmt.Println("    [ ! ] Error al enviar el nuevo torneo a RabbitMQ:", err)
			}
		}
	}
}
