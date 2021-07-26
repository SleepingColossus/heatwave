package main

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

type ClientMessage struct {
	MessageType int               `json:"messageType"`
	MessageBody map[string]string `json:"messageBody"`
}

type ServerMessage struct {
	MessageType int                 `json:"messageType"`
	MessageBody []map[string]string `json:"messageBody"`
}

func newClientMessage(msgType int, msgBody map[string]string) *ClientMessage {
	return &ClientMessage{
		MessageType: msgType,
		MessageBody: msgBody,
	}
}

func newServerMessage(msgType int, msgBody []map[string]string) *ServerMessage {
	return &ServerMessage{
		MessageType: msgType,
		MessageBody: msgBody,
	}
}
