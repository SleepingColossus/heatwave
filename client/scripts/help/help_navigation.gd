extends Control

export var main_menu_scene: PackedScene

func _process(delta):
	if Input.is_action_just_pressed("navigate_back"):
		get_tree().change_scene_to(main_menu_scene)

func _on_BackButton_pressed():
	get_tree().change_scene_to(main_menu_scene)
