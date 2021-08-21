extends KinematicBody2D

export var max_health : int = 5
var current_health : int

export var speed : int = 100
var direction : Vector2

onready var sprite = $AnimatedSprite

func _ready():
	current_health = max_health
	direction = Vector2()

func _process(delta):
	poll_movement()

	var velocity : Vector2 = Vector2(direction.x * speed, direction.y * speed)

	move_and_slide(velocity)

	if direction.x > 0:
		set_animation("MoveRight")
	elif direction.x < 0:
		set_animation("MoveLeft")
	elif direction.y > 0:
		set_animation("MoveBottom")
	elif direction.y < 0:
		set_animation("MoveTop")
	else:
		set_animation("IdleBottom")

func poll_movement() -> void:
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

func set_animation(animation_name):
	if sprite.animation != animation_name:
		sprite.play(animation_name)
