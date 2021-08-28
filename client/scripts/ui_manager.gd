class_name UIManager
extends Node

onready var notification_container: GridContainer = $Notifications
onready var game_state_text: Label = $GameStateText

onready var sprite_player_hp_5: Sprite = $Health/PlayerHealth5HP
onready var sprite_player_hp_4: Sprite = $Health/PlayerHealth4HP
onready var sprite_player_hp_3: Sprite = $Health/PlayerHealth3HP
onready var sprite_player_hp_2: Sprite = $Health/PlayerHealth2HP
onready var sprite_player_hp_1: Sprite = $Health/PlayerHealth1HP
onready var sprite_player_hp_0: Sprite = $Health/PlayerHealthDead
onready var label_player_hp: Label = $Health/PlayerHealthText

var fade_rate = 0.005

func _ready():
	pass

func _process(delta):
	var labels = notification_container.get_children()

	for label in labels:
		# If still visible, fade away
		if label.modulate.a >= 0:
			label.modulate.a -= fade_rate
		# Delete element
		else:
			label.queue_free()

func add_notification(notification) -> void:
	var n : Label = Label.new()
	n.text = notification
	n.modulate = Color(0, 0, 0, 1)
	notification_container.add_child(n)

func show_game_state_update(text) -> void:
	game_state_text.visible = true
	game_state_text.text = text

func hide_game_state_text() -> void:
	game_state_text.visible = false

func set_player_health(amount: int) -> void:
	label_player_hp.text = amount as String

	hide_all_health_sprites()

	if amount >= 5:
		sprite_player_hp_5.visible = true
	elif amount == 4:
		sprite_player_hp_4.visible = true
	elif amount == 3:
		sprite_player_hp_3.visible = true
	elif amount == 2:
		sprite_player_hp_2.visible = true
	elif amount == 1:
		sprite_player_hp_1.visible = true
	else: # dead
		sprite_player_hp_0.visible = true

func hide_all_health_sprites() -> void:
	sprite_player_hp_5.visible = false
	sprite_player_hp_4.visible = false
	sprite_player_hp_3.visible = false
	sprite_player_hp_2.visible = false
	sprite_player_hp_1.visible = false
	sprite_player_hp_0.visible = false
