package random

import "math"

type MarsagliaPolarGaussian struct {
	Source  RandomSource
	next    float64
	hasNext bool
}

func NewMarsagliaPolarGaussian(source RandomSource) *MarsagliaPolarGaussian {
	return &MarsagliaPolarGaussian{Source: source}
}

func (m *MarsagliaPolarGaussian) Reset() {
	m.hasNext = false
}

func (m *MarsagliaPolarGaussian) Next() float64 {
	if !m.hasNext {
		u := 2*m.Source.NextDouble() - 1
		v := 2*m.Source.NextDouble() - 1
		s := u*u + v*v
	do:
		u = 2*m.Source.NextDouble() - 1
		v = 2*m.Source.NextDouble() - 1
		s = u*u + v*v
		if s >= 1 || s == 0 {
			goto do
		}
		d := math.Sqrt(-2 * math.Log(s) / s)
		m.hasNext = true
		m.next = v * d
		return u * d
	} else {
		m.hasNext = false
		return m.next
	}
}
