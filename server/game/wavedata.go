package game

// TODO file contains a hardcoded list of enemies per wave
// TODO replace with an external resource such as XML or SQLite file
var (
	waveData []*Wave = []*Wave{
		newWave([]*Enemy{
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
			newEnemyMeleeBasic(),
		}),
	}
)
