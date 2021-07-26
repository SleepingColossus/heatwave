package main

import "github.com/gorilla/websocket"

// client message types
const (
	JoinGame int = iota
	LeaveGame
	Move
)

// server message types
const (
	Notification    int = iota
	SelfConnected       // you have joined the game
	PlayerConnected     // friendly player has joined the game
	PlayerDisconnected
	ActorsMoved
)

// message type received from game via websocket
type ClientMessage struct {
	MessageType int               `json:"messageType"`
	MessageBody map[string]string `json:"messageBody"`
}

// message type sent to game via websocket
type ServerMessage struct {
	MessageType int                 `json:"messageType"`
	MessageBody []map[string]string `json:"messageBody"`
}

func newServerMessage(msgType int, msgBody []map[string]string) ServerMessage {
	return ServerMessage{
		MessageType: msgType,
		MessageBody: msgBody,
	}
}

// message type sent to send and broadcast channels
type ServerChannelMessage struct {
	MessageType int
	Message     ServerMessage
	Connection  *websocket.Conn
}

func newChannelMessage(msgType int, msg ServerMessage, conn *websocket.Conn) *ServerChannelMessage {
	return &ServerChannelMessage{
		MessageType: msgType,
		Message:     msg,
		Connection:  conn,
	}
}
