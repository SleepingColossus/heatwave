package game

import (
	"github.com/google/uuid"
	"log"
)

// projectile alignment
const (
	friendly int = iota
	hostile
)

type Projectile struct {
	Actor
	alignment int
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
				velocity:  newVector2(2, 2), // TODO get from resource file
			},
		},
		alignment: friendly,
	}
}

func newHostileProjectile(parent Body2D, target *Player) *Projectile {
	return &Projectile{
		Actor: Actor{
			Id:    uuid.New().String(),
			Type:  projectileEnemyBullet,
			State: actorCreated,
			Body2D: Body2D{
				Position:  parent.Position,
				Direction: parent.Direction,
				hitbox:    zeroVector(),
				velocity:  setVelocity(parent.Position, target.Position),
			},
		},
		alignment: hostile,
	}
}

func (p *Projectile) update(enemies map[string]*Enemy) {
	p.move()
	p.deleteIfOffScreen()
	p.checkCollision(enemies)
}

func (p *Projectile) deleteIfOffScreen() {
	if p.Position.X < 0 - offset ||
	p.Position.X > screenWidth + offset ||
	p.Position.Y < 0 - offset ||
	p.Position.Y > screenHeight + offset {
		p.State = actorDeleted
	}
}

func (p *Projectile) checkCollision(enemies map[string]*Enemy) {
	if p.State != actorDeleted && p.alignment == friendly {
		for _, e := range enemies {
			if e.State != actorDeleted && e.Body2D.isColliding(p.Position) {
				dmg := p.dmgAmount()
				e.takeDamage(dmg)

				// mark for deletion
				p.State = actorDeleted
			}
		}
	}
}

func (p *Projectile) dmgAmount() int {
	switch p.Type {
	case projectilePlayerBullet:
		return 1
	case projectileEnemyBullet:
		return 1
	case projectilePlayerHarpoon:
		return 3
	case projectileEnemyHarpoon:
		return 3
	default:
		log.Println("unknown projectile type", p.Type)
		return 0
	}
}

func setVelocity(projectile Vector2, target Vector2) Vector2 {
	diffX := projectile.X - target.X
	diffY := projectile.Y - target.Y

	velX := diffX / 20
	velY := diffY / 20

	return newVector2(velX, velY)
}
