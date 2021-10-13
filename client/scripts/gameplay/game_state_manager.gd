extends Node

enum GameState {
	PENDING,
	PLAYING,
	VICTORY,
	DEFEAT
}

# prefabs
export var enemy_melee_basic: PackedScene
export var enemy_melee_fast: PackedScene
export var enemy_ranged_basic: PackedScene
export var enemy_ranged_advanced: PackedScene
export var enemy_tank: PackedScene
export var medkit: PackedScene

# playable scenes to load
export var main_menu_scene: PackedScene

onready var player = $Player
onready var ui_manager = $CanvasLayer/UI
onready var wave_start_timer = $WaveStartTimer
onready var game_over_notification_timer = $GameOverNotificationTimer

var game_state: int
var current_wave: int
var number_of_waves: int
var enemy_count: int
const wave_timer_max_duration: int = 6
var wave_timer_remaining_duration: int = 6

func _ready():
	game_state = GameState.PENDING
	current_wave = 0
	number_of_waves = Waves.wave_data.size()

	player.connect("health_changed", self, "_on_player_health_changed")
	player.connect("weapon_changed", self, "_on_player_weapon_changed")
	player.connect("ammo_changed", self, "_on_player_ammo_changed")

func _process(delta):
	poll_buttons()

	if game_state == GameState.PENDING:

		if wave_start_timer.is_stopped():
			wave_timer_remaining_duration -= 1
			wave_start_timer.start()
			display_wave_countdown(wave_timer_remaining_duration)

		if wave_timer_remaining_duration == 0:
			current_wave += 1
			start_wave(current_wave)

	elif game_state == GameState.PLAYING:

		if not player.is_alive:
			lose()

		# all enemies cleared in wave?
		if enemy_count == 0:
			# not final wave?
			if current_wave != number_of_waves:
				wait_for_next_wave()
			# final wave
			else:
				win()

func poll_buttons():
	if Input.is_action_just_pressed("restart"):
		get_tree().reload_current_scene()

	if Input.is_action_just_pressed("navigate_back"):
		get_tree().change_scene_to(main_menu_scene)

func _on_player_health_changed(current_hp: int, max_hp: int) -> void:
	ui_manager.set_player_health(current_hp)

func _on_player_weapon_changed(new_weapon: int) -> void:
	ui_manager.change_player_weapon(new_weapon)

func _on_player_ammo_changed(amount: String) -> void:
	ui_manager.set_player_ammo(amount)

func _on_enemy_died() -> void:
	enemy_count -= 1
	print_debug("Enemies remaining: %d" % enemy_count)

func start_wave(wave_number: int) -> void:
	game_state = GameState.PLAYING

	var wave_data : Dictionary = Waves.get_wave(wave_number)

	spawn_instance_batch(wave_data, EnemyType.EnemyType.MELEE_BASIC,     enemy_melee_basic)
	spawn_instance_batch(wave_data, EnemyType.EnemyType.MELEE_FAST,      enemy_melee_fast)
	spawn_instance_batch(wave_data, EnemyType.EnemyType.RANGED_BASIC,    enemy_ranged_basic)
	spawn_instance_batch(wave_data, EnemyType.EnemyType.RANGED_ADVANCED, enemy_ranged_advanced)
	spawn_instance_batch(wave_data, EnemyType.EnemyType.TANK,            enemy_tank)

	if wave_number % 5 == 0:
		spawn_medkit()

	ui_manager.update_wave_number(wave_number)
	ui_manager.hide_game_state_text()

func spawn_instance_batch(wave_dict: Dictionary, key: int, resource) -> void:
	if wave_dict.has(key):
		var amount = wave_dict[key]

		for n in range(amount):
			spawn_instance(resource)

func spawn_instance(resource) -> void:
	var instance = resource.instance()
	instance.position = generate_random_position()
	instance.connect("died", self, "_on_enemy_died")

	add_child(instance)
	enemy_count += 1

func spawn_medkit() -> void:
	var instance = medkit.instance()
	instance.position = generate_center_position()
	add_child(instance)

func generate_center_position() -> Vector2:
	var window_size = OS.get_real_window_size()
	return Vector2(window_size.x / 2, window_size.y / 2)

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

func win() -> void:
	game_state = GameState.VICTORY
	ui_manager.show_game_state_update("You win!")
	game_over_notification_timer.start()

func lose() -> void:
	game_state = GameState.DEFEAT
	ui_manager.show_game_state_update("You lose!")
	game_over_notification_timer.start()

func display_wave_countdown(seconds_remaining) -> void:
	ui_manager.show_game_state_update(seconds_remaining as String)

func wait_for_next_wave() -> void:
	game_state = GameState.PENDING
	wave_timer_remaining_duration = wave_timer_max_duration

# show instructions after game over
func _on_GameOverNotificationTimer_timeout():
	ui_manager.show_game_state_update("Press R to restart  |  Press ESC to exit")
