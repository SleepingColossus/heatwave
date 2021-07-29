extends KinematicBody2D

var direction: Vector2
onready var sprite = get_node("AnimatedSprite")

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
