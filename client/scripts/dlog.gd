class_name DebugLog

# Toggles console logging on and off
# Turn off to prevent crashes when receiving large amounts of data :(
const debug_on = false

static func log(s) -> void:
	if debug_on:
		print_debug(s)
