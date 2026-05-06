package main

import (
	"context"
	"fmt"
	"l2/types"
	pb "lcp/PokemonProto"
)

type ServidorLCP struct {
	pb.UnimplementedServicioLCPServer
	listaEntrenadores   *[]types.Entrenador
	listaTorneos        []types.Torneo
	listaInscripciones  []types.InscripcionTorneo
	gymClient           pb.ServicioGymClient
	listaPenalizaciones *[]types.Penalizaciones
}

func (s *ServidorLCP) RegistrarEntrenador(ctx context.Context, req *pb.Entrenador) (*pb.CodigoRespuesta, error) {
	fmt.Println("[RegistrarEntrenador] gRPC ejecutado\n    Datos recibidos:", req)

	// Verificar datos
	if req.Nombre == "" || req.Region == "" || req.Id == "" {
		fmt.Println("    [ ! ] Nombre, región o ID vacíos")
		return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "Nombre, región o ID vacíos"}, nil
	}
	if req.Ranking < 0 {
		fmt.Println("    [ ! ] Ranking inválido")
		return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "Ranking inválido"}, nil
	}
	if req.Estado != pb.EstadoEntrenador_ENTRENADOR_ACTIVO && req.Estado != pb.EstadoEntrenador_ENTRENADOR_INACTIVO && req.Estado != pb.EstadoEntrenador_ENTRENADOR_SUSPENDIDO {
		fmt.Println("    [ ! ] Estado inválido")
		return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "Estado inválido"}, nil
	}

	// Generar entrenador
	entrenador := types.Entrenador{
		Id:         req.Id,
		Nombre:     req.Nombre,
		Region:     req.Region,
		Ranking:    req.Ranking,
		Estado:     types.EstadoEntrenador(req.Estado),
		Suspencion: req.Suspencion,
	}
	fmt.Println("    [ OK ] Entrenador generado:", entrenador)

	// Verificar si el entrenador ya existe
	for _, e := range *s.listaEntrenadores {
		if e.Id == req.Id {
			fmt.Println("    [ ! ] Entrenador ya registrado")
			return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "Entrenador ya registrado"}, nil
		}
	}
	// Agregar entrenador a la lista
	*s.listaEntrenadores = append(*s.listaEntrenadores, entrenador)
	fmt.Println("    [ OK ] Entrenador agregado a la lista")

	// Agregar entrenador a la lista de penalizaciones.
	if req.Suspencion > 0 {
		*s.listaPenalizaciones = append(*s.listaPenalizaciones, types.Penalizaciones{
			IdEntrenador: entrenador.Id,
			Penalizacion: 1,
		})
	} else {
		*s.listaPenalizaciones = append(*s.listaPenalizaciones, types.Penalizaciones{
			IdEntrenador: entrenador.Id,
			Penalizacion: 0,
		})
	}

	/* Imprimir lista de entrenadores
	fmt.Println("    [ OK ] Lista de entrenadores actualizada:")
	for _, e := range s.listaEntrenadores {
		fmt.Printf("        - ID: %s, Nombre: %s, Región: %s, Ranking: %d, Estado: %d, Suspensión: %d\n",
			e.Id, e.Nombre, e.Region, e.Ranking, e.Estado, e.Suspencion)
	}
	*/

	return &pb.CodigoRespuesta{Codigo: 0, Mensaje: "Correcto"}, nil
}

func (s *ServidorLCP) ObtenerEntrenador(ctx context.Context, req *pb.EntrenadorID) (*pb.Entrenador, error) {
	fmt.Println("[ObtenerEntrenador] gRPC ejecutado\n    Datos recibidos:", req)

	// Verificar ID
	if req.Id == "" {
		fmt.Println("    [ ! ] ID vacío")
		return nil, fmt.Errorf("ID vacío")
	}

	// Buscar entrenador
	for _, e := range *s.listaEntrenadores {
		if e.Id == req.Id {
			fmt.Println("    [ OK ] Entrenador encontrado:", e)
			return &pb.Entrenador{
				Id:         e.Id,
				Nombre:     e.Nombre,
				Region:     e.Region,
				Ranking:    e.Ranking,
				Estado:     pb.EstadoEntrenador(e.Estado),
				Suspencion: e.Suspencion,
			}, nil
		}
	}

	fmt.Println("    [ ! ] Entrenador no encontrado")
	return nil, fmt.Errorf("entrenador no encontrado")
}

