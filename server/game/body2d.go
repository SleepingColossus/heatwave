package game

type Body2D struct {
	Position  Vector2 `json:"position"`
	Direction Vector2 `json:"direction"`
	hitbox    Hitbox
	velocity  Vector2
}

func (body *Body2D) setDirection(newX, newY int) {
	body.Direction.X = newX
	body.Direction.Y = newY
}

func (body *Body2D) move() {
	// move in current direction
	body.Position.X += body.Direction.X * body.velocity.X
	body.Position.Y += body.Direction.Y * body.velocity.Y
}

func (b *Body2D) isColliding(point Vector2) bool {
	leftEdge, rightEdge, topEdge, bottomEdge := b.hitbox.getEdges(b.Position)

	if point.X > leftEdge && point.X < rightEdge && point.Y > topEdge && point.Y < bottomEdge {
		return true
	} else {
		return false
	}
}
