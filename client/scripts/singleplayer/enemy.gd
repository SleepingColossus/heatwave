extends KinematicBody2D

onready var player = $"../../Player"
onready var sprite = $AnimatedSprite

export var speed : int = 50
export var attack_range : int = 200

func _ready():
	print_debug("Hi, I have spawned at: %d, %d" % [position.x, position.y])
	pass # Replace with function body.


func _process(delta):
	move()

func move() -> void:
	var distance_from_player = get_distance_between(position, player.position)

	if distance_from_player > attack_range:
		var direction = set_direction()
		var velocity : Vector2 = Vector2(direction.x * speed, direction.y * speed)

		move_and_slide(velocity)
		set_animation(direction)

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

func get_distance_between(from: Vector2, to: Vector2):
	var diff_x = abs(from.x - to.x)
	var diff_y = abs(from.y - to.y)

	var distance = sqrt(diff_x * diff_x + diff_y * diff_y)

	return distance

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
