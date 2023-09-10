package screen

import (
	"fmt"
	"github.com/Edouard127/go-mc/data/slots"
	"github.com/Edouard127/go-mc/nbt"
	"testing"
)

func TestNewManager(t *testing.T) {
	m := NewManager()
	m.Inventory.SetSlot(2, &slots.Slot{
		Index: 6969,
		ID:    0,
		Count: 0,
		NBT:   nbt.RawMessage{},
	})
	Containers[2].SetSlot(2, &slots.Slot{Index: 420})
	fmt.Println(m.Inventory.GetSlot(2))
	m.Inventory.OpenWith(Containers[2])
	fmt.Println(m.Inventory.GetSlot(2))
}
