class_name Obstacle

extends StaticBody2D

export var max_health: int
var health: int
var is_alive: bool = true

onready var sprite: AnimatedSprite = $AnimatedSprite
onready var animation_player: AnimationPlayer = $AnimationPlayer
onready var debris_particles: Particles2D = $Debris
onready var collider: CollisionShape2D = $CollisionShape2D

func _ready():
	health = max_health
	set_sprite()

func take_damage(amount: int) -> void:
	health -= amount

	set_sprite()
	animation_player.play("Hurt")
	debris_particles.emitting = true

	if health <= 0:
		is_alive = false
		collider.set_deferred("disabled", true)

func set_sprite() -> void:
	if health == max_health:
		sprite.play("Normal")
	elif health >= (max_health / 2): # above 50%
		sprite.play("DamagedLow")
	elif health <= (max_health / 2) and health > 0: # below 50%
		sprite.play("DamagedHigh")
	else:
		sprite.play("Destroyed")
