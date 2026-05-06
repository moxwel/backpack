package types

import (
	"time"
)

type Server struct {
	Status         ServerStatus
	Address        string
	CurrentMatchId string
	Player1        string
	Player2        string
}

type Player struct {
	Status        PlayerStatus
	ServerAddress string
	MatchId       string
	TimeInQueue   int32
	QueueSince    time.Time
}

type VectorClockMap map[string]int32

type MatchResult struct {
	Winner string
	Loser  string
}

type MatchResult2 struct {
	Winner  string
	Loser   string
	MatchId string
}
