package random

import (
	"github.com/Edouard127/go-mc/maths"
)

type Xoroshiro128 Seed128

func NewXoroshiro128(seed Seed128) *Xoroshiro128 {
	return doXoroShiro128(seed[0], seed[1])
}

func doXoroShiro128(lower, higher int64) *Xoroshiro128 {
	seed := Seed128{lower, higher}
	if lower|higher == 0 {
		seed[0] = -7046029254386353131
		seed[1] = 7640891576956012809
	}
	return (*Xoroshiro128)(&seed)
}

func (x *Xoroshiro128) Next() int64 {
	s0, s1 := x[0], x[1]
	result := maths.RotateLeft(s0+s1, 17) + s0
	s1 ^= s0
	x[0] = maths.RotateLeft(s0, 49) ^ s1 ^ s1<<14
	x[1] = maths.RotateLeft(s1, 28)
	return result
}
