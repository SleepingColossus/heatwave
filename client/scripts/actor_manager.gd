extends Node

# node all new actors will be appended to
var actor_container_node

# dictionary of all active actors
# k: actor id, v: actor
var actors : Dictionary = {}

# preload all actor resources
var player_resource = load("res://prefabs/player.tscn")
var enemy_melee_basic = load("res://prefabs/enemy_melee_basic.tscn")

func create_actor(actor_id: String, actor_type: int, position: Vector2):
	DebugLog.debug("creating actor type: %d with id: %s" % [actor_type, actor_id])

	var new_actor = create_resource_instance(actor_type)
	new_actor.position = position

	actor_container_node.add_child(new_actor)
	actors[actor_id] = new_actor

func delete_actor(actor_id: String):
	if actors.erase(actor_id):
		DebugLog.debug("deleted actor " % actor_id)
	else:
		DebugLog.error("unknown actor id: %s" % actor_id)

func move_actor(actor_id: String, new_position: Vector2):
	DebugLog.debug("moving actor %s to new position: %d, %d" % [actor_id, new_position.x, new_position.y])

	if actors.has(actor_id):
		var actor = actors[actor_id]
		actor.position = new_position
	else:
		DebugLog.error("unknown actor: %s" % actor_id)

func create_resource_instance(resource_type: int):
	match resource_type:
		ActorType.ActorType.PLAYER:
			return player_resource.instance()
		ActorType.ActorType.ENEMY_MELEE_BASIC:
			return enemy_melee_basic.instance()
		_:
			DebugLog.error("unknown resource type: %d" % resource_type)

func _ready():
	actor_container_node = get_node("/root/main/actor_container")