func (s *ServidorLCP) ExisteEntrenador(ctx context.Context, req *pb.EntrenadorID) (*pb.CodigoRespuesta, error) {
	fmt.Println("[ExisteEntrenador] gRPC ejecutado\n    Datos recibidos:", req)

	// Verificar ID
	if req.Id == "" {
		fmt.Println("    [ ! ] ID vacío")
		return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "ID vacío"}, nil
	}

	// Buscar entrenador
	for _, e := range *s.listaEntrenadores {
		if e.Id == req.Id {
			fmt.Println("    [ OK ] Entrenador existe:", e)
			return &pb.CodigoRespuesta{Codigo: 0, Mensaje: "Entrenador existe"}, nil
		}
	}

	fmt.Println("    [ ! ] Entrenador no encontrado")
	return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "Entrenador no encontrado"}, nil
}

func (s *ServidorLCP) ObtenerTorneos(ctx context.Context, req *pb.Empty) (*pb.ListaTorneos, error) {
	fmt.Println("[ObtenerTorneos] gRPC ejecutado\n    Datos recibidos:", req)

	var lista pb.ListaTorneos
	for _, t := range s.listaTorneos {
		lista.Torneos = append(lista.Torneos, &pb.Torneo{
			Id:     t.Id,
			Region: t.Region,
			Estado: pb.EstadoTorneo(t.Estado),
		})
	}
	fmt.Println("    [ OK ] Torneos encontrados:", len(lista.Torneos))
	return &lista, nil
}

func (s *ServidorLCP) ObtenerTorneosRegion(ctx context.Context, req *pb.RegionString) (*pb.ListaTorneos, error) {
	fmt.Println("[ObtenerTorneosRegion] gRPC ejecutado\n    Datos recibidos:", req)

	var lista pb.ListaTorneos
	for _, t := range s.listaTorneos {
		if t.Region == req.Region {
			lista.Torneos = append(lista.Torneos, &pb.Torneo{
				Id:     t.Id,
				Region: t.Region,
				Estado: pb.EstadoTorneo(t.Estado),
			})
		}
	}
	fmt.Println("    [ OK ] Torneos encontrados en región:", req.Region, "=>", len(lista.Torneos))
	return &lista, nil
}

