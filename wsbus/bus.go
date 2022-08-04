package wsbus

import "github.com/gorilla/websocket"

type WsClient struct {
	Conn    *websocket.Conn
	MsgChan chan *WsMessage
	// uid  string
}

type WsMessage struct {
	Message  string
	SourceId uint64
	TargetId uint64
}

var ClientMap = make(map[string]*WsClient)
