extends Node

var bullet_resource = load("res://prefabs/projectiles/projectile_player_bullet.tscn")

var projectile_container

var weapon_type: int

func _ready():
	weapon_type = 0

	projectile_container = get_node("/root/PlaySinglePlayer/ProjectileContainer")

func shoot(from: Vector2, to: Vector2) -> void:
	var projectile = bullet_resource.instance()

	projectile.position = from
	projectile.set_velocity(to)

	projectile_container.add_child(projectile)
