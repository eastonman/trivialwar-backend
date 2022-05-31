package model

const (
	GetAllUsers int = iota
	JoinUser
	StartMultiplayerGame
	ReportScore
	GetLeaderBoard
	Bye
)

type Command struct {
	Type  int `json:"type"`
	Param int `json:"param"`
}

type ClientPacket struct {
	Type int    `json:"type"`
	Data []byte `json:"data"`
}
