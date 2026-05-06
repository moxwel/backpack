package types

// CDN

type ResultadoCombate struct {
	IdTorneo          string
	IdEntrenador1     string
	NombreEntrenador1 string
	IdEntrenador2     string
	NombreEntrenador2 string
	IdGanador         string
	NombreGanador     string
	Fecha             string
	TipoMensaje       string
	Duracion          int32
}

type FalloCombate struct {
	IdTorneo          string
	IdEntrenador1     string
	NombreEntrenador1 string
	IdEntrenador2     string
	NombreEntrenador2 string
	Fecha             string
	TipoMensaje       string
}

// SDN

type RankingActualizado struct {
	IdEntrenador     string
	NombreEntrenador string
	NuevoRanking     int32
	Fecha            string
	TipoMensaje      string
}

type NuevoTorneo struct {
	IdTorneo    string
	Region      string
	Estado      bool
	Fecha       string
	TipoMensaje string
}

type PenalizacionEntrenador struct {
	IdEntrenador     string
	NombreEntrenador string
	Penalizacion     int32
	Fecha            string
	TipoMensaje      string
}

type ConfirmarcionInscripcion struct {
	IdTorneo         string
	IdEntrenador     string
	NombreEntrenador string
	Fecha            string
	TipoMensaje      string
}

type AlertaGenerica struct {
	Mensaje     string
	Fecha       string
	TipoMensaje string
}

type NotificacionGenerica struct {
	TipoMensaje string
	Payload     string
}

// GYM

type Gym struct {
	IdGym  string
	Region string
	Estado bool
	Llave  string
}

type CombateGymMessage struct {
	Region     string
	MsgCifrado string
}

// Entrenadores

type RegistroCombate struct {
	IdEntrenador   string
	Nombre         string
	IdTorneo       string
	Resultado      string
	RankingAntes   int32
	RankingDespues int32
}

//LCP

type Penalizaciones struct {
	IdEntrenador string
	Penalizacion int32
}
