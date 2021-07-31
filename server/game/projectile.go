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
			Id:    uuid.New().String(),
			Type:  projectilePlayerBullet,
			State: actorCreated,
			Body2D: Body2D{
				Position:  parent.Position,
				Direction: parent.Direction,
				hitbox:    zeroVector(),
				velocity:  10, // TODO get from resource file
			},
		},
	}
}

func newHostileProjectile(parent Body2D) *Projectile {
	return &Projectile{
		Actor: Actor{
			Id:    uuid.New().String(),
			Type:  projectileEnemyBullet,
			State: actorCreated,
			Body2D: Body2D{
				Position:  parent.Position,
				Direction: parent.Direction,
				hitbox:    zeroVector(),
				velocity:  5, // TODO get from resource file
			},
		},
	}
}

func (p *Projectile) update() {
	p.move()
	p.deleteIfOffScreen()
}

func (p *Projectile) deleteIfOffScreen() {
	if p.Position.X < 0 - offset ||
	p.Position.X > screenWidth + offset ||
	p.Position.Y < 0 - offset ||
	p.Position.Y > screenHeight + offset {
		p.State = actorDeleted
	}
}
