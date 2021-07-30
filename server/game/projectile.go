package game

import (
	"github.com/google/uuid"
)

type Projectile struct {
	Actor
}

func newFriendlyProjectile(parent Body2D) *Projectile {
	return &Projectile{
		Actor: Actor{
			Id:   uuid.New().String(),
			Type: projectilePlayerBullet,
			Body2D: Body2D{
				Position:  parent.Position,
				Direction: parent.Direction,
				Hitbox:    ZeroVector(),
				Velocity:  10, // TODO get from resource file
			},
		},
	}
}

func newHostileProjectile(parent Body2D) *Projectile {
	return &Projectile{
		Actor: Actor{
			Id:   uuid.New().String(),
			Type: projectileEnemyBullet,
			Body2D: Body2D{
				Position:  parent.Position,
				Direction: parent.Direction,
				Hitbox:    ZeroVector(),
				Velocity:  5, // TODO get from resource file
			},
		},
	}
}

func (p *Projectile) update() {
	p.move()
}
