package main

import (
	"log"
	"net/http"

	"github.com/eastonman/trivialwar-backend/app/game"
	"github.com/eastonman/trivialwar-backend/app/user"
	"github.com/google/uuid"
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}

	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("Error during message reading:", err)
		return
	}
	log.Printf("Received: %s", message)

	// Then create a user instance
	user := user.User{
		UUID:   uuid.New(),
		IP:     nil,
		WsConn: conn,
	}

	game.Game.AddUser(&user)

	// Handle the connection to user struct
	user.HandleConn()
}
