package block

import (
	"fmt"
	"github.com/Edouard127/go-mc/level/block/states"
	"testing"
)

// This is not the intended way of setting a property, if you set a property
// with another value, you will invalidate the block state array.
// But since I'm copying the block, there's no worry
func TestBlockStateHolder(t *testing.T) {
	temp := *TurtleEgg
	block := &temp
	// copying pointer keeps the map reference
	block.properties = make(map[states.Property]byte)

	// Initial print
	fmt.Println(block.Get(states.Hatch), block.Get(states.Eggs))

	// Simulate a block update
	block.set(states.Hatch, 2)
	block.set(states.Eggs, 4)

	// Updated print
	fmt.Println(block.Get(states.Hatch), block.Get(states.Eggs))
	fmt.Println(TurtleEgg.Get(states.Hatch), TurtleEgg.Get(states.Eggs))
}
