package game

import (
	"strconv"
)

// actor types
const (
	Player int = iota
	EnemyMeleeBasic
)

type Actor struct {
	Id        string
	Type      int
	Position  *Vector2
	Direction *Vector2
	Hitbox    *Hitbox
	Velocity  int
}


func NewPlayer(id string) *Actor {
	return &Actor{
		Id:        id,
		Type:      Player,
		Position:  center(),
		Direction: ZeroVector(),
		Hitbox:    smallHitbox(),
		Velocity:  1,
	}
}

func (actor *Actor) SetDirection(newDirection Vector2) {
	actor.Direction.X = newDirection.X
	actor.Direction.Y = newDirection.Y
}

func (actor *Actor) Move() {
	actor.Position.X += actor.Direction.X * actor.Velocity
	actor.Position.Y += actor.Direction.Y * actor.Velocity
}

func (actor *Actor) ToMap() map[string]string {
	return map[string]string{
		"clientId":  actor.Id,
		"positionX": strconv.Itoa(actor.Position.X),
		"positionY": strconv.Itoa(actor.Position.Y),
	}
}
