extends Node

# node all new actors will be appended to
var actor_container_node

# dictionary of all active actors
# k: actor id, v: actor
var actors : Dictionary = {}

# preload all actor resources
var player_resource =       load("res://prefabs/player.tscn")
var enemy_melee_basic =     load("res://prefabs/enemies/enemy_melee_basic.tscn")
var enemy_melee_fast =      load("res://prefabs/enemies/enemy_melee_fast.tscn")
var enemy_ranged_basic =    load("res://prefabs/enemies/enemy_ranged_basic.tscn")
var enemy_ranged_advanced = load("res://prefabs/enemies/enemy_ranged_advanced.tscn")
var enemy_tank =            load("res://prefabs/enemies/enemy_tank.tscn")

var projectile_player_bullet = load("res://prefabs/projectiles/projectile_player_bullet.tscn")
var projectile_enemy_bullet =  load("res://prefabs/projectiles/projectile_enemy_bullet.tscn")

onready var sound_manager = get_node("/root/main/AudioStreamPlayer")
onready var ui_manager = get_node("/root/main/CanvasLayer/UI")

func create_actor(actor_id: String, actor_type: int, position: Vector2, max_hp: int, current_hp: int, is_self = false):
	DebugLog.debug("creating actor type: %d with id: %s" % [actor_type, actor_id])

	var new_actor = create_resource_instance(actor_type)
	new_actor.position = position

	actor_container_node.add_child(new_actor)
	actors[actor_id] = new_actor

	if new_actor.has_method("set_max_health"):
		new_actor.set_max_health(max_hp)

	if new_actor.has_method("set_current_health"):
		new_actor.set_current_health(current_hp)

	if actor_type == ActorType.ActorType.PLAYER:
		if is_self:
			new_actor.is_self = true
			new_actor.set_self_indicator_visible()
		else:
			new_actor.set_ally_indicator_visible()

	if (actor_type == ActorType.ActorType.PROJECTILE_PLAYER_BULLET
		|| actor_type == ActorType.ActorType.PROJECTILE_PLAYER_HARPOON
		|| actor_type == ActorType.ActorType.PROJECTILE_ENEMY_BULLET
		|| actor_type == ActorType.ActorType.PROJECTILE_ENEMY_HARPOON):

		sound_manager.play_shoot()

func delete_actor(actor_id: String):
	if actors.has(actor_id):
		# delete game object
		var actor_to_delete = actors[actor_id]
		actor_to_delete.delete()

		if actor_to_delete.has_method("set_current_health"):
			actor_to_delete.set_current_health(0)

		# delete from dictionary
		actors.erase(actor_id)

		DebugLog.debug("deleted actor %s" % actor_id)
	else:
		DebugLog.error("unknown actor id: %s" % actor_id)

func update_actor(actor_id: String, new_position: Vector2, new_direction: Vector2, current_hp: int):
	DebugLog.debug("moving actor %s to new position: %d, %d" % [actor_id, new_position.x, new_position.y])

	if actors.has(actor_id):
		var actor = actors[actor_id]
		actor.position = new_position

		if actor.has_method("set_current_health"):
			actor.set_current_health(current_hp)

		if actor.has_method("set_direction"):
			actor.set_direction(new_direction)

		if actor is Player:
			if actor.is_self:
				ui_manager.set_player_health(current_hp)

	else:
		DebugLog.error("unknown actor: %s" % actor_id)

func create_resource_instance(resource_type: int):
	match resource_type:
		ActorType.ActorType.PLAYER:
			return player_resource.instance()
		ActorType.ActorType.ENEMY_MELEE_BASIC:
			return enemy_melee_basic.instance()
		ActorType.ActorType.ENEMY_MELEE_FAST:
			return enemy_melee_fast.instance()
		ActorType.ActorType.ENEMY_RANGED_BASIC:
			return enemy_ranged_basic.instance()
		ActorType.ActorType.ENEMY_RANGED_ADVANCED:
			return enemy_ranged_advanced.instance()
		ActorType.ActorType.ENEMY_TANK:
			return enemy_tank.instance()
		ActorType.ActorType.PROJECTILE_PLAYER_BULLET:
			return projectile_player_bullet.instance()
		ActorType.ActorType.PROJECTILE_ENEMY_BULLET:
			return projectile_enemy_bullet.instance()
		_:
			DebugLog.error("unknown resource type: %d" % resource_type)

func _ready():
	actor_container_node = get_node("/root/main/actor_container")
