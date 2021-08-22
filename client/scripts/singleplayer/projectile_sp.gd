extends KinematicBody2D

# base speed
export var speed = 20

# calculated x, y velocity
var velocity : Vector2

func _process(delta):
	move_and_slide(velocity)

func set_velocity(to: Vector2) -> void:
	var from = position

	var diff_x = to.x - from.x
	var diff_y = to.y - from.y

	var angle = atan2(diff_y, diff_x)

	var velocity_x = cos(angle) * speed
	var velocity_y = sin(angle) * speed

	velocity = Vector2(velocity_x * speed, velocity_y * speed)
