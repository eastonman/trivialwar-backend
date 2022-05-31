package main

import (
	"log"
	"net"
	"net/http"

	"github.com/eastonman/trivialwar-backend/app/game"
	"github.com/eastonman/trivialwar-backend/app/user"
	"github.com/google/uuid"
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Client connected: %s", r.RemoteAddr)
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	userIP := net.ParseIP(ip)

	// Then create a user instance
	user := user.User{
		UUID:                 uuid.New(),
		IP:                   userIP,
		WsConn:               conn,
		IsMultiplayerPlaying: false,
		Timer:                nil,
	}

	game.Game.AddUser(&user)
}
