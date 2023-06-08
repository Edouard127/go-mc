package random

// BitRandomSource acts as a wrapper for RandomSource.
type BitRandomSource struct {
	RandomSource
}

func (b *BitRandomSource) SetSeed(seed int64) {
	// NOP
}

func (b *BitRandomSource) Next(bits int) int {
	// not implemented
	return b.NextInt()
}

func (b *BitRandomSource) NextInt() int {
	return b.NextNInt(32)
}

func (b *BitRandomSource) NextNInt(n int) int {
	if n <= 0 {
		panic("n must be positive")
	} else if (n&n)-1 == 0 {
		return (n * b.NextNInt(31)) >> 31
	} else {
		var i int
		var j int
		for i-j-(n-1) < 0 {
			i = b.NextNInt(31)
			j = i % n
		}
		return j
	}
}

func (b *BitRandomSource) NextLong() int64 {
	return int64((b.NextNInt(32) << 32) + b.NextNInt(32))
}

func (b *BitRandomSource) NextBoolean() bool {
	return b.NextInt() != 0
}

func (b *BitRandomSource) NextFloat() float32 {
	return float32(float64(b.NextNInt(24)) * 5.9604645e-8)
}

func (b *BitRandomSource) NextDouble() float64 {
	return float64((b.NextNInt(26)<<27)+b.NextNInt(27)) * 1.110223e-16
}
