extends Node

# server settings
export var server_port = 8080
export var server_address = "localhost"
var server_url = "ws://%s:%s/ws" % [server_address, server_port]

var actor_manager
var direction: Vector2
var ws_client : WebSocketClient

# unique identifier
# used for player specific actions
# assigned after successful connection request
var client_id: String = ""

func _ready():
	actor_manager = get_node("/root/main/actor_manager")
	direction = Vector2(0, 0)

	ws_client = WebSocketClient.new()

	# setup websocket message handlers
	ws_client.connect("connection_closed", self, "_on_closed")
	ws_client.connect("connection_error", self, "_on_closed")
	ws_client.connect("connection_established", self, "_on_connected")
	ws_client.connect("data_received", self, "_on_data")

	var err = ws_client.connect_to_url(server_url)
	if err != OK:
		push_error("Unable to connect")
		set_process(false)

func _process(_delta):
	ws_client.poll()

	var new_direction = poll_inputs()

	# is the direction the same as last time
	if direction.x == new_direction.x && direction.y == new_direction.y:
		# yes - do nothing
		pass
	else:
		# no - send message to server
		var move_body = {
			"clientId": client_id,
			"x": new_direction.x as String,
			"y": new_direction.y as String,
		}

		var move_message = create_message(MessageType.ClientMessageType.MOVE, move_body)

		print_debug(move_message)

		var packet: PoolByteArray = JSON.print(move_message).to_utf8()
		ws_client.get_peer(1).put_packet(packet)

		# update last know direction
		direction = new_direction

func poll_inputs() -> Vector2:

	var new_direction = Vector2(0, 0)

	# horizontal
	if Input.is_action_pressed("move_left"):
		new_direction.x = -1
	elif Input.is_action_pressed("move_right"):
		new_direction.x = 1
	else:
		new_direction.x = 0

	# vertical
	if Input.is_action_pressed("move_up"):
		new_direction.y = -1
	elif Input.is_action_pressed("move_down"):
		new_direction.y = 1
	else:
		new_direction.y = 0

	return new_direction

func _on_connected(_protocol: String):
	var msg_body = Dictionary()
	var join_message = create_message(MessageType.ClientMessageType.JOIN_GAME, msg_body)

	# convert to JSON
	var packet: PoolByteArray = JSON.print(join_message).to_utf8()

	# send message to server
	ws_client.get_peer(1).put_packet(packet)

func _on_closed(was_clean = false):
	send_disconnect_message()
	set_process(false)

func _on_data():
	var packet = ws_client.get_peer(1).get_packet()
	var message = decode_message(packet)
	handle_message(message)

# Gracefully disconnect on quit
func _notification(what):
	if what == MainLoop.NOTIFICATION_WM_QUIT_REQUEST:
		print_debug("quitting game")
		send_disconnect_message()

func send_disconnect_message():
	print_debug("disconnect message start")

	if client_id == "":
		push_error("client id is blank")
		return

	var leave_body = {
		"clientId": client_id
	}

	var dc_msg = create_message(MessageType.ClientMessageType.LEAVE_GAME, leave_body)
	var packet: PoolByteArray = JSON.print(dc_msg).to_utf8()
	ws_client.get_peer(1).put_packet(packet)

	print_debug("disconnect message end")

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
		var msg_body : Array = msg["messageBody"]

		match msg_type:
			MessageType.ServerMessageTypes.SELF_CONNECTED:
				var body = msg_body[0]

				# create a new player actor
				var actor_id = body["clientId"]
				var actor_type = body["actorType"] as int # TODO error handling around type cast
				var x = body["x"] as int # TODO error handling around type cast
				var y = body["y"] as int # TODO error handling around type cast

				# spawn self
				var position = Vector2(x, y)
				actor_manager.create_actor(actor_id, actor_type, position)

				# set client id for use in future messages
				set_client_id(actor_id)

			MessageType.ServerMessageTypes.PLAYER_CONNECTED:
				for body in msg_body:
					var actor_id = body["clientId"]

					if actor_id == client_id:
						print_debug("received PLAYER_CONNECTED for self. pass")
						return

					var actor_type = body["actorType"] as int # TODO error handling around type cast
					var x = body["x"] as int # TODO error handling around type cast
					var y = body["y"] as int # TODO error handling around type cast

					# spawn friendly player
					var position = Vector2(x, y)
					actor_manager.create_actor(actor_id, actor_type, position)

			MessageType.ServerMessageTypes.PLAYER_DISCONNECTED:
				for body in msg_body:
					var actor_id = body["clientId"]
					actor_manager.delete_actor(actor_id)

			MessageType.ServerMessageTypes.ACTOR_MOVED:
				for body in msg_body:
					var actor_id = body["clientId"]
					var position_x = body["positionX"] as int
					var position_y = body["positionY"] as int

					var position = Vector2(position_x, position_y)

					actor_manager.move_actor(actor_id, position)

	print_debug("handle message end")

func set_client_id(id: String):
	print_debug("setting client id to %s" % id)
	client_id = id
