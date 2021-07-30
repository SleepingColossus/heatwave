extends Node

class_name MessageType

enum ClientMessageType {
	JOIN_GAME = 0
	LEAVE_GAME = 1
	MOVE = 2
	SHOOT = 3
}

enum ServerMessageTypes {
	NOTIFICATION = 0
	SELF_CONNECTED = 1          # you have joined the game
	PLAYER_CONNECTED = 2        # friendly player has joined the game
	PLAYER_DISCONNECTED = 3
	ACTOR_MOVED = 4
	ENEMY_SPAWNED = 5
	PROJECTILE_SPAWNED = 6
}
