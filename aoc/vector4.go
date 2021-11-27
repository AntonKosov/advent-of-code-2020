package aoc

type Vector4 struct {
	X, Y, Z, W int
}

func NewVector4(x, y, z, w int) Vector4 {
	return Vector4{X: x, Y: y, Z: z, W: w}
}
