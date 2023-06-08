package random

type RandomSource interface {
	SetSeed(seed int64)
	Next() int
	NextInt() int
	NextNInt(n int) int
	NextLong() int64
	NextBoolean() bool
	NextFloat() float32
	NextDouble() float64
	NextGaussian() int64
}
