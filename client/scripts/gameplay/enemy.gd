class_name Enemy
extends KinematicBody2D

signal died

onready var player = $"../Player"
onready var sprite: AnimatedSprite = $AnimatedSprite
onready var health_bar: ProgressBar = $HealthBar
onready var weapon : Weapon = $Weapon
onready var line_of_sight: Area2D = $LineOfSight

export var max_health: int = 5
var current_health: int

export var speed: int = 50
export var self_destruct : bool = false

var is_alive: bool = true
var target: Player = null

func _ready():
	current_health = max_health

	health_bar.max_value = max_health
	health_bar.value = current_health
	health_bar.modulate = Color(0, 1, 0, 1)

func _process(delta):
	if is_alive:
		var distance_from_player = get_distance_between(position, player.position)

		if target == null: # player not in shooting range
			var direction = set_direction()
			var velocity = get_velocity_towards(player.position)

			move_and_slide(velocity)
			set_animation(direction)
			sprite.play()
		else:
			sprite.stop()

			if weapon.can_shoot:
				weapon.shoot(target.global_position)

				if self_destruct:
					die()
					sprite.play()

func set_direction() -> Vector2:
	var chase_direction = Vector2()

	if position.x > player.position.x:
		chase_direction.x = -1
	if position.x < player.position.x:
		chase_direction.x = 1
	if position.y > player.position.y:
		chase_direction.y = -1
	if position.y < player.position.y:
		chase_direction.y = 1

	return chase_direction

func get_velocity_towards(to: Vector2) -> Vector2:
	var from = position

	var diff_x = to.x - from.x
	var diff_y = to.y - from.y

	var angle = atan2(diff_y, diff_x)

	var velocity_x = cos(angle) * speed
	var velocity_y = sin(angle) * speed

	return Vector2(velocity_x, velocity_y)

func get_distance_between(from: Vector2, to: Vector2) -> float:
	return from.distance_to(to)

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
	health_bar.value = current_health

	# full health
	if current_health == max_health:
		health_bar.modulate = Color(0, 1, 0, 1)
	# injured
	elif current_health < max_health && current_health > (max_health / 2):
		health_bar.modulate = Color(1, 1, 0, 1)
	# critical
	else:
		health_bar.modulate = Color(1, 0, 0, 1)

	if current_health <= 0:
		die()

func die() -> void:
	is_alive = false
	set_animation_by_name("Dying")
	health_bar.visible = false
	emit_signal("died")


func _on_LineOfSight_body_entered(body) -> void:
	if body is Player:
		target = body

func _on_LineOfSight_body_exited(body) -> void:
	if body is Player:
		target = null
