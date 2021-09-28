package azmath

import (
	"math"
)

const (
	quake64 = 0x5FE6EB50C7B537A9
)

func FastInvSqrt64(n float64) float64 {
	if n < 0 {
		return math.NaN()
	}
	halfN := n * 0.5
	threeHalves := float64(1.5)
	b := math.Float64bits(n)
	b = quake64 - (b >> 1)
	f := math.Float64frombits(b)
	f *= threeHalves - (halfN * f * f)
	return f
}
