extends Node

# node all new actors will be appended to
var actor_container_node

# dictionary of all active actors
# k: actor id, v: actor
var actors : Dictionary = {}

# preload all actor resources
var player_resource = load("res://prefabs/player.tscn")

func create_actor(actor_id: String, actor_type: int, position: Vector2):
	print("creating actor type: %d with id: %s" % [actor_type, actor_id])
	
	var new_actor = create_resource_instance(actor_type)
	new_actor.position = position
	
	actor_container_node.add_child(new_actor)
	actors[actor_id] = new_actor
	
func delete_actor(actor_id: String):
	if actors.erase(actor_id):
		print("deleted actor " % actor_id)
	else:
		push_error("unknown actor id: %s" % actor_id)

func create_resource_instance(resource_type: int):
	match resource_type:
		ActorType.ActorType.PLAYER:
			return player_resource.instance()
		_:
			push_error("unknown resource type: %d" % resource_type)

func _ready():
	actor_container_node = get_node("/root/main/actor_container")
