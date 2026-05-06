package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	pb "cazarec/servicioMensajes"

	"google.golang.org/grpc"
)

type cazarrecompensas struct {
	id         int32
	reputacion int32
	billetera  int32
	balance    int32
}

func main() {
	//=== SETEAR VARIABLES DE ENTORNO DESDE DOCKER-COMPOSE
	gob_ip := os.Getenv("GOB_IP_ADDR")
	if gob_ip == "" {
		gob_ip = "127.0.0.1"
	}
	gob_port := os.Getenv("GOB_PORT_NUMBR")
	if gob_port == "" {
		gob_port = "50053"
	}

	mar_ip := os.Getenv("MAR_IP_ADDR")
	if mar_ip == "" {
		mar_ip = "127.0.0.1"
	}
	mar_port := os.Getenv("MAR_PORT_NUMBR")
	if mar_port == "" {
		mar_port = "50052"
	}

	sub_ip := os.Getenv("SUB_IP_ADDR")
	if sub_ip == "" {
		sub_ip = "127.0.0.1"
	}
	sub_port := os.Getenv("SUB_PORT_NUMBR")
	if sub_port == "" {
		sub_port = "50051"
	}

	//convertir ITER_NUMBR a int
	iter_numbr := os.Getenv("ITER_NUMBR")
	var num_rondas int
	if iter_numbr == "" {
		num_rondas = 10
	} else {
		var err error
		num_rondas, err = strconv.Atoi(iter_numbr)
		if err != nil {
			fmt.Println("Error al convertir ITER_NUMBR a int:", err)
			return
		}
	}

	//=== FIN VARIABLES ENTORNO

	//=== CONEXIONES A SERVICIOS
	var gob_conn *grpc.ClientConn
	var gob_client pb.GobiernoServiceClient
	gob_conn, gob_client = gobConectar(gob_ip, gob_port)
	defer gob_conn.Close()

	var mar_conn *grpc.ClientConn
	var mar_client pb.MarinaServiceClient
	mar_conn, mar_client = marConectar(mar_ip, mar_port)
	defer mar_conn.Close()

	var sub_conn *grpc.ClientConn
	var sub_client pb.SubMundoServiceClient
	sub_conn, sub_client = subConectar(sub_ip, sub_port)
	defer sub_conn.Close()
	//=== FIN CONEXIONES A SERVICIOS

	//=== INICIO LOGICA
	// gobBasic(gob_client, "Cazarec a gobierno")

	// marBasic(mar_client, "Cazarec a Marina")

	// subBasic(sub_client, "Cazarec a SubMundo")

	cazador := cazarrecompensas{
		id:         0,
		reputacion: 100,
		billetera:  100000000,
		balance:    0,
	}

	for i := 0; i < num_rondas; i++ {
		fmt.Println("==================\nRonda:", i+1)

		lista_piratas := gobListaPiratas(gob_client)
		fmt.Println("Lista de piratas buscados del gobierno:")
		for _, pirata := range lista_piratas {
			fmt.Printf("    ID: %d, Nombre: %s, Peligrosidad: %s, Recompensa: %d, Estado: %s\n", pirata.Id, pirata.Nombre, pirata.Peligrosidad, pirata.Recompensa, pirata.Estado)
		}

		var pirata_selec *pb.Pirata
		if len(lista_piratas) == 0 {
			fmt.Println("No hay piratas para buscar. Terminando...\n===================")
			return
		} else {
			indice := rand.Intn(len(lista_piratas))
			pirata_selec = lista_piratas[indice]
			fmt.Printf("Pirata seleccionado: #%d - %s\n", pirata_selec.Id, pirata_selec.Nombre)
		}

		p_entrega := rand.Intn(2) // Probabilidad de entrega: 0 = marina, 1 = submundo
		var nuevo_estado string
		var pago int32
		entregado_a_submundo := false

		costo_entrega := int32(pirata_selec.Recompensa / 20)
		fmt.Printf("Costo de entrega: %d - Mi billetera: %d\n", costo_entrega, cazador.billetera)

		if cazador.billetera < costo_entrega {
			fmt.Printf("    No tengo suficiente dinero para entregar a este pirata...\n==================\n")
		} else {

			fmt.Println("Voy a decirle al gobierno que ire a capturarlo...")
			gobMarcarCaptura(gob_client, pirata_selec.Id)

			cazador.billetera -= costo_entrega
			if p_entrega == 0 {
				// Entregar a la marina
				fmt.Println("Voy a entregar el pirata a la marina...")
				nuevo_estado, pago = marEntregaPirata(mar_client, cazador.balance, pirata_selec, cazador.reputacion)
				if nuevo_estado == "Capturado" {
					fmt.Println("    El pirata le llegó a la marina.")
					cazador.balance += 1
				} else {
					fmt.Println("    El pirata no le llegó a la marina....")
					if nuevo_estado == "Perdido" {
						fmt.Println("        El pirata se perdió en el camino.")
					} else if nuevo_estado == "Rechazado" {
						fmt.Println("        El pirata fue rechazado por la marina.")
					}
				}
			} else {
				// Entregar al submundo
				fmt.Println("Voy a entregar el pirata al submundo...")
				nuevo_estado, pago = subEntregaPirata(sub_client, cazador.balance, pirata_selec, cazador.reputacion)
				if nuevo_estado == "Capturado" {
					fmt.Println("    El pirata le llegó al submundo.")
					entregado_a_submundo = true
					cazador.balance -= 1
				} else {
					fmt.Println("    El pirata no le llegó al submundo....")
					if nuevo_estado == "Perdido" {
						fmt.Println("        El pirata se perdió en el camino.")
					} else if nuevo_estado == "Confiscado" {
						fmt.Println("        El pirata fue confiscado por la marina.")
					}
				}
			}

			fmt.Println("Estado final del pirata:")
			fmt.Printf("    ID: %d, Nombre: %s, Peligrosidad: %s, Recompensa: %d, Estado: %s\n", pirata_selec.Id, pirata_selec.Nombre, pirata_selec.Peligrosidad, pirata_selec.Recompensa, nuevo_estado)

			cazador.billetera += pago
			fmt.Printf("Pago recibido: %d - Mi billetera: %d\n", pago, cazador.billetera)

			fmt.Println("Voy a reportar el estado del pirata al gobierno...")
			cambio_reputacion := gobActualizarReputacion(gob_client, pirata_selec.Id, nuevo_estado, entregado_a_submundo)

			cazador.reputacion += cambio_reputacion
			fmt.Printf("Cambio de reputacion: %d - Mi reputacion: %d\n", cambio_reputacion, cazador.reputacion)

			fmt.Println("Estado final:")
			fmt.Println("    ID Cazarrecompensas:", cazador.id)
			fmt.Println("    Nueva reputacion:", cazador.reputacion)
			fmt.Println("    Plata total:", cazador.billetera)
			fmt.Println("    Balance marina/submundo:", cazador.balance)
			fmt.Println("Fin de la ejecución\n==================")
		}

	}
}
