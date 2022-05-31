package model

const (
	GetAllUsers int = iota
	JoinUser
	Bye
)

type Command struct {
	Type  int `json:"type"`
	Param int `json:"param"`
}
