package game

type WavesTemplate struct {
	Waves []WaveTemplate `xml:"wave"`
}

type WaveTemplate struct {
	Enemies []EnemyTemplate `xml:"enemy"`
}

type EnemyTemplate struct {
	Type int `xml:"type,attr"`
}

func (wt WavesTemplate) toGameWave() []*Wave {
	var waves []*Wave

	for _, waveTemplate := range wt.Waves {
		enemies := make([]*Enemy, 0)

		for _, enemyTemplate := range waveTemplate.Enemies {
			switch enemyTemplate.Type {
			case enemyMeleeBasic:
				enemies = append(enemies, newEnemyMeleeBasic())
			case enemyMeleeFast:
				enemies = append(enemies, newEnemyMeleeFast())
			case enemyRangedBasic:
				enemies = append(enemies, newEnemyRangedBasic())
			case enemyRangedAdvanced:
				enemies = append(enemies, newEnemyRangedAdvanced())
			case enemyTank:
				enemies = append(enemies, newEnemyTank())
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
