extends Control

onready var sprite: AnimatedSprite = $Menu/CenterRow/AnimationContainer/AnimatedSprite

func _ready():
	for button in $Menu/CenterRow/Buttons.get_children():
		button.connect("pressed", self, "_on_button_pressed", [button.scene_to_load])

func _on_button_pressed(scene_to_load):
	get_tree().change_scene(scene_to_load)

func _on_QuitButton_pressed():
	get_tree().quit()

func _on_PlayButton_mouse_entered():
	sprite.play("MoveLeft")

func _on_HelpButton_mouse_entered():
	sprite.play("IdleBottom")

func _on_QuitButton_mouse_entered():
	sprite.play("Dying")
