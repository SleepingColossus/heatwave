package game

import "github.com/google/uuid"

// actor types
const(
	Player int = iota
)

type Actor struct {
	Id uuid.UUID
	Position Vector2
}

func newActor(pos Vector2) *Actor {
	return &Actor{
		Id: uuid.New(),
		Position: pos,
	}
}
