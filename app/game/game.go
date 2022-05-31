package game

import (
	"encoding/json"
	"log"

	"github.com/eastonman/trivialwar-backend/app/model"
	scoreboard "github.com/eastonman/trivialwar-backend/app/scoreBoard"
	"github.com/eastonman/trivialwar-backend/app/user"
	"github.com/gorilla/websocket"
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
	g.HandleConn(u)
}

// HandleConn is a non blocking function that start a goroutine to handle further WebSocket messages
func (g *game) HandleConn(u *user.User) {
	// Start a goroutine
	go func() {
		for {
			_, message_raw, err := u.WsConn.ReadMessage()
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				log.Printf("User %s exited normally", u.IP.String())
				return
			} else if err != nil {
				log.Println("Error during message reading:", err)
				break
			}
			log.Println("User struct received: ", string(message_raw))

			var command model.Command
			if err := json.Unmarshal(message_raw, &command); err != nil {
				// If error parsing command, only log
				log.Println("Error during message json unmarshal:", err)
				continue
			}

			switch command.Type {
			case model.GetAllUsers:
				message, _ := json.Marshal(g.GetAllUsers())
				if err := u.WsConn.WriteMessage(websocket.TextMessage, message); err != nil {
					log.Println("Error during websocket write: ", err)
					return
				}
			case model.GetLeaderBoard:
				message, _ := json.Marshal(scoreboard.ScoreBoard.Entries)
				if err := u.WsConn.WriteMessage(websocket.TextMessage, message); err != nil {
					log.Println("Error during websocket write: ", err)
					return
				}
			}
		}
	}()
}
