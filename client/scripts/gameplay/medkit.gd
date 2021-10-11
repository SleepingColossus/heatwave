extends Area2D

export var healing_amount : int = 1
onready var pickup_sound : AudioStreamPlayer2D = $PickupSound
onready var collect_animation : AnimationPlayer = $CollectAnimation
var is_collected: bool = false

func _on_HealthPickUp_body_entered(body):
	if !is_collected and body is Player:
		is_collected = true
		body.heal(healing_amount)
		collect_animation.play("Collected")
