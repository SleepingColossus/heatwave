package game

const (
	defaultPlayerX int = 100
	defaultPlayerY int = 100
)

type Vector2 struct {
	X int
	Y int
}

func NewVector2(x, y int) *Vector2 {
	return &Vector2{x, y}
}

func ZeroVector() *Vector2 {
	return NewVector2(0, 0)
}

func DefaultPlayerPosition() *Vector2 {
	return NewVector2(defaultPlayerX, defaultPlayerY)
}
