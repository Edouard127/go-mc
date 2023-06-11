package random

import "github.com/Edouard127/go-mc/maths"

type XoroshiroSource struct {
	RandomSource
	generator *Xoroshiro128
	gaussian  *MarsagliaPolarGaussian
}

func NewXoroshiroRandomSource(seed int64) *XoroshiroSource {
	src := &XoroshiroSource{generator: NewXoroshiro128(UpgradeSeed(seed))}
	src.gaussian = NewMarsagliaPolarGaussian(src)
	return src
}

func NewXoroshiroRandomSourceAt(x, y, z int32, seed int64) *XoroshiroSource {
	s := UpgradeSeed(seed)
	src := &XoroshiroSource{generator: doXoroShiro128(maths.ToSeed(x, y, z)^s[0], s[1])}
	src.gaussian = NewMarsagliaPolarGaussian(src)
	return src
}

func (x *XoroshiroSource) SetSeed(seed int64) {
	x.generator = NewXoroshiro128(UpgradeSeed(seed))
	x.gaussian.Reset()
}

func (x *XoroshiroSource) NextInt(bits int) int {
	if bits < 0 {
		panic("bits must be positive")
	}

	i := x.generator.Next() & 0xFFFFFFFF
	j := i * int64(bits)
	k := j & 4294967295
	if k < int64(bits) {
		for l := ^bits + 1%bits; k < int64(l); k = j & 4294967295 {
			i = x.generator.Next() & 0xFFFFFFFF
			j = i * int64(bits)
		}
	}
	i1 := j >> 32
	return int(i1)
}

func (x *XoroshiroSource) NextLong() int64 {
	return x.generator.Next()
}

func (x *XoroshiroSource) NextBoolean() bool {
	return x.NextLong()&1 != 0
}

func (x *XoroshiroSource) NextFloat() float32 {
	return float32(x.nextBits(24)) * 5.9604645e-8
}

func (x *XoroshiroSource) NextDouble() float64 {
	return float64(x.nextBits(53)) * 1.110223e-16
}

func (x *XoroshiroSource) NextGaussian() float64 {
	return x.gaussian.Next()
}

func (x *XoroshiroSource) nextBits(bits int) int64 {
	return x.NextLong() >> (64 - bits)
}
