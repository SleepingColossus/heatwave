package game

type Player struct {
	Actor
}

func newPlayer(id string) *Player {
	return &Player{
		Actor: Actor{
			Id:    id,
			Type:  player,
			State: actorCreated,
			CurrentHealth: 5,
			MaxHealth: 5,
			Body2D: Body2D{
				Position:  center(),
				Direction: zeroVector(),
				hitbox:    smallHitbox(),
				velocity:  newVector2(2, 2),
			},
		},
	}
}

// called on every tick
func (p *Player) update() {
	p.move()
}
