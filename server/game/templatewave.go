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

func (wt WavesTemplate) toGameWave(et map[int]EnemyTemplate) []*Wave {
	var waves []*Wave

	for _, waveTemplate := range wt.Waves {
		enemies := make([]*Enemy, 0)

		for _, enemyTemplate := range waveTemplate.Enemies {
			switch enemyTemplate.Type {
			case enemyMeleeBasic:
				enemies = append(enemies, newEnemyFromTemplate(et[enemyMeleeBasic]))
			case enemyMeleeFast:
				enemies = append(enemies, newEnemyFromTemplate(et[enemyMeleeFast]))
			case enemyRangedBasic:
				enemies = append(enemies, newEnemyFromTemplate(et[enemyRangedBasic]))
			case enemyRangedAdvanced:
				enemies = append(enemies, newEnemyFromTemplate(et[enemyRangedAdvanced]))
			case enemyTank:
				enemies = append(enemies, newEnemyFromTemplate(et[enemyTank]))
			}
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
