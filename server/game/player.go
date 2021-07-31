package game

type Player struct {
	Actor
}

func newPlayer(id string) *Player {
	return &Player{
		Actor: Actor{
			Id:        id,
			Type:      player,
			State:     actorCreated,
			Body2D: Body2D{
				Position:  center(),
				Direction: ZeroVector(),
				Hitbox:    smallHitbox(),
				Velocity:  2,
			},
		},
	}
}

// called on every tick
func (p *Player) update() {
	p.move()
}
