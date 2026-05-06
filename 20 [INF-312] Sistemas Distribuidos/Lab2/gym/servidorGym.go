package main

import (
	"context"
	"fmt"
	"l2/types"
	"math"
	"math/rand/v2"
	"os"
	"time"

	pb "gym/PokemonProto"

	"github.com/streadway/amqp"
)

func logWithTimestamp(format string, args ...interface{}) {
	now := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf(format, args...)
	logMsg := fmt.Sprintf("[%s] %s\n", now, msg)
	fmt.Print(logMsg)
	f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		defer f.Close()
		f.WriteString(logMsg)
	}
}

type ServidorGym struct {
	pb.UnimplementedServicioGymServer
	gimnasios          []types.Gym
	rabbit_ch          *amqp.Channel
	combatesReportados []string
}

func (s *ServidorGym) AsignarCombate(ctx context.Context, req *pb.CombateGym) (*pb.CodigoRespuesta, error) {
	logWithTimestamp("[AsignarCombate] gRPC ejecutado\n    Datos recibidos: %+v", req)
	combate := types.CombateGym{
		Id: req.Id,
		Entrenador1: types.Entrenador{
			Id:         req.GetEntrenador_1().Id,
			Nombre:     req.GetEntrenador_1().Nombre,
			Region:     req.GetEntrenador_1().Region,
			Ranking:    req.GetEntrenador_1().Ranking,
			Estado:     types.ENTRENADOR_ACTIVO,
			Suspencion: 0,
		},
		Entrenador2: types.Entrenador{
			Id:         req.GetEntrenador_2().Id,
			Nombre:     req.GetEntrenador_2().Nombre,
			Region:     req.GetEntrenador_2().Region,
			Ranking:    req.GetEntrenador_2().Ranking,
			Estado:     types.ENTRENADOR_ACTIVO,
			Suspencion: 0,
		},
		Region: req.Region,
	}
	logWithTimestamp("    [AsignarCombate] Combate recibido: %+v", combate)

	for _, id := range s.combatesReportados {
		if id == combate.Id {
			logWithTimestamp("    [AsignarCombate] ERROR: Combate duplicado detectado, no se reportará nuevamente.")
			return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "ERROR_COMBATE_REPETIDO"}, nil
		}
	}

	llave := ""
	for _, gym := range s.gimnasios {
		if gym.Region == combate.Region {
			llave = gym.Llave
			break
		}
	}
	if llave == "" {
		logWithTimestamp("    [AsignarCombate] Error: Región desconocida para llave de cifrado")
		return &pb.CodigoRespuesta{Codigo: 1, Mensaje: "ERROR_REGION"}, nil
	}

	// Procesar combate en segundo plano
	go func() {
		duracion := rand.IntN(20) + 1
		var id_ent_ganador string
		for i, gym := range s.gimnasios {
			if gym.Estado && (gym.Region == combate.Region || combate.Region == "") {
				s.gimnasios[i].Estado = false
				logWithTimestamp("    [AsignarCombate] Combate asignado a gimnasio: %s", gym.Region)
				id_ent_ganador = SimularCombate(combate.Entrenador1, combate.Entrenador2, duracion)
				s.gimnasios[i].Estado = true
				break
			}
		}
		if id_ent_ganador == "" {
			logWithTimestamp("    [AsignarCombate] Error: No hay gimnasios disponibles para el combate")
			return
		}
		var ent_ganador types.Entrenador
		if id_ent_ganador == combate.Entrenador1.Id {
			ent_ganador = combate.Entrenador1
		} else {
			ent_ganador = combate.Entrenador2
		}
		logWithTimestamp("    [AsignarCombate] Entrenador ganador: %+v", ent_ganador)
		resultado := types.ResultadoCombate{
			IdTorneo:          combate.Id,
			IdEntrenador1:     combate.Entrenador1.Id,
			NombreEntrenador1: combate.Entrenador1.Nombre,
			IdEntrenador2:     combate.Entrenador2.Id,
			NombreEntrenador2: combate.Entrenador2.Nombre,
			IdGanador:         ent_ganador.Id,
			NombreGanador:     ent_ganador.Nombre,
			Fecha:             time.Now().Format("2006-01-02"),
			TipoMensaje:       "ResultadoCombate",
			Duracion:          int32(duracion),
		}
		for i := 1; i <= 3; i++ {
			err := EnviarResultado(s.rabbit_ch, resultado, llave, combate.Region)
			if err == nil {
				s.combatesReportados = append(s.combatesReportados, combate.Id)
				logWithTimestamp("    [AsignarCombate] Resultado enviado al CDP: %+v", resultado)
				return
			}
			logWithTimestamp("    [AsignarCombate] Error enviando resultado (intento %d/3): %v", i, err)
			time.Sleep(1 * time.Second)
		}
		logWithTimestamp("    [AsignarCombate] Error: no se pudo enviar el resultado tras 3 intentos")
	}()

	return &pb.CodigoRespuesta{Codigo: 0, Mensaje: "COMBATE_EN_PROGRESO"}, nil
}

func SimularCombate(ent1 types.Entrenador, ent2 types.Entrenador, duracion int) string {
	logWithTimestamp("    [SimularCombate] Simulando batalla, duración: %d segundos", duracion)
	time.Sleep(time.Duration(duracion) * time.Second)
	diff := float64(ent1.Ranking - ent2.Ranking)
	k := 100.0
	prob := 1.0 / (1.0 + math.Exp(-diff/k))
	if rand.Float64() <= prob {
		return ent1.Id
	}
	return ent2.Id
}

func EnviarResultado(ch *amqp.Channel, resultado types.ResultadoCombate, llave string, region string) error {
	logWithTimestamp("[EnviarResultado] Enviando resultado al CDP...")
	var err error
	for i := 1; i <= 3; i++ {
		msg_cifrado, _ := EncriptarResultadoCombate(resultado, llave)
		mensaje_cola := types.CombateGymMessage{
			Region:     region,
			MsgCifrado: msg_cifrado,
		}
		err = SendGymResultQueue(ch, mensaje_cola)
		if err == nil {
			logWithTimestamp("    [EnviarResultado] Resultado enviado al CDP: %+v", resultado)
			return nil
		}
		logWithTimestamp("    [EnviarResultado] Error enviando resultado (intento %d/3): %v", i, err)
		time.Sleep(1 * time.Second)
	}
	// Si falla tras 3 intentos, usar logWithTimestamp
	logWithTimestamp("[EnviarResultado] ERROR FATAL: No se pudo enviar el resultado tras 3 intentos: %v", err)
	return fmt.Errorf("    [ ! ] Error al enviar el resultado al CDP tras 3 intentos: %v", err)
}
