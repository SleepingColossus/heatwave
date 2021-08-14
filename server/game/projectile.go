package game

import (
	"github.com/google/uuid"
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
	maxDistance float64
	distanceTravelled float64
	damage int
}

func newProjectileFromTemplate(parent Actor, targetPosition Vector2, template ProjectileTemplate) *Projectile {
	return &Projectile{
		Actor: Actor{
			Id:    uuid.New().String(),
			Type:  template.Type,
			State: actorCreated,
			Body2D: Body2D{
				Position:  parent.Position,
				velocity:  setVelocity(parent.Position, targetPosition, template.Velocity),
			},
		},
		alignment: getAlignment(template.Type),
		maxDistance: template.MaxDistance,
		distanceTravelled: 0,
		damage: template.Damage,
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
					e.takeDamage(p.damage)

					// mark for deletion
					p.State = actorDeleted
				}
			}
		} else { // hostile projectile
			for _, pl := range players {
				if pl.State != actorDeleted && pl.Body2D.isColliding(p.Position) {
					pl.takeDamage(p.damage)

					// mark for deletion
					p.State = actorDeleted
				}
			}
		}
	}
}

func setVelocity(projectile Vector2, target Vector2, baseVelocity float64) Vector2 {
	diffX := float64(target.X - projectile.X)
	diffY := float64(target.Y - projectile.Y)

	angle := math.Atan2(diffY, diffX)

	velX := int(math.Cos(angle) * baseVelocity)
	velY := int(math.Sin(angle) * baseVelocity)

	return newVector2(velX, velY)
}

func getAlignment(projectileType int) int {
	switch projectileType {
	case projectilePlayerBullet:
		return friendly
	case projectilePlayerHarpoon:
		return friendly
	default:
		return hostile
	}
}
