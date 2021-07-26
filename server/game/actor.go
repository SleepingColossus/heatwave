package game

import "strconv"

// actor types
const (
	Player int = iota
)

type Actor struct {
	Id        string
	Position  *Vector2
	Direction *Vector2
	Velocity  int
}

func NewActor(id string, pos *Vector2) *Actor {
	return &Actor{
		Id:        id,
		Position:  pos,
		Direction: ZeroVector(),
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
