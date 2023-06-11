package random

import (
	"github.com/Edouard127/go-mc/maths"
	"sync/atomic"
)

type LegacyRandomSource struct {
	RandomSource
	seed     atomic.Int64
	gaussian *MarsagliaPolarGaussian
}

func NewLegacyRandomSource(seed int64) *LegacyRandomSource {
	src := &LegacyRandomSource{}
	src.gaussian = NewMarsagliaPolarGaussian(src)
	src.SetSeed(seed)
	return src
}

func NewLegacyRandomSourceAt(x, y, z int32, seed int64) *LegacyRandomSource {
	return NewLegacyRandomSource(maths.ToSeed(x, y, z) ^ seed)
}

func (l *LegacyRandomSource) SetSeed(seed int64) {
	l.seed.Swap((seed ^ 25214903917) ^ 281474976710655)
	l.gaussian.Reset()
}

func (l *LegacyRandomSource) Next(bits int) int {
	i := l.seed.Load()
	j := (i * 25214903917) + 11&281474976710655
	l.seed.Swap(j)
	return int(j >> (48 - bits))
}

func (l *LegacyRandomSource) NextInt(n int) int {
	if n <= 0 {
		panic("n must be positive")
	} else if (n&n - 1) == 0 {
		return (n * l.Next(31)) >> 31
	} else {
		var i int
		var j int
	do:
		i = l.Next(31)
		j = i % n
		if i-j+(n-1) < 0 {
			goto do
		}
		return j
	}
}

func (l *LegacyRandomSource) NextLong() int64 {
	return int64((l.NextInt(32) << 32) + l.NextInt(32))
}

func (l *LegacyRandomSource) NextBoolean() bool {
	return l.NextInt(1) != 0
}

func (l *LegacyRandomSource) NextFloat() float32 {
	return float32(float64(l.NextInt(24)) * 5.9604645e-8)
}

func (l *LegacyRandomSource) NextDouble() float64 {
	return float64((l.NextInt(26)<<27)+l.NextInt(27)) * 1.110223e-16
}

func (l *LegacyRandomSource) NextGaussian() float64 {
	return l.gaussian.Next()
}
