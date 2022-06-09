package game

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

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
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if err := g.messageLoop(u, ctx); err != nil {
			log.Printf("User %s message loop exited with error: %s", u.UUID.String(), err)
		}
	}()
}

func (g *game) messageLoop(u *user.User, ctx context.Context) error {
	for {
		// Read message from user
		_, message_raw, err := u.WsConn.ReadMessage()

		// If is a stop message, do clean up jobs
		if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
			err = fmt.Errorf("user %s exited normally", u.IP.String())
			return err
		} else if err != nil {
			log.Println("Error during message reading:", err)
			continue
		}

		// DEBUG
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
				err = fmt.Errorf("error during websocket write: %s", err)
				return err
			}
		case model.GetLeaderBoard:
			message, _ := json.Marshal(scoreboard.ScoreBoard.Entries)
			packet, _ := json.Marshal(model.ClientPacket{
				Type: model.GetLeaderBoard,
				Data: string(message),
			})
			if err := u.WsConn.WriteMessage(websocket.TextMessage, packet); err != nil {
				err = fmt.Errorf("error during websocket write: %s", err)
				return err
			}
		case model.ReportScore:
			if u.Score, err = strconv.ParseUint(command.Param, 10, 64); err != nil {
				log.Println("Error parsing score: ", err)
				continue
			}
		case model.JoinUser:
			if command.Param != "0" {
				u.PairUser = g.Users[command.Param]
			} else {
				// If not a valid user uuid, pair with the user himself
				u.PairUser = u
			}

		case model.StartMultiplayerGame:
			log.Printf("User %s started multiplayer game", u.UUID.String())
			// Start a multiplayer game

			// Setup a score report goroutine
			go u.ReportScore(ctx)
		}
	}
}
