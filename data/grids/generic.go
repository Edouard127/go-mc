package grids

import (
	"fmt"
	"github.com/Edouard127/go-mc/data/slots"
)

type Container interface {
	GetSlot(int) *slots.Slot
	SetSlot(int, *slots.Slot) error
	OnClose() error
	GetType() int
	GetSize() int
}

// A Generic is a grid that can be used for any type of window.
// Generic is different from GenericInventory
// When you open a chest, the chest is a Generic, but the Generic
// has access to the player inventory, and places the items in
// the slots, either in the chest or in the player inventory.
// If index >= size, then the slot is in the player inventory.
// If index < size, then the slot is in the chest.
type Generic struct {
	Name string // Name of the grid.
	Type int    // Type is the int corresponding to the window type.
	Size int
	Data []*slots.Slot
}

func InitGenericContainer(name string, id, size int) *Generic {
	return &Generic{
		Name: name,
		Type: id,
		Size: size,
		Data: make([]*slots.Slot, size),
	}
}

func (g *Generic) OnClose() error { return nil }

func (g *Generic) GetSlot(i int) *slots.Slot {
	if i < 0 || i >= len(g.Data) {
		return nil
	}

	if i >= g.Size {
		return Containers[0].GetSlot(i - g.Size)
	}

	return g.Data[i]
}

func (g *Generic) SetSlot(i int, s *slots.Slot) error {
	if i < 0 || i >= g.Size+Containers[0].GetSize() {
		return fmt.Errorf("slot index %d out of bounds. maximum index is %d", i, len(g.Data)-1)
	}
	if i >= g.Size {
		return Containers[0].SetSlot(i-g.Size, s)
	}
	g.Data[i] = s
	return nil
}

func (g *Generic) GetType() int {
	return g.Type
}

func (g *Generic) GetSize() int {
	return g.Size
}
