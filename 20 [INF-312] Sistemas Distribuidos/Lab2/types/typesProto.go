package types

// Enums equivalentes a EstadoEntrenador y EstadoTorneo

type EstadoEntrenador int32

const (
	ENTRENADOR_ACTIVO     EstadoEntrenador = 0
	ENTRENADOR_INACTIVO   EstadoEntrenador = 1
	ENTRENADOR_SUSPENDIDO EstadoEntrenador = 2
)

type EstadoTorneo int32

const (
	TORNEO_ACTIVO     EstadoTorneo = 0
	TORNEO_PENDIENTE  EstadoTorneo = 1
	TORNEO_FINALIZADO EstadoTorneo = 2
)

// Mensajes equivalentes a los protos

type Entrenador struct {
	Id         string
	Nombre     string
	Region     string
	Ranking    int32
	Estado     EstadoEntrenador
	Suspencion int32
}

type CombateGym struct {
	Id          string
	Entrenador1 Entrenador
	Entrenador2 Entrenador
	Region      string
}

type CodigoRespuesta struct {
	Codigo  int32
	Mensaje string
}

type Torneo struct {
	Id     string
	Region string
	Estado EstadoTorneo
}

type ListaTorneos struct {
	Torneos []Torneo
}

type RegionString struct {
	Region string
}

type InscripcionTorneo struct {
	EntrenadorID string
	TorneoID     string
}

type EntrenadorID struct {
	Id string
}

type ListaEntrenadores struct {
	Entrenadores []Entrenador
}

type Empty struct{}
