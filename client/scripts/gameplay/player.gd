class_name Player
extends KinematicBody2D

enum WeaponType {
	PISTOL  = 0,
	UZI     = 1,
	SHOTGUN = 2,
	HARPOON = 3,
}

signal health_changed(current_hp, max_hp)
signal weapon_changed(new_weapon)
signal ammo_changed(amount)

export var max_health: int = 5
var current_health: int
var is_alive: bool = true

export var base_speed: int = 100

onready var sprite = $AnimatedSprite
onready var animation_player: AnimationPlayer = $AnimationPlayer
onready var hurt_sound = $HurtSound

# weapons
onready var pistol:  Weapon = $Pistol
onready var uzi:     Weapon = $Uzi
onready var shotgun: Weapon = $Shotgun
onready var harpoon: Weapon = $Harpoon
var current_weapon_type = WeaponType.PISTOL
var current_ammo = 0
onready var equip_weapon_timer: Timer = $EquipWeaponTimer
var ready_to_fire: bool
export var uzi_ammo: int
export var shotgun_ammo: int
export var harpoon_ammo: int

# dash
onready var dash_cooldown_timer: Timer = $DashCooldownTimer
onready var dash_duration_timer: Timer = $DashDurationTimer
onready var dash_sound: AudioStreamPlayer2D = $DashSound
onready var dash_particle: Particles2D = $DashParticle
var can_dash: bool
export var dash_speed_multiplier: int
var dash_value: int # current multiplier

func _ready():
	can_dash = true
	dash_value = 1
	dash_particle.emitting = false
	current_health = max_health
	$HealthBar.visible = false
	ready_to_fire = true

func _process(delta):
	if is_alive:
		poll_action()

		var direction: Vector2 = poll_movement()
		var speed = base_speed * dash_value
		var velocity : Vector2 = Vector2(direction.x * speed, direction.y * speed)
		velocity.normalized()

		move_and_slide(velocity)
		set_animation(direction)

		# adjust position to stay in bounds
		var window_size = OS.get_real_window_size()
		var sprite_size = 32
		position.x = clamp(position.x, 0, window_size.x - sprite_size)
		position.y = clamp(position.y, 0, window_size.y - sprite_size)

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
	if Input.is_action_pressed("shoot"):
		var current_weapon = get_current_weapon()
		if current_weapon.can_shoot and ready_to_fire:
			var mouse_position = get_viewport().get_mouse_position()
			current_weapon.shoot(mouse_position)

			if current_weapon_type != WeaponType.PISTOL:
				current_ammo -= 1
				emit_signal("ammo_changed", current_ammo as String)

			if current_ammo <= 0 and current_weapon_type != WeaponType.PISTOL:
				change_weapon(WeaponType.PISTOL)

	if Input.is_action_just_pressed("dash"):
		if can_dash:
			dash()

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

	if current_health <= 0:
		die()

	hurt_sound.play()
	animation_player.play("Hurt")
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
			current_ammo = uzi_ammo
		WeaponType.SHOTGUN:
			current_weapon_type = WeaponType.SHOTGUN
			current_ammo = shotgun_ammo
		WeaponType.HARPOON:
			current_weapon_type = WeaponType.HARPOON
			current_ammo = harpoon_ammo

	emit_signal("weapon_changed", w)

	if w == WeaponType.PISTOL:
		emit_signal("ammo_changed", "Inf")
	else:
		emit_signal("ammo_changed", current_ammo as String)

	# disable firing for short duriation when switching weapons
	ready_to_fire = false
	equip_weapon_timer.start()

func die() -> void:
	is_alive = false
	set_animation_by_name("Dying")

func dash() -> void:
	can_dash = false
	dash_value = dash_speed_multiplier
	dash_sound.play()
	dash_particle.emitting = true
	dash_duration_timer.start()

func _on_DashDurationTimer_timeout() -> void:
	dash_value = 1
	dash_particle.emitting = false
	dash_cooldown_timer.start()

func _on_DashCooldownTimer_timeout() -> void:
	can_dash = true

func _on_EquipWeaponTimer_timeout() -> void:
	ready_to_fire = true
