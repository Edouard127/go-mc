package maths

import (
	"golang.org/x/exp/constraints"
	"math"
)

type Vec2b = Vec2[int8]
type Vec2s = Vec2[int16]
type Vec2i = Vec2[int]
type Vec2f = Vec2[float32]
type Vec2d = Vec2[float64]

type Vec2[T constraints.Integer | constraints.Float] struct {
	X, Y T
	// Pitch, Yaw
}

func ParseVec2[T constraints.Float](x, y T) Vec2[T] {
	return Vec2[T]{X: x, Y: y}
}

func (v Vec2[T]) Add(v2 Vec2[T]) Vec2[T] {
	return Vec2[T]{X: v.X + v2.X, Y: v.Y + v2.Y}
}

func (v Vec2[T]) AddScalar(s T) Vec2[T] {
	return Vec2[T]{X: v.X + s, Y: v.Y + s}
}

func (v Vec2[T]) Sub(v2 Vec2[T]) Vec2[T] {
	return Vec2[T]{X: v.X - v2.X, Y: v.Y - v2.Y}
}

func (v Vec2[T]) SubScalar(s T) Vec2[T] {
	return Vec2[T]{X: v.X - s, Y: v.Y - s}
}

func (v Vec2[T]) Mul(v2 Vec2[T]) Vec2[T] {
	return Vec2[T]{X: v.X * v2.X, Y: v.Y * v2.Y}
}

func (v Vec2[T]) MulScalar(s T) Vec2[T] {
	return Vec2[T]{X: v.X * s, Y: v.Y * s}
}

func (v Vec2[T]) Div(v2 Vec2[T]) Vec2[T] {
	return Vec2[T]{X: v.X / v2.X, Y: v.Y / v2.Y}
}

func (v Vec2[T]) DivScalar(s T) Vec2[T] {
	return Vec2[T]{X: v.X / s, Y: v.Y / s}
}

func (v Vec2[T]) DistanceTo(v2 Vec2[T]) T {
	xDiff, yDiff := v.X-v2.X, v.Y-v2.Y
	return T(math.Sqrt(float64(xDiff*xDiff + yDiff*yDiff)))
}
