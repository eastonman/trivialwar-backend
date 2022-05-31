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

	// LAN IP for WebRTC
	IP net.IP

	// WebSocket connection to user
	WsConn *websocket.Conn

	// Indicates if is playing and in multiplayer-mode
	IsMultiplayerPlaying bool

	Timer *time.Ticker
}
