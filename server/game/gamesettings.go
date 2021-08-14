package game

type GameSettings struct {
	EnemySettings map[int]EnemyTemplate
	WaveSettings []*Wave
}

func NewGameSettings(wavesPath, enemyPath string) *GameSettings {
	enemies := ReadEnemyData(enemyPath)
	waves := ReadWaveData(wavesPath, enemies)

	return &GameSettings{
		enemies,
		waves,
	}
}
