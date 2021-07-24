extends Node

# server settings
export var server_port = 8080
export var server_address = "localhost"
var server_url = "ws://%s:%s/ws" % [server_address, server_port]

var actor_manager
var ws_client : WebSocketClient

func _ready():
	actor_manager = get_node("/root/main/actor_manager")
	
	ws_client = WebSocketClient.new()
	
	# Connect base signals to get notified of connection open, close, and errors.
	ws_client.connect("connection_closed", self, "_on_closed")
	ws_client.connect("connection_error", self, "_on_closed")
	ws_client.connect("connection_established", self, "_on_connected")

	# This signal is emitted when not using the Multiplayer API every time
	# a full packet is received.
	# Alternatively, you could check get_peer(1).get_available_packets() in a loop.
	ws_client.connect("data_received", self, "_on_data")

	# Initiate connection to the given URL.
	var err = ws_client.connect_to_url(server_url)
	if err != OK:
		print("Unable to connect")
		set_process(false)
		
	var actor_type = 0 #ActorType.PLAYER
	actor_manager.create_actor("test-123", actor_type, Vector2(1000, 600))

func _on_connected(_protocol: String):
	var msg_body = Dictionary()
	var join_message = create_message(ClientEventType.JOIN_GAME, msg_body)
	
	# convert to JSON
	var packet: PoolByteArray = JSON.print(join_message).to_utf8()
	
	# send message to server
	ws_client.get_peer(1).put_packet(packet)

func _on_closed(was_clean = false):
	# was_clean will tell you if the disconnection was correctly notified
	# by the remote peer before closing the socket.
	print("Closed, clean: ", was_clean)
	set_process(false)

func _on_data():
	# Print the received packet, you MUST always use get_peer(1).get_packet
	# to receive data from server, and not get_packet directly when not
	# using the MultiplayerAPI.
	var data = ws_client.get_peer(1).get_packet().get_string_from_utf8()
	print("Got data from server: ", data)

# Called every frame. 'delta' is the elapsed time since the previous frame.
func _process(_delta):
	ws_client.poll()

# ---

enum ClientEventType {
	JOIN_GAME = 0,
}

# ---

enum ServerEventTypes {
	PLAYER_CONNECTED = 0
	PLAYER_DISCONNECTED = 1
}

# ---

func create_message(msg_type, msg_body: Dictionary) -> Dictionary:
	return {
		messageType = msg_type,
		messageBody = msg_body,
	}
