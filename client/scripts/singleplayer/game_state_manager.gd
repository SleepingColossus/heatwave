extends Node

# preload resources prior to spawning -----------------------------------------------
var enemy_melee_basic =     load("res://prefabs/enemies/enemy_melee_basic.tscn")
var enemy_melee_fast =      load("res://prefabs/enemies/enemy_melee_fast.tscn")
var enemy_ranged_basic =    load("res://prefabs/enemies/enemy_ranged_basic.tscn")
var enemy_ranged_advanced = load("res://prefabs/enemies/enemy_ranged_advanced.tscn")
var enemy_tank =            load("res://prefabs/enemies/enemy_tank.tscn")

var bullet_resource = load("res://prefabs/projectiles/projectile_player_bullet.tscn")
# -----------------------------------------------------------------------------------

onready var player = $"../Player"
onready var sound_manager = $"../AudioStreamPlayer2D"

var current_wave: int
var number_of_waves: int

func _ready():
	current_wave = 1
	number_of_waves = Waves.wave_data.size()

	player.connect("shot_fired", self, "_on_shot_fired")

	start_wave(current_wave)

func _process(delta):
	pass

func _on_shot_fired(from: Vector2,to: Vector2, weapon_type: int) -> void:
	var projectile = bullet_resource.instance()

	projectile.position = from
	projectile.set_velocity(to)

	add_child(projectile)

	sound_manager.play_shoot()

func start_wave(wave_number: int) -> void:
	var wave_data : Dictionary = Waves.get_wave(wave_number)

	if wave_data.has(EnemyType.EnemyType.MELEE_BASIC):
		var amount = wave_data[EnemyType.EnemyType.MELEE_BASIC]

		for n in range(amount + 1):
			spawn_instance(enemy_melee_basic)

func spawn_instance(resource) -> void:
	var instance = resource.instance()
	instance.position = generate_random_position()
	add_child(instance)

func generate_random_position() -> Vector2:
	var window_size = OS.get_real_window_size()
	var screen_offset = 40
	var min_x = -screen_offset
	var min_y = -screen_offset
	var max_x = screen_offset + window_size.x
	var max_y = screen_offset + window_size.y

	var rng = RandomNumberGenerator.new()
	rng.randomize()
	var screen_edge = rng.randi_range(ScreenEdge.ScreenEdge.TOP, ScreenEdge.ScreenEdge.RIGHT)
	var spawn_position = Vector2()

	if screen_edge == ScreenEdge.ScreenEdge.TOP :
		spawn_position.y = min_y
		spawn_position.x = rng.randi_range(min_x, max_x)
	elif screen_edge == ScreenEdge.ScreenEdge.BOTTOM :
		spawn_position.y = max_y
		spawn_position.x = rng.randi_range(min_x, max_x)
	elif screen_edge == ScreenEdge.ScreenEdge.LEFT :
		spawn_position.x = min_x
		spawn_position.y = rng.randi_range(min_y, max_y)
	else :
		spawn_position.x = max_x
		spawn_position.y = rng.randi_range(min_y, max_y)

	return spawn_position
