package types

// PlayerStatus enum
type PlayerStatus int32

const (
	IDLE     PlayerStatus = 0
	IN_QUEUE PlayerStatus = 1
	IN_MATCH PlayerStatus = 2
	UNKNOWN  PlayerStatus = 3
)

// ServerStatus enum
type ServerStatus int32

const (
	DISPONIBLE  ServerStatus = 0
	OCUPADO     ServerStatus = 1
	CAIDO       ServerStatus = 2
	DESCONOCIDO ServerStatus = 3
)

// VectorClockEntry representa una entrada del reloj vectorial
type VectorClockEntry struct {
	ServerId  string
	Timestamp int32
}

// VectorClock representa el reloj vectorial completo
type VectorClock struct {
	Entries []VectorClockEntry
}

// PlayerInfoRequest
type PlayerInfoRequest struct {
	PlayerId       string
	ModePreference int32
	VectorClock    VectorClockMap
}

// QueuePlayerResponse
type QueuePlayerResponse struct {
	StatusCode  int32
	Mensaje     string
	VectorClock VectorClockMap
}

// PlayerStatusRequest
type PlayerStatusRequest struct {
	PlayerId    string
	VectorClock VectorClockMap
}

// PlayerStatusResponse
type PlayerStatusResponse struct {
	Status        PlayerStatus
	MatchId       string
	ServerAddress string
	VectorClock   VectorClockMap
}

// CancelQueueRequest
type CancelQueueRequest struct {
	PlayerId    string
	VectorClock VectorClockMap
}

// CancelQueueResponse
type CancelQueueResponse struct {
	StatusCode  int32
	Message     string
	VectorClock VectorClockMap
}

// AssignMatchRequest
type AssignMatchRequest struct {
	Id          string
	PlayerId1   string
	PlayerId2   string
	VectorClock VectorClockMap
}

// AssignMatchResponse
type AssignMatchResponse struct {
	StatusCode  int32
	Message     string
	VectorClock VectorClockMap
}

// PingServerRequest
type PingServerRequest struct {
	VectorClock VectorClockMap
}

// PingServerResponse
type PingServerResponse struct {
	StatusCode  int32
	Message     string
	VectorClock VectorClockMap
}

// ServerStatusUpdateRequest
type ServerStatusUpdateRequest struct {
	ServerId      string
	NewStatus     ServerStatus
	ServerAddress string
	VectorClock   VectorClockMap
}

// ServerStatusUpdateResponse
type ServerStatusUpdateResponse struct {
	StatusCode  int32
	Message     string
	VectorClock VectorClockMap
}

// AdminRequest
type AdminRequest struct {
	VectorClock VectorClockMap
}

// ServerState
type ServerState struct {
	Id             string
	Status         ServerStatus
	Address        string
	CurrentMatchId string
	VectorClock    VectorClockMap
}

// PlayerQueueEntry
type PlayerQueueEntry struct {
	PlayerId    string
	TimeInQueue int32
	VectorClock VectorClockMap
}

// SystemStatusResponse
type SystemStatusResponse struct {
	ServerList  []ServerState
	PlayerQueue []PlayerQueueEntry
	VectorClock VectorClockMap
}

// AdminServerUpdateRequest
type AdminServerUpdateRequest struct {
	ServerId    string
	NewStatus   ServerStatus
	VectorClock VectorClockMap
}

// AdminUpdateResponse
type AdminUpdateResponse struct {
	StatusCode  int32
	Message     string
	VectorClock VectorClockMap
}

// UpdateGameServerStateRequest - MATCHMAKER -> GAME SERVER
type UpdateGameServerStateRequest struct {
	ServerId    string
	NewStatus   ServerStatus
	VectorClock VectorClockMap
}

// UpdateGameServerStateResponse
type UpdateGameServerStateResponse struct {
	StatusCode  int32
	Message     string
	VectorClock VectorClockMap
}

// GetMatchResultRequest
type GetMatchResultRequest struct {
	MatchId     string
	VectorClock VectorClockMap
}

// GetMatchResultResponse
type GetMatchResultResponse struct {
	StatusCode  int32
	Message     string
	WinnerId    string
	LoserId     string
	VectorClock VectorClockMap
}
