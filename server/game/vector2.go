package game

import (
	"math"
	"math/rand"
)

// screen edges
const (
	top int = iota
	bottom
	left
	right
)

// movement directions - horizontal
const (
	moveLeft  = -1
	moveRight = 1
)

// movement directions - vertical
const (
	moveUp = -1
	moveDown = 1
)

const (
	screenWidth  = 1920
	screenHeight = 1080
	offset = 50 // how far off screen to spawn and despawn actors
)

var (
	boundary = Vector2{screenWidth, screenHeight}
)

type Vector2 struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func newVector2(x, y int) Vector2 {
	return Vector2{x, y}
}

func zeroVector() Vector2 {
	return newVector2(0, 0)
}

func center() Vector2 {
	return newVector2(boundary.X / 2, boundary.Y / 2)
}

func (from *Vector2) distanceTo(to Vector2) int {
	distanceX := float64(absDiff(from.X, to.X))
	distanceY := float64(absDiff(from.Y, to.Y))

	distance := math.Sqrt(math.Pow(distanceX, 2) + math.Pow(distanceY, 2))

	return int(distance)
}

// get absolute difference between two ints
func absDiff(a, b int) int {
	diff := a - b

	if diff >= 0 {
		return diff
	} else {
		return -diff
	}
}

func randomPosition() Vector2 {
	edge := rand.Intn(4)

	if edge == top {
		return newVector2(rand.Intn(boundary.X), -offset)
	} else if edge == bottom {
		return newVector2(rand.Intn(boundary.X), boundary.Y+ offset)
	} else if edge == left {
		return newVector2(-offset, rand.Intn(boundary.Y))
	} else {
		return newVector2(boundary.X+ offset, rand.Intn(boundary.Y))
	}
}
