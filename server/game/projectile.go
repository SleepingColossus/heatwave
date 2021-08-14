package game

import (
	"github.com/google/uuid"
	"log"
	"math"
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

func newFriendlyProjectile(parent Body2D, targetPosition Vector2) *Projectile {
	return &Projectile{
		Actor: Actor{
			Id:    uuid.New().String(),
			Type:  projectilePlayerBullet,
			State: actorCreated,
			Body2D: Body2D{
				Position:  parent.Position,
				Direction: parent.Direction,
				hitbox:    zeroVector(),
				velocity:  setVelocity(parent.Position, targetPosition),
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
				Direction: zeroVector(),
				hitbox:    zeroVector(),
				velocity:  setVelocity(parent.Position, target.Position),
			},
		},
		alignment: hostile,
	}
}

func (p *Projectile) update(players map[string]*Player, enemies map[string]*Enemy) {
	p.move()
	p.deleteIfOffScreen()
	p.checkCollision(players, enemies)
}

func (p *Projectile) move() {
	p.Position.X += p.velocity.X
	p.Position.Y += p.velocity.Y
}

func (p *Projectile) deleteIfOffScreen() {
	if p.Position.X < 0 - offset ||
	p.Position.X > screenWidth + offset ||
	p.Position.Y < 0 - offset ||
	p.Position.Y > screenHeight + offset {
		p.State = actorDeleted
	}
}

func (p *Projectile) checkCollision(players map[string]*Player ,enemies map[string]*Enemy) {
	if p.State != actorDeleted {
		if p.alignment == friendly {
			for _, e := range enemies {
				if e.State != actorDeleted && e.Body2D.isColliding(p.Position) {
					dmg := p.dmgAmount()
					e.takeDamage(dmg)

					// mark for deletion
					p.State = actorDeleted
				}
			}
		} else { // hostile projectile
			for _, pl := range players {
				if pl.State != actorDeleted && pl.Body2D.isColliding(p.Position) {
					dmg := p.dmgAmount()
					pl.takeDamage(dmg)

					// mark for deletion
					p.State = actorDeleted
				}
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
	baseVelocity := 5.0

	diffX := float64(target.X - projectile.X)
	diffY := float64(target.Y - projectile.Y)

	angle := math.Atan2(diffY, diffX)

	velX := int(math.Cos(angle) * baseVelocity)
	velY := int(math.Sin(angle) * baseVelocity)

	return newVector2(velX, velY)
}
