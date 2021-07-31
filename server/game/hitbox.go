package game

type Hitbox = Vector2

func smallHitbox() Hitbox {
	return newVector2(32, 32)
}
