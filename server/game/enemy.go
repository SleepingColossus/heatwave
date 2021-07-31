package game

import "github.com/google/uuid"

type Enemy struct {
	Actor
	Target *Player // target player to chase
}

func newEnemy(t int) *Enemy {
	return &Enemy{
		Actor: Actor{
			Id:        uuid.New().String(),
			State:     actorCreated,
			Type:      t,
			Body2D: Body2D{
				Position:  randomPosition(),
				Direction: ZeroVector(),
				Hitbox:    smallHitbox(),
				Velocity:  1,
			},
		},
		Target: nil,
	}
}

func (e *Enemy) update() {
	e.setChaseDirection()
	e.move()
}

func (e *Enemy) setChaseDirection() {
	if e.Target != nil {
		var moveh, movev int

		// is target on right side?
		if e.Position.X < e.Target.Position.X {
			moveh = moveRight
		} else if e.Position.X > e.Target.Position.X {
			moveh = moveLeft
		} else {
			moveh = 0
		}

		// is target below?
		if e.Position.Y < e.Target.Position.Y {
			movev = moveDown
		} else if e.Position.Y > e.Target.Position.Y {
			movev = moveUp
		} else {
			movev = 0
		}

		e.setDirection(moveh, movev)
	}
}

func newEnemyMeleeBasic() *Enemy {
	return newEnemy(enemyMeleeBasic)
}

func (e *Enemy) setTarget(players []*Player) {
	if len(players) > 0 {
		e.Target = players[0]
	}
}
