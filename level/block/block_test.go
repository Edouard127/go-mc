package block

import (
	"fmt"
	"github.com/Edouard127/go-mc/level/block/states"
	"testing"
)

func TestBlock(t *testing.T) {
	fmt.Println(SandstoneWall.StateHolder.Neighbors)
}

func TestBlockStateHolder(t *testing.T) {
	var block = TurtleEgg

	fmt.Println("Default state id", block.Default())

	// Initial print
	fmt.Println(block.GetValue(states.HatchProperty), block.GetValue(states.EggsProperty), ToStateID[block])

	// Simulate a block update
	block.SetValue(states.HatchProperty, 2)
	block.SetValue(states.EggsProperty, 4)

	// Updated print
	fmt.Println(block.GetValue(states.HatchProperty), block.GetValue(states.EggsProperty), ToStateID[block])
}
