package user

import (
	"log"
	"net"

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
			_, message, err := u.WsConn.ReadMessage()
			if err != nil {
				log.Println("Error during message reading:", err)
				break
			}
			log.Println("User struct received: ", message)
		}
	}()
}
