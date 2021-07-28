package game

import (
	"strconv"
)

// actor types
const (
	player int = iota
	enemyMeleeBasic
)

type Updatable interface {
	update()
}

type Actor struct {
	Id        string
	Type      int
	Position  *Vector2
	Direction *Vector2
	Hitbox    *Hitbox
	Velocity  int
}

func (actor *Actor) SetDirection(newX, newY int) {
	actor.Direction.X = newX
	actor.Direction.Y = newY
}

func (actor *Actor) move() {
	// move in current direction
	actor.Position.X += actor.Direction.X * actor.Velocity
	actor.Position.Y += actor.Direction.Y * actor.Velocity
}

func (actor *Actor) toMap() map[string]string {
	return map[string]string{
		"clientId":  actor.Id,
		"actorType": strconv.Itoa(actor.Type),
		"positionX": strconv.Itoa(actor.Position.X),
		"positionY": strconv.Itoa(actor.Position.Y),
	}
}
