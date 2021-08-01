extends KinematicBody2D


func _ready():
	pass

#func _process(delta):
#	pass

func delete():
	queue_free()
