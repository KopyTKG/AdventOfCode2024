package vector

type Vector struct {
	X, Y int
}

func New(x int, y int) *Vector {
	return &Vector{X: x, Y: y}
}

func (v *Vector) Copy() *Vector {
	return New(v.X, v.Y)
}

func (v *Vector) IsEqual(v2 *Vector) bool {
	return v.X == v2.X && v.Y == v2.Y
}
