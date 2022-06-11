package model

const (
	GetAllUsers int = iota
	JoinUser
	StartMultiplayerGame
	ReportScore
	GetLeaderBoard
	Login
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

type LoginInfo struct {
	Username string `json:"username"`
	Hash     string `json:"hash"` // should be SHA-256
}
