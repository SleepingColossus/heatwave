package main

// client message types
const(
	JoinGame int = iota
)

// server message types
const (
	PlayerConnected int = iota
	PlayerDisconnected
)

type Message struct {
	MessageType int `json:"messageType"`
	MessageBody map[string]string `json:"messageBody"`
}

func newMessage(msgType int, msgBody map[string]string) *Message {
	return &Message{
		MessageType: msgType,
		MessageBody: msgBody,
	}
}