func (s *ServidorLCP) InscribirEnTorneo(ctx context.Context, req *pb.InscripcionTorneo) (*pb.CodigoRespuesta, error) {
	fmt.Println("[InscribirEnTorneo] gRPC ejecutado\n    Datos recibidos:", req)

	// Validar datos
	if req.EntrenadorId == "" || req.TorneoId == "" {
		fmt.Println("    [ ! ] ID de entrenador o torneo vacío")
		return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "ID de entrenador o torneo vacío"}, nil
	}

	// Verificar que el entrenador existe
	var entrenadorExistente bool
	for _, e := range *s.listaEntrenadores {
		if e.Id == req.EntrenadorId {
			entrenadorExistente = true
			break
		}
	}
	if !entrenadorExistente {
		fmt.Println("    [ ! ] Entrenador no encontrado")
		return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "Entrenador no encontrado"}, nil
	}

	// Verificar que el torneo existe
	var torneoExistente bool
	for _, t := range s.listaTorneos {
		if t.Id == req.TorneoId {
			torneoExistente = true
			if t.Estado != types.TORNEO_ACTIVO {
				fmt.Println("    [ ! ] Torneo no activo")
				return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "Torneo no activo"}, nil
			}
			break
		}
	}
	if !torneoExistente {
		fmt.Println("    [ ! ] Torneo no encontrado")
		return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "Torneo no encontrado"}, nil
	}

	// Verificar si el entrenador ya está inscrito en el torneo
	for _, insc := range s.listaInscripciones {
		if insc.EntrenadorID == req.EntrenadorId && insc.TorneoID == req.TorneoId {
			fmt.Println("    [ ! ] Entrenador ya inscrito en el torneo")
			return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "Entrenador ya inscrito en el torneo"}, nil
		}
	}

	// Verificar que el entrenador no esté inscrito en otro torneo activo o pendiente
	for _, insc := range s.listaInscripciones {
		if insc.EntrenadorID == req.EntrenadorId {
			// Buscar el torneo de esta inscripción
			for _, t := range s.listaTorneos {
				if t.Id == insc.TorneoID && (t.Estado == types.TORNEO_ACTIVO || t.Estado == types.TORNEO_PENDIENTE) {
					fmt.Println("    [ ! ] Entrenador ya inscrito en otro torneo activo o pendiente")
					return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "Entrenador ya inscrito en otro torneo activo o pendiente"}, nil
				}
			}
		}
	}

	// Verificar regiones de entrenador y torneo
	var torneoRegion string
	for _, t := range s.listaTorneos {
		if t.Id == req.TorneoId {
			torneoRegion = t.Region
			break
		}
	}
	if torneoRegion != "" {
		for _, e := range *s.listaEntrenadores {
			if e.Id == req.EntrenadorId {
				if e.Region != torneoRegion {
					fmt.Println("    [ ! ] Entrenador no pertenece a la región del torneo")
					return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "Entrenador no pertenece a la región del torneo"}, nil
				}
				break
			}
		}
	} else {
		fmt.Println("    [ ! ] Región del torneo no encontrada")
		return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "Región del torneo no encontrada"}, nil
	}

	//Verificar si el jugador está inactivo o suspendido
	for i, e := range *s.listaEntrenadores {
		if e.Id == req.EntrenadorId {
			if e.Estado == types.ENTRENADOR_INACTIVO {
				fmt.Println("    [ ! ] Entrenador Expulsado")
				return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "Entrenador Expulsado"}, nil
			} else if e.Estado == types.ENTRENADOR_SUSPENDIDO {
				fmt.Println("    [ ! ] Entrenador Suspendido")
				// Reducir la suspensión en 1 en la lista de entrenadores
				(*s.listaEntrenadores)[i].Suspencion--
				if (*s.listaEntrenadores)[i].Suspencion == 0 {
					(*s.listaEntrenadores)[i].Estado = types.ENTRENADOR_ACTIVO
					fmt.Println("    [ OK ] Entrenador reactivado tras finalizar suspensión")
					return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "Entrenador reactivado tras finalizar suspensión"}, nil
				} else {
					fmt.Println("    [ OK ] Suspensión reducida, quedan", (*s.listaEntrenadores)[i].Suspencion, "torneos")
					return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "Suspensión reducida, quedan " + fmt.Sprint((*s.listaEntrenadores)[i].Suspencion) + " torneos"}, nil
				}
			} else {
				fmt.Println("    [ OK ] Entrenador activo")
				break
			}
		}
	}

	// Agregar inscripción
	nuevaInscripcion := types.InscripcionTorneo{
		EntrenadorID: req.EntrenadorId,
		TorneoID:     req.TorneoId,
	}
	s.listaInscripciones = append(s.listaInscripciones, nuevaInscripcion)
	fmt.Println("    [OK] Inscripción agregada:", nuevaInscripcion)

	// Buscar inscripciones para el mismo torneo
	var entrenadoresInscritos []string
	for _, insc := range s.listaInscripciones {
		if insc.TorneoID == req.TorneoId {
			entrenadoresInscritos = append(entrenadoresInscritos, insc.EntrenadorID)
		}
	}

	if len(entrenadoresInscritos) == 2 {
		// Buscar datos de los entrenadores
		var entrenador1, entrenador2 *types.Entrenador
		for i := range *s.listaEntrenadores {
			if (*s.listaEntrenadores)[i].Id == entrenadoresInscritos[0] {
				entrenador1 = &(*s.listaEntrenadores)[i]
			}
			if (*s.listaEntrenadores)[i].Id == entrenadoresInscritos[1] {
				entrenador2 = &(*s.listaEntrenadores)[i]
			}
		}
		if entrenador1 != nil && entrenador2 != nil {
			// Llamar a AsignarCombate en el Gym usando la función de clienteGym.go
			err := gymAsignarCombate(s.gymClient, entrenador1, entrenador2, req.TorneoId, entrenador1.Region)
			if err != nil {
				fmt.Println("    [ ! ] Error al asignar combate:", err)
				return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "Error al asignar combate en el Gym"}, nil
			}
			fmt.Println("    [OK] Combate asignado en el Gym")
		}
		// Cerrar inscripción del torneo
		for i := 0; i < len(s.listaTorneos); i++ {
			if s.listaTorneos[i].Id == req.TorneoId {
				s.listaTorneos[i].Estado = types.TORNEO_PENDIENTE // Cambiar estado a pendiente
				fmt.Println("    [OK] Torneo pendiente:", s.listaTorneos[i])
				break
			}
		}
	}

	// imprimir lista de inscripciones
	fmt.Println("    [ OK ] Lista de inscripciones actualizada:")
	for _, insc := range s.listaInscripciones {
		fmt.Printf("        - Entrenador ID: %s, Torneo ID: %s\n", insc.EntrenadorID, insc.TorneoID)
	}

	return &pb.CodigoRespuesta{Codigo: 0, Mensaje: "Inscripción exitosa"}, nil
}
