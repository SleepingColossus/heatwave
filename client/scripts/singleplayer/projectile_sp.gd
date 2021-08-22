extends KinematicBody2D

# base speed
export var speed = 20

export var is_friendly: bool
export var damage: int

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


func _on_Area2D_body_entered(body):
	if body is Enemy and body.is_alive and is_friendly:
		body.take_damage(damage)
		queue_free()

	elif body is Player and not is_friendly:
		print_debug("hit player")
		queue_free()
