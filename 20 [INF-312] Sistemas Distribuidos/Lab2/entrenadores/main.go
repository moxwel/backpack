package main

import (
	"encoding/json"
	pb "entrenadores/PokemonProto"
	"fmt"
	"l2/types"
	"os"
	"strings"
)

func main() {
	//=== VARIABLES DE ENTORNO ===//
	lcp_ip := os.Getenv("LCP_IP")
	if lcp_ip == "" {
		lcp_ip = "127.0.0.1"
	}
	lcp_port := os.Getenv("LCP_PORT")
	if lcp_port == "" {
		lcp_port = "50051"
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

	//=== Conectar a RabbitMQ y cliente gRPC de LCP ===//
	rabbitConn, rabbitCh, err := conectarRabbit(rabbit_ip, rabbit_port, rabbit_user, rabbit_pass)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rabbitConn.Close()
	defer rabbitCh.Close()

	lcpConn, clienteLCP, err := conectarLCP(lcp_ip, lcp_port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer lcpConn.Close()
	//===============

	//=== Registrar entrenadores en LCP ===//
	lista_entrenadores, _ := procesarArchivoEntrenadores("entrenadores.json")
	for i, entrenador := range lista_entrenadores {
		fmt.Println(entrenador)
		err := lcpRegistrarEntrenador(clienteLCP, entrenador)
		if err != nil {
			fmt.Printf("Error al registrar el entrenador %s: %v\n", entrenador.Nombre, err)
			// Entrenador ya esta registrado, se consulta al LCP y se obtienen sus datos
			entrenadorLCP, err2 := lcpObtenerEntrenador(clienteLCP, entrenador.Id)
			if err2 == nil {
				lista_entrenadores[i] = entrenadorLCP
				fmt.Printf("Entrenador %s actualizado desde LCP.\n", entrenador.Nombre)
			} else {
				fmt.Printf("No se pudo obtener datos del entrenador %s desde LCP: %v\n", entrenador.Nombre, err2)
			}
		} else {
			fmt.Printf("Entrenador %s registrado exitosamente.\n", entrenador.Nombre)
		}
	}
	//===============

	var lista_historial []types.RegistroCombate
	go escucharSNPNotifyQueue(rabbitCh, &lista_entrenadores, &lista_historial)

	// Selección de entrenador
	var entrenadorSeleccionado *types.Entrenador
	seleccionarEntrenador := func(lista_entrenadores []types.Entrenador) int {
		for {
			fmt.Println("=== Seleccione un entrenador para usar ===")
			for i, e := range lista_entrenadores {
				fmt.Printf("    %d. %s (Región: %s, Ranking: %d, Estado: %s, Suspension: %d)\n", i+1, e.Nombre, e.Region, e.Ranking, mapEstadoEntrenador(e.Estado), e.Suspencion)
			}
			fmt.Print("Ingrese el número del entrenador: ")
			var seleccion int
			_, err := fmt.Scanf("%d\n", &seleccion)
			if err != nil || seleccion < 1 || seleccion > len(lista_entrenadores) {
				fmt.Println("Selección inválida. Intente nuevamente.")
				continue
			}
			return seleccion - 1
		}
	}

	indiceSeleccionado := seleccionarEntrenador(lista_entrenadores)
	entrenadorSeleccionado = &lista_entrenadores[indiceSeleccionado]

	var opcion int
	for {
		fmt.Printf("==== MENU PRINCIPAL - %s %s (%s) %s (%d) ====\n    1. Mostrar estado de entrenadores\n    2. Ver torneos disponibles\n    3. Inscribirse a un torneo\n    4. Ver historial\n    5. Ver notificaciones\n    6. Cambiar entrenador activo\n    0. Salir\n",
			entrenadorSeleccionado.Id,
			entrenadorSeleccionado.Nombre,
			entrenadorSeleccionado.Region,
			mapEstadoEntrenador(entrenadorSeleccionado.Estado),
			entrenadorSeleccionado.Suspencion,
		)
		fmt.Print("Ingrese una opción: ")
		_, err := fmt.Scanf("%d", &opcion)
		if err != nil {
			fmt.Println("Error al leer la opción:", err)
			continue
		}

		switch opcion {
		case 1:
			mostrarEstadoEntrenadores(lista_entrenadores)
		case 2:
			mostrarTorneosDisponibles(clienteLCP)
		case 3:
			ejecutarInscripcion(clienteLCP, entrenadorSeleccionado)
		case 4:
			mostrarHistorial(lista_historial)
		case 5:
			mostrarNotificaciones()
		case 6:
			fmt.Println("Cambiando entrenador activo...")
			indiceSeleccionado = seleccionarEntrenador(lista_entrenadores)
			entrenadorSeleccionado = &lista_entrenadores[indiceSeleccionado]
		case 0:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción no válida. Intente nuevamente.")
		}
	}
}

func mostrarNotificaciones() {
	fmt.Println("=== Notificaciones ===")
	file, err := os.Open("notify_logs.txt")
	if err != nil {
		fmt.Println("No hay notificaciones disponibles.")
		return
	}
	defer file.Close()
	buf := make([]byte, 4096)
	for {
		n, err := file.Read(buf)
		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
		if err != nil {
			break
		}
	}
	fmt.Println("======================")
}

func procesarArchivoEntrenadores(file_name string) ([]types.Entrenador, error) {
	type entrenadorJSON struct {
		Id         string `json:"id"`
		Nombre     string `json:"nombre"`
		Region     string `json:"region"`
		Ranking    int32  `json:"ranking"`
		Estado     string `json:"estado"`
		Suspension int32  `json:"suspension"`
	}
	var entrenadoresJSONTemp []entrenadorJSON

	mapEstado := func(estado string) types.EstadoEntrenador {
		switch strings.ToLower(estado) {
		case "activo":
			return types.ENTRENADOR_ACTIVO
		case "inactivo":
			return types.ENTRENADOR_INACTIVO
		case "suspendido":
			return types.ENTRENADOR_SUSPENDIDO
		default:
			return types.ENTRENADOR_INACTIVO
		}
	}

	file, err := os.Open(file_name)
	if err != nil {
		fmt.Printf("Error al abrir el archivo entrenadores.json: %v\n", err)
		return nil, err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&entrenadoresJSONTemp); err != nil {
		fmt.Printf("Error al decodificar el archivo entrenadores.json: %v\n", err)
		return nil, err
	}

	var entrenadores []types.Entrenador
	for _, e := range entrenadoresJSONTemp {
		entrenadores = append(entrenadores, types.Entrenador{
			Id:         e.Id,
			Nombre:     e.Nombre,
			Region:     strings.ToLower(e.Region),
			Ranking:    e.Ranking,
			Estado:     mapEstado(e.Estado),
			Suspencion: e.Suspension,
		})
	}

	if len(entrenadores) == 0 {
		fmt.Println("No se encontraron entrenadores en el archivo.")
		return nil, fmt.Errorf("no se encontraron entrenadores en el archivo")
	}
	fmt.Printf("Se encontraron %d entrenadores en el archivo.\n", len(entrenadores))
	return entrenadores, nil
}

func mostrarEstadoEntrenadores(entrenadores []types.Entrenador) {
	fmt.Println("=== Estado de los entrenadores ===")
	for _, entrenador := range entrenadores {
		fmt.Printf("    - ID: %s, Nombre: %s, Región: %s, Ranking: %d, Estado: %d (%s), Suspensión: %d\n",
			entrenador.Id, entrenador.Nombre, entrenador.Region, entrenador.Ranking,
			entrenador.Estado, mapEstadoEntrenador(entrenador.Estado), entrenador.Suspencion)
	}
	fmt.Println("===================================")
}

func mapEstadoEntrenador(estado types.EstadoEntrenador) string {
	switch estado {
	case types.ENTRENADOR_ACTIVO:
		return "Activo"
	case types.ENTRENADOR_INACTIVO:
		return "Inactivo"
	case types.ENTRENADOR_SUSPENDIDO:
		return "Suspendido"
	default:
		return "Desconocido"
	}
}

func mostrarTorneosDisponibles(clienteLCP pb.ServicioLCPClient) {
	fmt.Println("=== Torneos Disponibles ===")

	torneos, err := lcpObtenerTorneos(clienteLCP)
	if err != nil {
		fmt.Println("Error al obtener torneos disponibles:", err)
		return
	}
	if len(torneos) == 0 {
		fmt.Println("No hay torneos disponibles en este momento.")
		return
	}

	for _, torneo := range torneos {
		fmt.Printf("    - ID: %s, Región: %s, Estado: %d (%s)\n",
			torneo.Id, torneo.Region, torneo.Estado, mapEstadoTorneo(torneo.Estado))
	}
}

func mapEstadoTorneo(estado types.EstadoTorneo) string {
	switch estado {
	case types.TORNEO_ACTIVO:
		return "Activo"
	case types.TORNEO_PENDIENTE:
		return "Pendiente"
	case types.TORNEO_FINALIZADO:
		return "Finalizado"
	default:
		return "Desconocido"
	}
}

func ejecutarInscripcion(clienteLCP pb.ServicioLCPClient, entrenador *types.Entrenador) {
	torneos, err := lcpObtenerTorneosRegion(clienteLCP, entrenador.Region)
	if err != nil {
		fmt.Println("Error al obtener torneos de la región:", err)
		return
	}
	if len(torneos) == 0 {
		fmt.Println("No hay torneos activos disponibles en tu región.")
		return
	}
	fmt.Println("=== Torneos Activos en tu Región ===")
	for i, t := range torneos {
		fmt.Printf("    %d. ID: %s, Región: %s, Estado: %s\n", i+1, t.Id, t.Region, mapEstadoTorneo(t.Estado))
	}
	fmt.Print("Seleccione el número del torneo para inscribirse: ")
	var seleccion int
	_, err = fmt.Scanf("%d\n", &seleccion)
	if err != nil || seleccion < 1 || seleccion > len(torneos) {
		fmt.Println("Selección inválida.")
		return
	}
	torneoSeleccionado := torneos[seleccion-1]
	err = lcpInscribirTorneo(clienteLCP, entrenador, torneoSeleccionado.Id)
	if err != nil {
		fmt.Println("No se pudo inscribir en el torneo:", err)
	} else {
		fmt.Println("Inscripción exitosa en el torneo.")
	}
}

func mostrarHistorial(historial []types.RegistroCombate) {
	fmt.Println("=== Historial de Combates ===")
	if len(historial) == 0 {
		fmt.Println("No hay registros de combates disponibles.")
		return
	}
	for i, registro := range historial {
		fmt.Printf("%d. Entrenador: %s (%s) | Torneo: %s | Resultado: %s | Ranking: %d -> %d\n",
			i+1, registro.Nombre, registro.IdEntrenador, registro.IdTorneo, registro.Resultado, registro.RankingAntes, registro.RankingDespues)
	}
	fmt.Println("==============================")
}
