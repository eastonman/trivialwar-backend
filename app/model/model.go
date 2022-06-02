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
	Type  int    `json:"type"`
	Param string `json:"param"`
}

type ClientPacket struct {
	Type int    `json:"type"`
	Data string `json:"data"`
}
