package model

const (
	GetAllUsers int = iota
	JoinUser
	StartMultiplayerGame
	ReportScore
	GetLeaderBoard
	Login
	Signup
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

type SignupInfo struct {
	Username string `json:"username"`
	Hash     string `json:"hash"` // should be SHA-256
}
