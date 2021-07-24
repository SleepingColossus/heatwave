package game

// actor types
const(
	Player int = iota
)

type Actor struct {
	Id string
	Position *Vector2
}

func NewActor(id string, pos *Vector2) *Actor {
	return &Actor{
		Id: id,
		Position: pos,
	}
}

func Move(position, direction *Vector2) Vector2 {
	return Vector2{
		X: position.X + direction.X,
		Y: position.Y + direction.Y,
	}
}
