extends Area2D

export var weapon_type: int
onready var collect_animation: AnimationPlayer = $CollectAnimation
onready var despawn_animation: AnimationPlayer = $DespawnAnimation
var is_collected: bool = false

func _on_collected(body):
	if !is_collected and body is Player:
		is_collected = true
		body.change_weapon(weapon_type)
		collect_animation.play("Collected")

func _on_DespawnTimer_timeout():
	despawn_animation.play("Despawned")
