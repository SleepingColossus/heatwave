class_name Waves

const wave_data : Dictionary = {
		1: {
			EnemyType.EnemyType.MELEE_BASIC: 1,
		},
		2: {
			EnemyType.EnemyType.MELEE_BASIC: 2,
			EnemyType.EnemyType.MELEE_FAST: 1,
		},
	}

static func get_wave(wave_number: int) -> Dictionary:
	return wave_data[wave_number]
