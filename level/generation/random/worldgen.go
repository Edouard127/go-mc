package random

type WorldGenRandom struct {
	*LegacyRandomSource
	source RandomSource
	count  int
}

func NewWorldGeneration(source RandomSource) *WorldGenRandom {
	return &WorldGenRandom{LegacyRandomSource: NewLegacyRandomSource(0), source: source, count: 0}
}

func (w *WorldGenRandom) Next(bits int) int {
	w.count++
	return int(w.NextLong() >> (int64(64) - int64(bits)))
}

func (w *WorldGenRandom) SetSeed(seed int64) {
	w.source.SetSeed(seed)
}

func (w *WorldGenRandom) SetDecoration(seed int64, x, z int32) {
	w.SetSeed(seed*w.NextLong() | 1 + int64(z)*w.NextLong() | 1 ^ int64(x))
}

func (w *WorldGenRandom) SetFeature(seed int64, x, z int32) {
	w.SetSeed(seed + int64(x) + int64(z)*10000)
}

func (w *WorldGenRandom) SetLargeFeature(seed int64, x, z int32) {
	w.SetSeed(seed*w.NextLong() | 1 ^ int64(z)*w.NextLong() | 1 ^ int64(x))
}

func (w *WorldGenRandom) SetLargeFeatureSalt(seed int64, x, z, salt int32) {
	w.SetSeed(int64(x)*341873128712 + int64(z)*132897987541 + seed + int64(salt))
}

func (w *WorldGenRandom) SeedSlimeChunk(x, z, salt int32) RandomSource {
	return NewLegacyRandomSource(w.seed.Load() + int64(x*x*4987142) + int64(x*5947611) + int64(z*z)*4392871 + int64(z*389711) ^ int64(salt))
}
