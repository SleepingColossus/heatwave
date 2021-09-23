class_name Weapon
extends Node2D

export var attack_range: int
export var attack_rate: int
export var projectile_scene: PackedScene

onready var reload_time: Timer = $ReloadTimer
onready var line_of_sight: CollisionShape2D = $AttackRange/CollisionShape2D
onready var shoot_sound: AudioStreamPlayer2D = $ShootSound

var target: Player = null
var can_shoot: bool = true

func _ready():
	var visibility_radius = CircleShape2D.new()
	visibility_radius.radius = attack_range
	line_of_sight.shape = visibility_radius

	reload_time.wait_time = attack_rate

func shoot() -> void:
	spawn_projectile()
	shoot_sound.play()
	can_shoot = false
	reload_time.start()

func spawn_projectile() -> void:
	var projectile = projectile_scene.instance()
	projectile.position = get_global_position()
	projectile.set_velocity(target.global_position)

	# add child to game scene
	get_parent().get_parent().add_child(projectile)

func _on_AttackRange_body_entered(body):
	if body is Player:
		target = body

func _on_AttackRange_body_exited(body):
	if body is Player:
		target = null

func _on_ReloadTimer_timeout():
	can_shoot = true
