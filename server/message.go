package main

// client message types
const (
	JoinGame int = iota
	LeaveGame
)

// server message types
const (
	Notification    int = iota
	SelfConnected       // you have joined the game
	PlayerConnected     // friendly player has joined the game
	PlayerDisconnected
)

type Message struct {
	MessageType int               `json:"messageType"`
	MessageBody map[string]string `json:"messageBody"`
}

func newMessage(msgType int, msgBody map[string]string) *Message {
	return &Message{
		MessageType: msgType,
		MessageBody: msgBody,
	}
}
