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

func (g *game) AddUser(u *user.User) {
	if g.Users == nil {
		g.Users = make(map[string]*user.User)
	}
	g.Users[u.UUID.String()] = u
}
