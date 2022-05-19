package game

import (
	"github.com/eastonman/trivialwar-backend/app/user"
)

var Game = &game{}

type game struct {
	// User HashMap, UUID -> User struct
	Users map[string]*user.User
}

func (g *game) GetAllUsers() map[string]*user.User {
	return g.Users
}

func (g *game) AddUser(user *user.User) {
	g.Users[user.UUID.String()] = user
}
