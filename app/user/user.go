package user

import (
	"net"
	"time"

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
