extends Node

onready var player = $"../Player"
onready var sound_manager = $"../AudioStreamPlayer2D"

var bullet_resource = load("res://prefabs/projectiles/projectile_player_bullet.tscn")

func _ready():

	player.connect("shot_fired", self, "_on_shot_fired")

func _process(delta):
	pass

func _on_shot_fired(from: Vector2,to: Vector2, weapon_type: int) -> void:
	var projectile = bullet_resource.instance()

	projectile.position = from
	projectile.set_velocity(to)

	add_child(projectile)

	sound_manager.play_shoot()
