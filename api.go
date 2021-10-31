package sg_api

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

func Listen(w *WsConn, cb func(resp string)) error {
	defer w.conn.Close()
	for {
		var data []byte
		_, err := w.conn.Read(data)
		if err != nil {
			return err
		} else {
			cb(string(data))
		}
	}
}

func Send(w *WsConn, message string) error {
	_, err := w.conn.Write([]byte(message))
	if err != nil {
		w.conn.Close()
		return err
	}
	return nil
}

func New(i Instance, u uuid.UUID) (*WsConn, error) {
	endpoint := fmt.Sprintf("%s/%s", i.Endpoint, u)
	conn, err := websocket.Dial(endpoint, "", "http://localhost")
	if err != nil { return nil, err }
	return &WsConn{
		conn: conn,
		Id: u,
	}, nil
}