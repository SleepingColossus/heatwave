package game

type Wave struct {
	State int
	Enemies []*Enemy
}

func newWave(enemies []*Enemy) *Wave {
	return &Wave{
		State: ready,
		Enemies: enemies,
	}
}

func (wave *Wave) Start(players []*Actor) {
	wave.State = started
	wave.setTargets(players)
}

func (wave *Wave) End() {
	wave.State = over
}

// TODO set target to closets player
func (wave *Wave) setTargets(players []*Actor) {
	for _, enemy := range wave.Enemies {
		enemy.setTarget(players)
	}
}
