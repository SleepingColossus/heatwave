class_name Player
extends KinematicBody2D

signal shot_fired(from, to, type)

export var max_health: int = 5
var current_health: int
var is_alive: bool = true

export var speed: int = 100

export var weapon_type: int = 0
onready var reload_timer: Timer = $ReloadTimer

onready var sprite = $AnimatedSprite

func _ready():
	current_health = max_health
	$HealthBar.visible = false

func _process(delta):
	poll_action()

	var direction: Vector2 = poll_movement()
	var velocity : Vector2 = Vector2(direction.x * speed, direction.y * speed)
	velocity.normalized()

	move_and_slide(velocity)
	set_animation(direction)

func poll_movement() -> Vector2:
	var direction : Vector2 = Vector2()

	# horizontal
	if Input.is_action_pressed("move_left"):
		direction.x = -1
	elif Input.is_action_pressed("move_right"):
		direction.x = 1
	else:
		direction.x = 0

	# vertical
	if Input.is_action_pressed("move_up"):
		direction.y = -1
	elif Input.is_action_pressed("move_down"):
		direction.y = 1
	else:
		direction.y = 0

	return direction

func poll_action() -> void:
	if Input.is_action_just_pressed("shoot") and reload_timer.is_stopped():
		var mouse_position = get_viewport().get_mouse_position()
		emit_signal("shot_fired", position, mouse_position, weapon_type)
		$ShootSound.play()
		reload_timer.start()

func set_animation(direction: Vector2) -> void:
	if direction.x > 0:
		set_animation_by_name("MoveRight")
	elif direction.x < 0:
		set_animation_by_name("MoveLeft")
	elif direction.y > 0:
		set_animation_by_name("MoveBottom")
	elif direction.y < 0:
		set_animation_by_name("MoveTop")
	else:
		set_animation_by_name("IdleBottom")

func set_animation_by_name(animation_name) -> void:
	if sprite.animation != animation_name:
		sprite.play(animation_name)

func take_damage(amount: int) -> void:
	current_health -= amount

