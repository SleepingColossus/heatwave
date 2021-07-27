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

func (wave *Wave) Start() {
	wave.State = started
}

func (wave *Wave) End() {
	wave.State = over
}

// TODO set target to closets player
func (wave *Wave) SetTargets(players []*Actor) {
	for _, enemy := range wave.Enemies {
		enemy.setTarget(players)
	}
}
