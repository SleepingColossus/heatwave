extends AudioStreamPlayer

onready var audio_steam_player : AudioStreamPlayer = get_node("/root/main/AudioStreamPlayer")

var sound_game_over = load("res://audio/sound/game_over.ogg")
var sound_pick_up = load("res://audio/sound/pick_up.ogg")
var sound_shoot = load("res://audio/sound/shoot.ogg")
var sound_take_damage = load("res://audio/sound/take_damage.ogg")
var sound_wave_clear = load("res://audio/sound/wave_clear.ogg")

func play_game_over():
	audio_steam_player.stream = sound_game_over
	audio_steam_player.play()

func play_pick_up():
	audio_steam_player.stream = sound_pick_up
	audio_steam_player.play()

func play_shoot():
	audio_steam_player.stream = sound_shoot
	audio_steam_player.play()

func play_take_damage():
	audio_steam_player.stream = sound_take_damage
	audio_steam_player.play()

func play_wave_clear():
	audio_steam_player.stream = sound_wave_clear
	audio_steam_player.play()
