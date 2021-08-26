extends Node

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

func _process(delta):
	if can_shoot and target != null:
		shoot()

func shoot():
	shoot_sound.play()
	can_shoot = false
	reload_time.start()

func _on_AttackRange_body_entered(body):
	if body is Player:
		target = body

func _on_AttackRange_body_exited(body):
	if body is Player:
		target = null

func _on_ReloadTimer_timeout():
	can_shoot = true
