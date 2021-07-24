package game

type Vector2 struct {
	X int
	Y int
}

func newVector2(x, y int) *Vector2 {
	return &Vector2{X:x, Y: y}
}
