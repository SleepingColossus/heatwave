class_name Waves

const wave_data : Dictionary = {
		1: {
			EnemyType.EnemyType.MELEE_BASIC: 1,
		},
		2: {
			EnemyType.EnemyType.MELEE_BASIC: 4,
			EnemyType.EnemyType.RANGED_BASIC: 2,

		},
		3: {
			EnemyType.EnemyType.MELEE_BASIC: 4,
			EnemyType.EnemyType.RANGED_BASIC: 4,
			EnemyType.EnemyType.TANK: 1,
		},
		4: {
			EnemyType.EnemyType.MELEE_BASIC: 6,
			EnemyType.EnemyType.MELEE_FAST: 1,
			EnemyType.EnemyType.RANGED_BASIC: 2,
			EnemyType.EnemyType.RANGED_ADVANCED: 1,
			EnemyType.EnemyType.TANK: 1,
		},
		5: {
			EnemyType.EnemyType.MELEE_BASIC: 10,
			EnemyType.EnemyType.MELEE_FAST: 3,
			EnemyType.EnemyType.RANGED_BASIC: 5,
			EnemyType.EnemyType.TANK: 2,
		},
		6: {
			EnemyType.EnemyType.MELEE_BASIC: 5,
			EnemyType.EnemyType.MELEE_FAST: 3,
			EnemyType.EnemyType.RANGED_BASIC: 10,
			EnemyType.EnemyType.RANGED_ADVANCED: 2,
			EnemyType.EnemyType.TANK: 2,
		},
		7: {
			EnemyType.EnemyType.MELEE_BASIC: 20,
			EnemyType.EnemyType.RANGED_BASIC: 7,
			EnemyType.EnemyType.RANGED_ADVANCED: 3,
			EnemyType.EnemyType.TANK: 1,
		},
		8: {
			EnemyType.EnemyType.RANGED_BASIC: 10,
			EnemyType.EnemyType.RANGED_ADVANCED: 5,
			EnemyType.EnemyType.TANK: 5,
		},
		9: {
			EnemyType.EnemyType.MELEE_BASIC: 15,
			EnemyType.EnemyType.RANGED_BASIC: 15,
			EnemyType.EnemyType.RANGED_ADVANCED: 5,
			EnemyType.EnemyType.TANK: 5,
		},
		10: {
			EnemyType.EnemyType.MELEE_FAST: 20,
		},
		11: {
			EnemyType.EnemyType.MELEE_FAST: 7,
			EnemyType.EnemyType.RANGED_BASIC: 20,
			EnemyType.EnemyType.RANGED_ADVANCED: 5,
		},
		12: {
			EnemyType.EnemyType.RANGED_ADVANCED: 10,
			EnemyType.EnemyType.TANK: 15,
		},
		13: {
			EnemyType.EnemyType.MELEE_FAST: 20,
			EnemyType.EnemyType.RANGED_ADVANCED: 5,
			EnemyType.EnemyType.TANK: 5,
			EnemyType.EnemyType.RANGED_BASIC: 20,
			EnemyType.EnemyType.MELEE_BASIC: 20,
		},
		14: {
			EnemyType.EnemyType.RANGED_BASIC: 40,
		},
		15: {
			EnemyType.EnemyType.MELEE_BASIC: 100,
		},
	}

static func get_wave(wave_number: int) -> Dictionary:
	return wave_data[wave_number]
