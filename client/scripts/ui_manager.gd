extends Node

onready var notification_container: GridContainer = $Notifications

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

func add_notification(notification):
	var n : Label = Label.new()
	n.text = notification
	n.modulate = Color(0, 0, 0, 1)
	notification_container.add_child(n)
