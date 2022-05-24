package user

import (
	"encoding/json"
	"log"
	"net"

	scoreboard "github.com/eastonman/trivialwar-backend/app/scoreBoard"
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
			log.Println("User struct received: ", string(message))
			message, _ = json.Marshal(scoreboard.ScoreBoard.Entries)
			err = u.WsConn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("Error during message writing", err)
				break
			}
			log.Println("User struct sent: ", string(message))
		}
	}()
}
