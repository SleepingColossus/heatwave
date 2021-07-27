package game

import "github.com/google/uuid"

type Enemy struct {
	Actor
	Target *Actor // target player to chase
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

func newEnemyMeleeBasic() * Enemy {
	return newEnemy(EnemyMeleeBasic)
}

func (e *Enemy) setTarget(players []*Actor) {
	if len(players) > 0 {
		e.Target = players[0]
	}
}
