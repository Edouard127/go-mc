package screen

import (
	"fmt"
	"github.com/Edouard127/go-mc/data/grids"
	"github.com/Edouard127/go-mc/data/slots"
	"testing"
)

func TestNewManager(t *testing.T) {
	m := NewManager()
	m.Inventory.SetSlot(2, &slots.Slot{
		Index: 6969,
		ID:    0,
		Count: 0,
	})
	grids.Containers[2].SetSlot(2, &slots.Slot{Index: 420})
	fmt.Println(grids.Containers[2].GetSlot(2), m.Inventory.GetSlot(2))
}
