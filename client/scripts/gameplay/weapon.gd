class_name Weapon
extends Node2D

export var attack_rate: float
export var spread_amount: int
export var spread_angle: float
export var projectile_scene: PackedScene
export var shoot_sound: AudioStreamOGGVorbis

onready var reload_time: Timer = $ReloadTimer
onready var audio_stream: AudioStreamPlayer2D = $ShootSound

# angle offset when firing even number of shots
const even_offset = 5

var can_shoot: bool = true

func _ready():
	reload_time.wait_time = attack_rate

	audio_stream.stream = shoot_sound

func shoot(target_position: Vector2) -> void:
	spawn_projectile(target_position)
	audio_stream.play()
	can_shoot = false
	reload_time.start()

func spawn_projectile(target_position: Vector2) -> void:
	var spread_range = create_spread_array()

	var spread_angle_rad = deg2rad(spread_angle)
	var offset_rad = 0 if (spread_amount == 0 or spread_amount % 2 != 0) else deg2rad(5)

	for i in spread_range:
		var projectile = projectile_scene.instance()
		projectile.position = get_global_position()
		projectile.set_velocity(target_position, i * spread_angle_rad + offset_rad)

		# add child to game scene
		get_parent().get_parent().add_child(projectile)

func create_spread_array() -> Array:
	if spread_amount == 0: # single bullet
		return [0]
	elif spread_amount % 2 != 0: # multiple bullets (odd)
		var half: int = int(spread_amount / 2)
		return range(-half, half + 1)
	else: # multiple bullets (even)
		var half: int = int(spread_amount / 2)
		return range(-half, half)

func _on_ReloadTimer_timeout():
	can_shoot = true
