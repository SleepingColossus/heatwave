package game

import "github.com/google/uuid"

type Enemy struct {
	Actor
	target             *Player // target player to chase
	attackRange        int     // minimal range to perform attack
	fireRate           int     // number of ticks between shots
	ticksSinceLastShot int
	projectileTemplate ProjectileTemplate
}

func newEnemy(t, maxHp, atkRange int, vel Vector2, fireRate int, pt ProjectileTemplate) *Enemy {
	return &Enemy{
		Actor: Actor{
			Id:            uuid.New().String(),
			State:         actorCreated,
			Type:          t,
			MaxHealth:     maxHp,
			CurrentHealth: maxHp,
			Body2D: Body2D{
				Position:  randomPosition(),
				Direction: zeroVector(),
				hitbox:    smallHitbox(),
				velocity:  vel,
			},
		},
		target:             nil,
		attackRange:        atkRange,
		fireRate:           fireRate,
		ticksSinceLastShot: 0,
		projectileTemplate: pt,
	}
}

func newEnemyFromTemplate(t EnemyTemplate, p ProjectileTemplate) *Enemy {
	return newEnemy(t.Type, t.MaxHP, t.Range, newVector2(t.Velocity, t.Velocity), t.FireRate, p)
}

func (e *Enemy) update() *Projectile {
	e.setChaseDirection()
	e.moveWithinRange()
	return e.shoot()
}

func (e *Enemy) moveWithinRange() {
	if e.target != nil {
		dist := e.Position.distanceTo(e.target.Position)

		if dist > e.attackRange {
			e.move()
		}
	}
}

func (e *Enemy) setChaseDirection() {
	if e.target != nil {
		var moveh, movev int

		// is target on right side?
		if e.Position.X < e.target.Position.X {
			moveh = moveRight
		} else if e.Position.X > e.target.Position.X {
			moveh = moveLeft
		} else {
			moveh = 0
		}

		// is target below?
		if e.Position.Y < e.target.Position.Y {
			movev = moveDown
		} else if e.Position.Y > e.target.Position.Y {
			movev = moveUp
		} else {
			movev = 0
		}

		e.setDirection(moveh, movev)
	}
}

func (e *Enemy) shoot() *Projectile {
	if e.target != nil && e.ticksSinceLastShot >= e.fireRate {
		dist := e.Position.distanceTo(e.target.Position)

		// am I within range?
		if dist <= e.attackRange {
			e.ticksSinceLastShot = 0
			return newProjectileFromTemplate(e.Actor, e.target.Position, e.projectileTemplate)
		}
	}

	e.ticksSinceLastShot++
	return nil
}

func (e *Enemy) setTarget(players []*Player) {
	if len(players) > 0 {
		e.target = players[0]
	}
}
