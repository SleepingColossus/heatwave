package game

// TODO file contains a hardcoded list of enemies per wave
// TODO replace with an external resource such as XML or SQLite file
var (
	waveData = []*Wave{
		newWave(toMap([]*Enemy{
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
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