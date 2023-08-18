package random

import (
	"fmt"
	"github.com/Edouard127/go-mc/internal/util"
	"testing"
)

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
	util.AssertEqual(t, true, world.SeedSlimeChunk(int32(3), int32(-11), 987234911).NextInt(10) == 0 && world.NextInt(10) == 0)
}
