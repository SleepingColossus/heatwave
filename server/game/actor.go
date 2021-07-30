package game

import (
	"strconv"
)

// actor types
const (
	player int = iota
	enemyMeleeBasic
	enemyMeleeFast
	enemyRangedBasic
	enemyRangedAdvanced
	enemyTank
	projectilePlayerBullet
	projectilePlayerHarpoon
	projectileEnemyBullet
	projectileEnemyHarpoon
)

type Updatable interface {
	update()
}

type Actor struct {
	Id        string
	Type      int
	Body2D
}

func (actor *Actor) ToMap() map[string]string {
	return map[string]string{
		"clientId":  actor.Id,
		"actorType": strconv.Itoa(actor.Type),
		"positionX": strconv.Itoa(actor.Position.X),
		"positionY": strconv.Itoa(actor.Position.Y),
		"directionX": strconv.Itoa(actor.Direction.X),
		"directionY": strconv.Itoa(actor.Direction.Y),
	}
}
