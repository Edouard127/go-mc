package maths

import (
	"golang.org/x/exp/constraints"
)

func RotateLeft(x int64, k uint) int64 {
	return (x << k) | (x >> -k)
}

func ToSeed(x, y, z int32) int64 {
	i := x*3129871 ^ y*116129781 ^ z
	return int64((i*i*42317861 + i*11) >> 16)
}

func Lerp[T constraints.Integer | constraints.Float](x, y, z T) T {
	return x + (y-x)*z
}
