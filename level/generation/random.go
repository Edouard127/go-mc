package generation

import (
	"sync/atomic"
	"time"
)

var seedUniquifier atomic.Int64

func init() {
	seedUniquifier.Store(8682522807148012)
}

type Seed128 [2]int64

// GenerateUniqueSeed generates a unique seed.
func GenerateUniqueSeed() int64 {
	return seedUniquifier.Add(181783497276652981) ^ time.Now().UnixNano()
}

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
