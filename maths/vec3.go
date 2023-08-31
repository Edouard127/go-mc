package maths

import (
	"golang.org/x/exp/constraints"
	"math"
)

type Vec3b = Vec3[int8]
type Vec3s = Vec3[int16]
type Vec3i = Vec3[int]
type Vec3f = Vec3[float32]
type Vec3d = Vec3[float64]

type Vec3[T constraints.Integer | constraints.Float] struct {
	X, Y, Z T
}

func (v Vec3[T]) Add(vec3 Vec3[T]) Vec3[T] {
	return Vec3[T]{X: v.X + vec3.X, Y: v.Y + vec3.Y, Z: v.Z + vec3.Z}
}

func (v Vec3[T]) AddScalar(x, y, z T) Vec3[T] {
	return Vec3[T]{X: v.X + x, Y: v.Y + y, Z: v.Z + z}
}

func (v Vec3[T]) Sub(vec3 Vec3[T]) Vec3[T] {
	return Vec3[T]{X: v.X - vec3.X, Y: v.Y - vec3.Y, Z: v.Z - vec3.Z}
}

func (v Vec3[T]) SubScalar(x, y, z T) Vec3[T] {
	return Vec3[T]{X: v.X - x, Y: v.Y - y, Z: v.Z - z}
}

func (v Vec3[T]) Mul(vec3 Vec3[T]) Vec3[T] {
	return Vec3[T]{X: v.X * vec3.X, Y: v.Y * vec3.Y, Z: v.Z * vec3.Z}
}

func (v Vec3[T]) MulScalar(x, y, z T) Vec3[T] {
	return Vec3[T]{X: v.X * x, Y: v.Y * y, Z: v.Z * z}
}

func (v Vec3[T]) Div(vec3 Vec3[T]) Vec3[T] {
	return Vec3[T]{X: v.X / vec3.X, Y: v.Y / vec3.Y, Z: v.Z / vec3.Z}
}

func (v Vec3[T]) DivScalar(x, y, z T) Vec3[T] {
	return Vec3[T]{X: v.X / x, Y: v.Y / y, Z: v.Z / z}
}

func (v Vec3[T]) DistanceTo(vec3 Vec3[T]) T {
	xDiff, yDiff, zDiff := v.X-vec3.X, v.Y-vec3.Y, v.Z-vec3.Z
	return T(math.Sqrt(float64(xDiff*xDiff + yDiff*yDiff + zDiff*zDiff)))
}

func (v Vec3[T]) Scale(scale T) Vec3[T] {
	return v.MulScalar(scale, scale, scale)
}

func (v Vec3[T]) Reverse() Vec3[T] {
	return v.Scale(-1)
}

func (v Vec3[T]) Floor() Vec3[T] {
	return Vec3[T]{X: T(math.Floor(float64(v.X))), Y: T(math.Floor(float64(v.Y))), Z: T(math.Floor(float64(v.Z)))}
}

func (v Vec3[T]) Length() T {
	return T(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

func (v Vec3[T]) LengthSquared() T {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3[T]) HorizontalDistance() T {
	return T(math.Sqrt(float64(v.X*v.X + v.Z*v.Z)))
}

func (v Vec3[T]) HorizontalDistanceSquared() T {
	return v.X*v.X + v.Z*v.Z
}

func (v Vec3[T]) Lerp(vec3 Vec3[T], t T) Vec3[T] {
	return Vec3[T]{X: Lerp[T](v.X, vec3.X, t), Y: Lerp[T](v.Y, vec3.Y, t), Z: Lerp[T](v.Z, vec3.Z, t)}
}

func (v Vec3[T]) XRotation(rotation T) Vec3[T] {
	f := T(math.Cos(float64(rotation)))
	f1 := T(math.Sin(float64(rotation)))
	return Vec3[T]{X: v.X, Y: v.Y*f + v.Z*f1, Z: v.Z*f - v.Y*f1}
}

func (v Vec3[T]) YRotation(rotation T) Vec3[T] {
	f := T(math.Cos(float64(rotation)))
	f1 := T(math.Sin(float64(rotation)))
	return Vec3[T]{X: v.X*f + v.Z*f1, Y: v.Y, Z: v.Z*f - v.X*f1}
}

func (v Vec3[T]) ZRotation(rotation T) Vec3[T] {
	f := T(math.Cos(float64(rotation)))
	f1 := T(math.Sin(float64(rotation)))
	return Vec3[T]{X: v.X*f + v.Y*f1, Y: v.Y*f - v.X*f1, Z: v.Z}
}

func (v Vec3[T]) Spread() (T, T, T) {
	return v.X, v.Y, v.Z
}

func (v Vec3[T]) ToAABB() AxisAlignedBB[T] {
	return AxisAlignedBB[T]{MinX: v.X, MinY: v.Y, MinZ: v.Z, MaxX: v.X, MaxY: v.Y, MaxZ: v.Z}
}

func (v Vec3[T]) ToChunkPos() Vec2i {
	return Vec2i{int(v.X) >> 4, int(v.Z) >> 4}
}

func (v Vec3[T]) IsValid() bool {
	return !(math.IsNaN(float64(v.X)) || math.IsNaN(float64(v.Y)) || math.IsNaN(float64(v.Z))) &&
		!(math.IsInf(float64(v.X), 0) || math.IsInf(float64(v.Y), 0) || math.IsInf(float64(v.Z), 0))
}

func (v Vec3[T]) IsZero() bool {
	return v.X == 0 && v.Y == 0 && v.Z == 0 || !v.IsValid()
}
