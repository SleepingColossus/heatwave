package main

import (
	"github.com/SleepingColossus/heatwave/game"
	"github.com/gorilla/websocket"
)

// client message types
const (
	JoinGame int = iota
	LeaveGame
	Move
	Shoot
)

// server message types
const (
	InitNewPlayer int = iota
	GameStateUpdated
)

// message type received from game via websocket
type ClientMessage struct {
	MessageType int               `json:"messageType"`
	MessageBody map[string]string `json:"messageBody"`
}

// message type sent to game via websocket
type ServerMessage struct {
	MessageType int                   `json:"messageType"`
	MessageBody game.GameStateUpdate  `json:"messageBody"`
}

func newServerMessage(msgType int, msgBody game.GameStateUpdate) ServerMessage {
	return ServerMessage{
		MessageType: msgType,
		MessageBody: msgBody,
	}
}

// message type sent to send and broadcast channels
type ChannelMessage struct {
	Message     ServerMessage
	Connection  *websocket.Conn
}

func newChannelMessage(msg ServerMessage, conn *websocket.Conn) *ChannelMessage {
	return &ChannelMessage{
		Message:     msg,
		Connection:  conn,
	}
}
