package sg_api

import (
	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type Instance struct {
	Endpoint string
}

type WsConn struct {
	Id uuid.UUID
	conn *websocket.Conn
}

