package random

import (
	"sync/atomic"
	"time"
)

var seedUniquifier int64

func init() {
	atomic.AddInt64(&seedUniquifier, 8682522807148012)
}

type Seed128 [2]int64

// GenerateUniqueSeed generates a unique seed.
func GenerateUniqueSeed() int64 {
	return atomic.AddInt64(&seedUniquifier, seedUniquifier*1181783497276652981) ^ time.Now().UnixNano()*8394769045352766035
}

// UpgradeSeed takes a int64 seed and transforms it into a Seed128
func UpgradeSeed(seed int64) Seed128 {
	lower := seed ^ 7640891576956012809
	upper := lower + -7046029254386353131
	return Seed128{Stafford13(lower), Stafford13(upper)}
}

func Stafford13(seed int64) int64 {
	seed = (seed ^ seed>>30) * -4658895280553007687
	seed = (seed ^ seed>>27) * -7723592293110705685
	return seed ^ seed>>31
}
