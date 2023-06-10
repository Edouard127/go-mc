package random

import (
	"fmt"
	"reflect"
	"testing"
)

// AssertEqual checks if values are equal
func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Errorf("Received %v (type %v), expected %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
}

func TestUpgradeSeed(t *testing.T) {
	fmt.Println(UpgradeSeed(84978056))
	fmt.Println(UpgradeSeed(-8735875436988056))

	// output:
	//	[4120648339100951164 3023117482333613152]
	//	[7499608215528407055 610144910520614932]
}

func TestGenerateUniqueSeed(t *testing.T) {
	fmt.Println(GenerateUniqueSeed())
	fmt.Println(GenerateUniqueSeed())
	fmt.Println(GenerateUniqueSeed())
	fmt.Println(GenerateUniqueSeed())
}

func TestWorldGenRandom_SeedSlimeChunk(t *testing.T) {
	world := NewWorldGeneration(NewLegacyRandomSource(9209794931264193696))
	AssertEqual(t, true, world.SeedSlimeChunk(int32(3), int32(-11), 987234911).NextNInt(10) == 0 && world.NextNInt(10) == 0)
}
