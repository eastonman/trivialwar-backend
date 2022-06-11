package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"

	"github.com/eastonman/trivialwar-backend/app/game"
	"github.com/eastonman/trivialwar-backend/app/model"
	"github.com/eastonman/trivialwar-backend/app/user"
	"github.com/gorilla/websocket"
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Client connected: %s", r.RemoteAddr)
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}

	// Read message from user
	_, message_raw, err := conn.ReadMessage()
	if err != nil {
		log.Print("Error during websocket reading:", err)
		return
	}

	var command model.Command
	if err := json.Unmarshal(message_raw, &command); err != nil {
		// If error parsing command, only log
		log.Println("Error during message json unmarshal:", err)
		return
	}

	switch command.Type {
	case model.Login:
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		userIP := net.ParseIP(ip)

		// Resolve LoginInfo
		var loginInfo model.LoginInfo
		if err := json.Unmarshal([]byte(command.Param), &loginInfo); err != nil {
			// If error parsing command, only log
			log.Println("Error during message json unmarshal:", err)
			return
		}

		// If user existes
		if user, ok := game.Game.Users[loginInfo.Username]; ok {
			// Check password
			if loginInfo.Hash != user.Password {
				message := model.ClientPacket{
					Type: model.Login,
					Data: "denied",
				}
				message_raw, _ = json.Marshal(message)
				if err := conn.WriteMessage(websocket.TextMessage, message_raw); err != nil {
					log.Printf("error during websocket write: %s", err)
					return
				}
				return
			}

			// If passed, create a goroutine to handle the reset of communication
			message := model.ClientPacket{
				Type: model.Login,
				Data: "passed",
			}
			message_raw, _ = json.Marshal(message)
			if err := conn.WriteMessage(websocket.TextMessage, message_raw); err != nil {
				log.Printf("error during websocket write: %s", err)
				return
			}
			user.WsConn = conn
			user.Online = true
			game.Game.HandleConn(user)
			return
		}

		// If not, create a user instance
		log.Printf("%s", loginInfo.Username)
		user := user.User{
			Username:             loginInfo.Username,
			Password:             loginInfo.Hash,
			IP:                   userIP,
			WsConn:               conn,
			IsMultiplayerPlaying: false,
			PairUser:             nil,
			Score:                0,
			Timer:                nil,
		}
		message := model.ClientPacket{
			Type: model.Login,
			Data: "passed",
		}
		message_raw, _ = json.Marshal(message)
		if err := conn.WriteMessage(websocket.TextMessage, message_raw); err != nil {
			log.Printf("error during websocket write: %s", err)
			return
		}
		game.Game.AddUser(&user)
		game.Game.HandleConn(&user)
	default:
		return
	}

}
