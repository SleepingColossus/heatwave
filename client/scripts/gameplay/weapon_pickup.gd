extends Area2D

export var weapon_type: int

func _on_collected(body):
	if body is Player:
		body.change_weapon(weapon_type)
		queue_free()
