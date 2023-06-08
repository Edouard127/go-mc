package maths

import "golang.org/x/exp/constraints"

func EuclideanDistance2d[T constraints.Integer | constraints.Float](p1, p2 Vec2d[float32]) float64 {
	xDiff, yDiff := p1.X-p2.X, p1.Y-p2.Y
	return float64(xDiff*xDiff + yDiff*yDiff)
}

func EuclideanDistance3d(p1, p2 Vec3d[float64]) float64 {
	xDiff, yDiff, zDiff := p1.X-p2.X, p1.Y-p2.Y, p1.Z-p2.Z
	return xDiff*xDiff + yDiff*yDiff + zDiff*zDiff
}

func RotateLeft(x int64, k uint) int64 {
	return (x << k) | (x >> -k)
}
