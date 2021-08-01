package game

import "github.com/google/uuid"

type Enemy struct {
	Actor
	target *Player // target player to chase
	attackRange int // minimal range to perform attack
}

func newEnemy(t, maxHp, vel, atkRange int) *Enemy {
	return &Enemy{
		Actor: Actor{
			Id:    uuid.New().String(),
			State: actorCreated,
			Type:  t,
			MaxHealth: maxHp,
			CurrentHealth: maxHp,
			Body2D: Body2D{
				Position:  randomPosition(),
				Direction: zeroVector(),
				hitbox:    smallHitbox(),
				velocity:  vel,
			},
		},
		target: nil,
		attackRange: atkRange,
	}
}

func (e *Enemy) update() {
	e.setChaseDirection()
	e.moveWithinRange()
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

func (e *Enemy) setTarget(players []*Player) {
	if len(players) > 0 {
		e.target = players[0]
	}
}
