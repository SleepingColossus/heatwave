class_name MessageType

enum ClientMessageType {
	JOIN_GAME = 0
	LEAVE_GAME = 1
	MOVE = 2
	SHOOT = 3
}

enum ServerMessageTypes {
	INIT_NEW_PLAYER = 0
	GAME_STATE_UPDATED = 1
}
