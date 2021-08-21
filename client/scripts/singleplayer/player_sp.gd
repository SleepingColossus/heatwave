extends KinematicBody2D

export var max_health : int = 5
var current_health : int

export var speed : int = 100

onready var sprite = $AnimatedSprite

func _ready():
	current_health = max_health

func _process(delta):
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
