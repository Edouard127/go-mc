package random

type RandomSource interface {
	SetSeed(seed int64)
	Next(bits int) int
	NextInt(n int) int
	NextLong() int64
	NextBoolean() bool
	NextFloat() float32
	NextDouble() float64
	NextGaussian() float64
}
