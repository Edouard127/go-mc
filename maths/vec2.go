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
	X, Y T // Pitch, Yaw
}

func (v *Vec2[T]) Set(v2 Vec2[T]) {
	v.X = v2.X
	v.Y = v2.Y
}

func (v *Vec2[T]) Add(v2 Vec2[T]) {
	v.X += v2.X
	v.Y += v2.Y
}

func (v *Vec2[T]) AddScalar(s T) {
	v.X += s
	v.Y += s
}

func (v *Vec2[T]) Sub(v2 Vec2[T]) {
	v.X -= v2.X
	v.Y -= v2.Y
}

func (v *Vec2[T]) SubScalar(s T) {
	v.X -= s
	v.Y -= s
}

func (v *Vec2[T]) Mul(v2 Vec2[T]) {
	v.X *= v2.X
	v.Y *= v2.Y
}

func (v *Vec2[T]) MulScalar(s T) {
	v.X *= s
	v.Y *= s
}

func (v *Vec2[T]) Div(v2 Vec2[T]) {
	v.X /= v2.X
	v.Y /= v2.Y
}

func (v *Vec2[T]) DivScalar(s T) {
	v.X /= s
	v.Y /= s
}

func (v *Vec2[T]) DistanceTo(v2 Vec2[T]) T {
	xDiff, yDiff := v.X-v2.X, v.Y-v2.Y
	return T(math.Sqrt(float64(xDiff*xDiff + yDiff*yDiff)))
}
