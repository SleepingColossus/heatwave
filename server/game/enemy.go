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
			Type:      t,
			Position:  randomPosition(),
			Direction: ZeroVector(),
			Hitbox:    smallHitbox(),
			Velocity:  1,
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
		} else {
			moveh = moveLeft
		}

		// is target below?
		if e.Position.Y < e.Target.Position.Y {
			movev = moveDown
		} else {
			movev = moveUp
		}

		e.SetDirection(moveh, movev)
	}
}

func newEnemyMeleeBasic() * Enemy {
	return newEnemy(enemyMeleeBasic)
}

func (e *Enemy) setTarget(players []*Player) {
	if len(players) > 0 {
		e.Target = players[0]
	}
}
