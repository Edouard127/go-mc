package block

import (
	"fmt"
	"testing"
)

func TestBlock(t *testing.T) {
	fmt.Println(SandstoneWall.StateHolder.Neighbors)
}

func TestBlockStateHolder(t *testing.T) {
	var block = TurtleEgg

	fmt.Println("Default state id", block.Default())

	// Initial print
	fmt.Println(block.GetValue(HatchProperty), block.GetValue(EggsProperty), ToStateID[block])

	// Simulate a block update
	block.SetValue(HatchProperty, 2)
	block.SetValue(EggsProperty, 4)

	// Updated print
	fmt.Println(block.GetValue(HatchProperty), block.GetValue(EggsProperty), ToStateID[block])
}
