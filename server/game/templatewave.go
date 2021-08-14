package game

type WavesTemplate struct {
	Waves []WaveTemplate `xml:"wave"`
}

type WaveTemplate struct {
	Enemies []EnemyWaveTemplate `xml:"enemy"`
}

type EnemyWaveTemplate struct {
	Type int `xml:"type,attr"`
}

func (wt WavesTemplate) toGameWave(et map[int]EnemyTemplate, pt map[int]ProjectileTemplate) []*Wave {
	var waves []*Wave

	for _, waveTemplate := range wt.Waves {
		enemies := make([]*Enemy, 0)

		for _, e := range waveTemplate.Enemies {
			var enemyTemplate EnemyTemplate

			switch e.Type {
			case enemyMeleeBasic:
				enemyTemplate = et[enemyMeleeBasic]
			case enemyMeleeFast:
				enemyTemplate = et[enemyMeleeFast]
			case enemyRangedBasic:
				enemyTemplate = et[enemyRangedBasic]
			case enemyRangedAdvanced:
				enemyTemplate = et[enemyRangedAdvanced]
			case enemyTank:
				enemyTemplate = et[enemyTank]
			}

			projectileTemplate := pt[enemyTemplate.ProjectileType]
			enemies = append(enemies, newEnemyFromTemplate(enemyTemplate, projectileTemplate))
		}

		wave := newWave(waveToMap(enemies))
		waves = append(waves, wave)
	}

	return waves
}

func waveToMap(enemySlice []*Enemy) map[string]*Enemy {
	m := make(map[string]*Enemy)

	for _, enemy := range enemySlice {
		m[enemy.Id] = enemy
	}

	return m
}
