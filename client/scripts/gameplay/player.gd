class_name Player
extends KinematicBody2D

enum WeaponType {
	PISTOL  = 0,
	UZI     = 1,
	SHOTGUN = 2,
	HARPOON = 3,
}

signal health_changed(current_hp, max_hp)

export var max_health: int = 5
var current_health: int
var is_alive: bool = true

export var speed: int = 100

export var weapon_type: int = 0
onready var reload_timer: Timer = $ReloadTimer

onready var sprite = $AnimatedSprite

onready var pistol:  Weapon = $Pistol
onready var uzi:     Weapon = $Uzi
onready var shotgun: Weapon = $Shotgun
onready var harpoon: Weapon = $Harpoon
var current_weapon_type = WeaponType.PISTOL
var current_ammo = 0

func _ready():
	current_health = max_health
	$HealthBar.visible = false

func _process(delta):
	poll_action()

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

func poll_action() -> void:
	if Input.is_action_just_pressed("shoot"):
		var current_weapon = get_current_weapon()
		if current_weapon.can_shoot:
			var mouse_position = get_viewport().get_mouse_position()
			current_weapon.shoot(mouse_position)

			if current_weapon_type != WeaponType.PISTOL:
				current_ammo -= 1

			if current_ammo <= 0 and current_weapon_type != WeaponType.PISTOL:
				change_weapon(WeaponType.PISTOL)

func get_current_weapon() -> Weapon:
	match current_weapon_type:
		WeaponType.PISTOL:
			return pistol
		WeaponType.UZI:
			return uzi
		WeaponType.SHOTGUN:
			return shotgun
		WeaponType.HARPOON:
			return harpoon

	# default case
	return null

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

	if current_health < 0:
		current_health = 0

	emit_signal("health_changed", current_health, max_health)

func heal(amount: int) -> void:
	current_health += amount

	if current_health > max_health:
		current_health = max_health

	emit_signal("health_changed", current_health, max_health)

func change_weapon(w) -> void:
	match w:
		WeaponType.PISTOL:
			current_weapon_type = WeaponType.PISTOL
			current_ammo = 0
		WeaponType.UZI:
			current_weapon_type = WeaponType.UZI
			current_ammo = 20
		WeaponType.SHOTGUN:
			current_weapon_type = WeaponType.SHOTGUN
			current_ammo = 5
		WeaponType.HARPOON:
			current_weapon_type = WeaponType.HARPOON
			current_ammo = 3

func is_alive() -> bool:
	return current_health > 0
