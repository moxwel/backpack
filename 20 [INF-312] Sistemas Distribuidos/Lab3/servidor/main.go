package main

import (
	"fmt"
	"l3/types"
	"math/rand"
	"net"
	"os"
	pb "servidor/MatchmakingProto"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {
	//=== VARIABLES DE ENTORNO ===//

	server_ip := os.Getenv("IP")
	if server_ip == "" {
		server_ip = "0.0.0.0"
	}
	server_port := os.Getenv("PORT")
	if server_port == "" {
		server_port = "50000"
	}
	server_id := os.Getenv("SERVERID")
	if server_id == "" {
		server_id = "defaultserver"
	}

	real_ip := os.Getenv("REAL_IP")
	if real_ip == "" {
		real_ip = server_id
	}

	match_ip := os.Getenv("MATCH_IP")
	if match_ip == "" {
		match_ip = "matchmaker"
	}
	match_port := os.Getenv("MATCH_PORT")
	if match_port == "" {
		match_port = "50000"
	}
	//============

	fmt.Printf("[GameServer] Iniciando servidor de juegos...\n    IP: %s\n    PORT: %s\n    SERVERID: %s\n    MATCH_IP: %s\n    MATCH_PORT: %s\n", server_ip, server_port, server_id, match_ip, match_port)
	str_server_address_grpc := fmt.Sprintf("%s:%s", server_ip, server_port) // String para abrir el listener gRPC
	str_server_address := fmt.Sprintf("%s:%s", real_ip, server_port)        // String que se enviará al matchmaker

	//======= CONEXION MATCHMAKER ======//
	conn, matchmakerClient, err := conectarMatchmaking(match_ip, match_port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	clock := types.VectorClockMap{
		server_id: 0, // Inicializar el reloj vectorial con el servidor actual
	}

	// Registrar el servidor en el matchmaker
	cod_resp, msg, reloj, err := matchmakingUpdateServerStatus(matchmakerClient, server_id, str_server_address, types.DISPONIBLE, clock)

	fusionarMayores(clock, reloj)

	if err != nil {
		fmt.Printf("[%s] [ ! ] Error al registrar el servidor en el matchmaker: %v\n", server_id, err)
		// Reintentar la conexión al matchmaker despues de 5 segundos 5 veces
		for i := 0; i < 5; i++ {
			fmt.Printf("[%s] [ ! ] Reintentando conexión al matchmaker (%d/5)...\n", server_id, i+1)
			time.Sleep(5 * time.Second)
			cod_resp, msg, reloj, err = matchmakingUpdateServerStatus(matchmakerClient, server_id, str_server_address, types.DISPONIBLE, clock)
			fusionarMayores(clock, reloj)
			if err == nil && cod_resp == 0 {
				fmt.Printf("[%s] [INFO] Servidor registrado correctamente en el matchmaker tras reintento.\n", server_id)
				break
			}
		}
		if err != nil {
			fmt.Printf("[%s] [ ! ] Error final al registrar el servidor en el matchmaker: %v\n", server_id, err)
			return
		}
	}
	if cod_resp != 0 {
		fmt.Printf("[%s] [ ! ] Respuesta del matchmaker: Codigo %d, Mensaje: %s\n", server_id, cod_resp, msg)
		return
	}
	fmt.Printf("[%s] [INFO] Servidor registrado correctamente en el matchmaker.\n", server_id)
	//==== Fin conexion matchmaker ======//

	// Crear instancia del servidor de juegos
	server := &ServidorJuegos{
		server_id:        server_id,
		server_address:   str_server_address,
		current_match_id: "",
		current_player_1: "",
		current_player_2: "",
		status:           0, // DISPONIBLE
		matchmakerClient: matchmakerClient,
		clock:            clock,
		resultsMutex:     sync.Mutex{},
		matchResults:     map[string]types.MatchResult{},
	}

	listener, err := net.Listen("tcp", str_server_address_grpc)
	if err != nil {
		panic("    [ ! ] No se pudo iniciar el listener: " + err.Error())
	}
	go prob_caida(server) // Llamar a la función de probabilidad de caída

	grpcServer := grpc.NewServer()
	pb.RegisterGameServerServiceServer(grpcServer, server)

	fmt.Printf("[%s] Servidor escuchando en %s\n", server_id, str_server_address_grpc)
	if err := grpcServer.Serve(listener); err != nil {
		panic("    [ ! ] Fallo al servir: " + err.Error())
	}
}

func prob_caida(server *ServidorJuegos) {
	//el server cada 20 segundos tiene un 30% de probabilidad de caerse
	for {
		if server.status == types.CAIDO {
			return // Si ya está caído, no hacemos nada
		}
		if probabilidad := rand.Intn(100); probabilidad < 30 { // 30% de probabilidad
			fmt.Printf("[ProbCaida] El servidor '%s' ha caído.\n", server.server_id)
			server.status = types.CAIDO
			server.current_match_id = ""
			server.current_player_1 = ""
			server.current_player_2 = ""
		}
		// Esperar 20 segundos antes de la siguiente comprobación
		time.Sleep(20 * time.Second)
	}

}
