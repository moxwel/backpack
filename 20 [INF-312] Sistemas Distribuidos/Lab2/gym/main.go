package main

import (
	"fmt"
	"l2/types"
	"net"
	"os"

	"google.golang.org/grpc"

	// Importa tu paquete proto, por ejemplo:
	pb "gym/PokemonProto"
)

func main() {
	//=== VARIABLES DE ENTORNO ===//
	ip := os.Getenv("IP")
	if ip == "" {
		ip = "0.0.0.0"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "50052"
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

	rabbitConn, rabbitCh, err := conectarRabbit(rabbit_ip, rabbit_port, rabbit_user, rabbit_pass)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rabbitConn.Close()
	defer rabbitCh.Close()

	err = declararGymResultQueue(rabbitCh)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Obtener llaves AES256 desde variables de entorno
	llaves := map[string]string{
		"kanto":  os.Getenv("AES_KEY_KANTO"),
		"johto":  os.Getenv("AES_KEY_JOHTO"),
		"sinnoh": os.Getenv("AES_KEY_SINNOH"),
	}

	// Definir gimnasios por región con su llave correspondiente
	gimnasios := []types.Gym{
		{IdGym: "gym-kanto", Region: "kanto", Estado: true, Llave: llaves["kanto"]},
		{IdGym: "gym-johto", Region: "johto", Estado: true, Llave: llaves["johto"]},
		{IdGym: "gym-sinnoh", Region: "sinnoh", Estado: true, Llave: llaves["sinnoh"]},
	}

	fmt.Println("[GYM] Iniciando servidor...")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		panic("    [ ! ] No se pudo iniciar el listener: " + err.Error())
	}

	server := &ServidorGym{
		gimnasios:          gimnasios,
		rabbit_ch:          rabbitCh,
		combatesReportados: []string{},
	}
	grpcServer := grpc.NewServer()
	pb.RegisterServicioGymServer(grpcServer, server)

	fmt.Println("[GYM] Servidor escuchando en", fmt.Sprintf("%s:%s", ip, port))
	if err := grpcServer.Serve(listener); err != nil {
		panic("    [ ! ] Fallo al servir: " + err.Error())
	}
}
