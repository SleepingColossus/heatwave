extends Area2D

export var healing_amount : int = 1
onready var pickup_sound : AudioStreamPlayer2D = $PickupSound

func _on_HealthPickUp_body_entered(body):
	if body is Player:
		body.heal(healing_amount)
		pickup_sound.play()
		queue_free()
