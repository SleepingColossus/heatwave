class_name DebugLog

enum LogLevel {
	DEBUG = 0
	INFO  = 1
	WARN  = 2
	ERROR = 3
}

# Filters out unwanted log messages
# Used to prevent crashes when logging large amounts of data :(
const min_log_level = LogLevel.INFO

static func _log(level: int, obj) -> void:
	if level >= min_log_level:
		if level == LogLevel.WARN:
			push_warning(obj)
		elif level == LogLevel.ERROR:
			push_error(obj)
		else:
			print_debug(obj)

static func debug(obj):
	_log(LogLevel.DEBUG, obj)

static func info(obj):
	_log(LogLevel.INFO, obj)

static func warn(obj):
	_log(LogLevel.WARN, obj)

static func error(obj):
	_log(LogLevel.ERROR, obj)
