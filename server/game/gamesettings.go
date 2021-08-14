package game

type GameSettings struct {
	EnemySettings      map[int]EnemyTemplate
	ProjectileSettings map[int]ProjectileTemplate
	WaveSettings       []*Wave
}

func NewGameSettings(wavesPath, enemyPath, projectilePath string) *GameSettings {
	enemies := ReadEnemyData(enemyPath)
	waves := ReadWaveData(wavesPath, enemies)
	projectiles := ReadProjectileData(projectilePath)

	return &GameSettings{
		enemies,
		projectiles,
		waves,
	}
}
