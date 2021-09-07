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
		DebugLog.error("Unable to connect")
		set_process(false)

func _process(_delta):
	ws_client.poll()

	var new_direction = poll_movement_inputs()

	# is the direction the same as last frame
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

		DebugLog.debug(move_message)

		var packet: PoolByteArray = JSON.print(move_message).to_utf8()
		ws_client.get_peer(1).put_packet(packet)

		# update last know direction
		direction = new_direction

	var shoot_target = poll_action_inputs()

	if shoot_target != null:
		var shoot_body = {
			"clientId": client_id,
			"x": shoot_target.x as String,
			"y": shoot_target.y as String,
		}

		var shoot_message = create_message(MessageType.ClientMessageType.SHOOT, shoot_body)

		var packet: PoolByteArray = JSON.print(shoot_message).to_utf8()
		ws_client.get_peer(1).put_packet(packet)

func poll_movement_inputs() -> Vector2:

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

func poll_action_inputs():
	if Input.is_action_just_pressed("shoot"):
		return get_viewport().get_mouse_position()
	else:
		return null

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
		DebugLog.info("quitting game")
		send_disconnect_message()

func send_disconnect_message():
	DebugLog.debug("disconnect message start")

	if client_id == "":
		DebugLog.error("client id is blank")
		return

	var leave_body = {
		"clientId": client_id
	}

	var dc_msg = create_message(MessageType.ClientMessageType.LEAVE_GAME, leave_body)
	var packet: PoolByteArray = JSON.print(dc_msg).to_utf8()
	ws_client.get_peer(1).put_packet(packet)

	DebugLog.debug("disconnect message end")

func create_message(msg_type, msg_body: Dictionary) -> Dictionary:
	return {
		messageType = msg_type,
		messageBody = msg_body,
	}

func decode_message(packet: PoolByteArray) -> Dictionary:
	DebugLog.debug("decode start")
	var msg_str = packet.get_string_from_utf8()
	var parse_result: JSONParseResult = JSON.parse(msg_str)

	if parse_result.error == OK:
		DebugLog.debug("decode ok")
		return parse_result.result
	else:
		DebugLog.debug("decode error")
		DebugLog.error(parse_result.error_string)
		return Dictionary()

func handle_message(msg: Dictionary):
	DebugLog.debug("handle message start")
	DebugLog.debug(msg)

	if msg.empty():
		# failed to decode - do nothing
		return
	else:
		var msg_type : int = msg["messageType"]
		var msg_body : Dictionary = msg["messageBody"]

		match msg_type:
			MessageType.ServerMessageTypes.INIT_NEW_PLAYER:
				set_client_id(msg_body["clientId"])

				var actor_updates = msg_body["actorUpdates"]

				for a in actor_updates:
					# Since this is our first time joining
					# Create all actors except ones marked for deletion
					if a["state"] != ActorState.ActorState.DELETED:
						var id = a["id"]
						var type = a["type"]
						var position_x = a["position"]["x"]
						var position_y = a["position"]["y"]
						var current_hp = a["currentHealth"]
						var max_hp = a["maxHealth"]

						var pos = Vector2(position_x, position_y)
						actor_manager.create_actor(id, type, pos, max_hp, current_hp, true)

			MessageType.ServerMessageTypes.GAME_STATE_UPDATED:
				var actor_updates = msg_body["actorUpdates"]

				for a in actor_updates:
					# must conver to int
					# otherwise will be considered a float
					# and will fail to match!
					var state = a["state"] as int

					match state:
						ActorState.ActorState.CREATED:
							var id = a["id"]

							# prevent spawning self multiple times by accident
							if id == client_id:
								pass
							else:
								var type = a["type"]
								var position_x = a["position"]["x"]
								var position_y = a["position"]["y"]
								var current_hp = a["currentHealth"]
								var max_hp = a["maxHealth"]

								var pos = Vector2(position_x, position_y)
								actor_manager.create_actor(id, type, pos, max_hp, current_hp)

						ActorState.ActorState.DELETED:
							var id = a["id"]
							actor_manager.delete_actor(id)

						ActorState.ActorState.UPDATED:
							var id = a["id"]
							var type = a["type"]
							var position_x = a["position"]["x"]
							var position_y = a["position"]["y"]
							var direction_x = a["direction"]["x"]
							var direction_y = a["direction"]["y"]
							var current_hp = a["currentHealth"]

							var pos = Vector2(position_x, position_y)
							var dir = Vector2(direction_x, direction_y)

							actor_manager.update_actor(id, pos, dir, current_hp)
						_:
							DebugLog.warn("unable to match actor state: %d" % state)

	DebugLog.debug("handle message end")

func set_client_id(id: String):
	DebugLog.info("setting client id to %s" % id)
	client_id = id