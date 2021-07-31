package game

type Body2D struct {
	Position  Vector2
	Direction Vector2
	Hitbox    Hitbox
	Velocity  int
}

func (actor *Body2D) setDirection(newX, newY int) {
	actor.Direction.X = newX
	actor.Direction.Y = newY
}

func (actor *Body2D) move() {
	// move in current direction
	actor.Position.X += actor.Direction.X * actor.Velocity
	actor.Position.Y += actor.Direction.Y * actor.Velocity
}