extends Actor

class_name PlayerMP

var is_self: bool
onready var self_indicator = $Indicator/ArrowIndicatorPlayer
onready var ally_indicator = $Indicator/ArrowIndicatorAlly

func set_self_indicator_visible():
	self_indicator.visible = true
	health_bar.visible = false

func set_ally_indicator_visible():
	ally_indicator.visible = true
