package game

// TODO file contains a hardcoded list of enemies per wave
// TODO replace with an external resource such as XML or SQLite file
var (
	waveData = []*Wave{
		// 1
		newWave(toMap([]*Enemy{
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
		})),
		// 2
		newWave(toMap([]*Enemy{
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeFast(),
			newEnemyMeleeFast(),
		})),
		// 3
		newWave(toMap([]*Enemy{
			newEnemyRangedBasic(),
			newEnemyRangedBasic(),
			newEnemyRangedBasic(),
			newEnemyRangedAdvanced(),
			newEnemyRangedAdvanced(),
		})),
		// 4
		newWave(toMap([]*Enemy{
			newEnemyTank(),
			newEnemyTank(),
			newEnemyRangedAdvanced(),
			newEnemyRangedAdvanced(),
			newEnemyRangedAdvanced(),
			newEnemyRangedAdvanced(),
		})),
	}
)

func toMap(enemySlice []*Enemy) map[string]*Enemy {
	m := make(map[string]*Enemy)

	for _, enemy := range enemySlice {
		m[enemy.Id] = enemy
	}

	return m
}
