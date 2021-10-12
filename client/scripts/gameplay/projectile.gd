extends KinematicBody2D

# base speed
export var speed = 20

export var is_friendly: bool
export var damage: int
export var piercing: bool
export var durability: int

# calculated x, y velocity
var velocity : Vector2

const rotation_adjustment: float = 1.57 # 45 degrees in radians

onready var sprite: AnimatedSprite = $Sprite
onready var despawn_animation: AnimationPlayer = $DespawnAnimation

func _process(delta):
	move_and_slide(velocity)
	destroy_off_screen()

func set_velocity(to: Vector2, offset: float) -> void:
	var from = position

	var diff_x = to.x - from.x
	var diff_y = to.y - from.y

	var angle = atan2(diff_y, diff_x) + offset

	var velocity_x = cos(angle) * speed
	var velocity_y = sin(angle) * speed

	velocity = Vector2(velocity_x * speed, velocity_y * speed)

	self.rotation = angle + rotation_adjustment

func destroy_off_screen() -> void:
	var window_size = OS.get_real_window_size()

	if (position.x < 0
		or position.y < 0
		or position.x > window_size.x
		or position.y > window_size.y):
			queue_free()

func _on_Area2D_body_entered(body):
	if body.is_alive and ((body is Enemy and is_friendly) or (body is Player and not is_friendly)):
		body.take_damage(damage)

		if !piercing:
			queue_free()
		else:
			if body is Enemy:
				durability -= body.piercing_resistance

				if durability == 0:
					queue_free()

func _on_DespawnTimer_timeout():
	despawn_animation.play("Despawned")
