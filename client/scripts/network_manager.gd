extends Node

# server settings
export var server_port = 8080
export var server_address = "localhost"
var server_url = "ws://%s:%s/ws" % [server_address, server_port]

var actor_manager
var ws_client : WebSocketClient

# unique identifier
# used for player specific actions
# assigned after successful connection request
var client_id: String

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

func _process(_delta):
	ws_client.poll()

func _on_connected(_protocol: String):
	var msg_body = Dictionary()
	var join_message = create_message(MessageType.ClientMessageType.JOIN_GAME, msg_body)
	
	# convert to JSON
	var packet: PoolByteArray = JSON.print(join_message).to_utf8()
	
	# send message to server
	ws_client.get_peer(1).put_packet(packet)

func _on_closed(was_clean = false):
	send_disconnect_message()
	
	# was_clean will tell you if the disconnection was correctly notified
	# by the remote peer before closing the socket.	
	print("Closed, clean: ", was_clean)
	set_process(false)

func _on_data():
	# Print the received packet, you MUST always use get_peer(1).get_packet
	# to receive data from server, and not get_packet directly when not
	# using the MultiplayerAPI.
	var packet = ws_client.get_peer(1).get_packet()
	var message = decode_message(packet)
	handle_message(message)

# Gracefully disconnect on quit
func _notification(what):
	if what == MainLoop.NOTIFICATION_WM_QUIT_REQUEST:
		send_disconnect_message()

func send_disconnect_message():
	var dc_msg = create_message(MessageType.ClientMessageType.LEAVE_GAME, Dictionary())
	var packet: PoolByteArray = JSON.print(dc_msg).to_utf8()
	ws_client.get_peer(1).put_packet(packet)

func create_message(msg_type, msg_body: Dictionary) -> Dictionary:
	return {
		messageType = msg_type,
		messageBody = msg_body,
	}

func decode_message(packet: PoolByteArray) -> Dictionary:
	print_debug("decode start")
	var msg_str = packet.get_string_from_utf8()
	var parse_result: JSONParseResult = JSON.parse(msg_str)
	
	if parse_result.error == OK:
		print_debug("decode ok")
		return parse_result.result
	else:
		print_debug("decode error")
		push_error(parse_result.error_string)
		return Dictionary()

func handle_message(msg: Dictionary):
	print_debug("handle message start")
	print_debug(msg)
	
	if msg.empty():
		# failed to decode - do nothing
		return
	else:
		var msg_type : int = msg["messageType"]
		var msg_body : Dictionary = msg["messageBody"]
		
		match msg_type:
			MessageType.ServerMessageTypes.PLAYER_CONNECTED:
				# create a new player actor
				var actor_id = msg_body["clientId"]
				var actor_type = msg_body["actorType"] as int # TODO error handling around type cast
				var x = msg_body["x"] as int # TODO error handling around type cast
				var y = msg_body["y"] as int # TODO error handling around type cast
				
				var position = Vector2(x, y)
				
				actor_manager.create_actor(actor_id, actor_type, position)

	print_debug("handle message end")
