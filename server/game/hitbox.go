package game

type Hitbox = Vector2

func smallHitbox() *Hitbox {
	return NewVector2(32, 32)
}
