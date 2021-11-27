package aoc

type Vector3 struct {
	X, Y, Z int
}

func NewVector3(x, y, z int) Vector3 {
	return Vector3{X: x, Y: y, Z: z}
}
