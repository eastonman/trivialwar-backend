package user

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/eastonman/trivialwar-backend/app/model"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type User struct {
	// Global unique identifier
	UUID uuid.UUID `json:"uuid"`

	// Access from this IP
	IP net.IP

	// WebSocket connection to user
	WsConn *websocket.Conn `json:"-"`

	// Indicates if is playing and in multiplayer-mode
	IsMultiplayerPlaying bool

	// PairUser is the user that this user playing with
	PairUser *User

	// User score
	Score uint64 `json:"score"`

	Timer *time.Ticker `json:"-"`
}

func (u *User) ReportScore(ctx context.Context) {
	for {
		// Send score to user

		// Prepare
		data := strconv.FormatInt(int64(u.PairUser.Score), 10)
		clientPacket := model.ClientPacket{
			Type: model.ReportScore,
			Data: data,
		}
		clientPacketRaw, _ := json.Marshal(clientPacket)

		// Send
		if err := u.WsConn.WriteMessage(websocket.TextMessage, clientPacketRaw); err != nil {
			log.Println("Error during websocket write: ", err)
			return
		}

		select {

		// If stop signal received, stop the goroutine
		case <-ctx.Done():
			return
		// else continue to send score
		default:
			continue
		}

	}
}
