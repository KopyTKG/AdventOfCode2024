package vectors

type Vector2d struct {
	X, Y int
}

func New(x int, y int) *Vector2d {
	return &Vector2d{X: x, Y: y}
}

func (v *Vector2d) Copy() *Vector2d {
	return New(v.X, v.Y)
}

func (v *Vector2d) IsEqual(v2 *Vector2d) bool {
	return v.X == v2.X && v.Y == v2.Y
}
