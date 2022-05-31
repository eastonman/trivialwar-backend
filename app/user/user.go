package user

import (
	"encoding/json"
	"log"
	"net"

	"github.com/eastonman/trivialwar-backend/app/game"
	"github.com/eastonman/trivialwar-backend/app/model"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type User struct {
	// Global unique identifier
	UUID uuid.UUID `json:"uuid"`

	// LAN IP for WebRTC
	IP net.Addr

	// WebSocket connection to user
	WsConn *websocket.Conn
}

// HandleConn is a non blocking function that start a goroutine to handle further WebSocket messages
func (u *User) HandleConn() {
	// Start a goroutine
	go func() {
		for {
			_, message_raw, err := u.WsConn.ReadMessage()
			if err != nil {
				log.Println("Error during message reading:", err)
				break
			}
			log.Println("User struct received: ", string(message_raw))

			var command model.Command
			if err := json.Unmarshal(message_raw, &command); err != nil {
				log.Println("Error during message json unmarshal:", err)
				break
			}

			switch command.Type {
			case model.GetAllUsers:
				message, _ := json.Marshal(game.Game.GetAllUsers())
				if err := u.WsConn.WriteMessage(websocket.TextMessage, message); err != nil {
					log.Println("Error during websocket write: ", err)
					break
				}
			}

		}
	}()
}
