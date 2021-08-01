extends KinematicBody2D

class_name Actor

var direction: Vector2
var max_health: int
var current_health: int

onready var sprite = $AnimatedSprite

func _ready():
	direction = Vector2()

func _process(delta):
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

func set_animation(animation_name):
	if sprite.animation != animation_name:
		sprite.play(animation_name)

func set_direction(d: Vector2):
	direction = d

func set_current_health(hp):
	current_health = hp

func set_max_health(hp):
	max_health = hp

func delete():
	queue_free()
