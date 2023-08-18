package grids

import (
	"fmt"
	"github.com/Edouard127/go-mc/data/slots"
)

type Generic struct {
	Name      string // Name of the grid.
	Type      int    // Type is the int corresponding to the window type.
	Size      int
	Data      []slots.Slot
	Inventory *GenericInventory
}

func InitGenericContainer(name string, id, size int, inventory *GenericInventory) *Generic {
	return &Generic{
		Name:      name,
		Type:      id,
		Size:      size,
		Data:      make([]slots.Slot, size),
		Inventory: inventory,
	}
}

func (g *Generic) OnClose() error { return nil }

func (g *Generic) GetSlot(i int) *slots.Slot {
	if i < 0 || i >= len(g.Inventory.Slots) {
		return nil
	}

	if i < g.Size {
		return &g.Data[i]
	} else {
		return &g.Inventory.Slots[g.Size:][i]
	}
}

func (g *Generic) SetSlot(i int, s slots.Slot) error {
	if i < g.Size {
		g.Data[i] = s
	} else {
		fmt.Println("SetSlot", i, s)
		g.Inventory.Slots[g.Size:][i] = s
	}

	return nil
}
