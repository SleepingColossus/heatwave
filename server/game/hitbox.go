package game

// represents size of hitbox on x & y scales
type Hitbox = Vector2

func smallHitbox() Hitbox {
	return newVector2(32, 32)
}

// gets left, right, top and bottom edge positions
// assumes origin is in center
func (h *Hitbox) getEdges(from Vector2) (int, int, int, int) {
	halfWidth := h.X / 2
	halfHeight := h.Y / 2

	leftEdge := from.X - halfWidth
	rightEdge := from.X + halfWidth
	topEdge := from.Y - halfHeight
	bottomEdge := from.Y + halfHeight

	return leftEdge, rightEdge, topEdge, bottomEdge
}
