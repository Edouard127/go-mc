package grids

import (
	"fmt"
	"github.com/Edouard127/go-mc/data/item"
	"github.com/Edouard127/go-mc/data/slots"
)

type GenericInventory struct {
	Slots [46]*slots.Slot
}

func (g *GenericInventory) OnClose() error { return nil }

func (g *GenericInventory) GetSlot(i int) *slots.Slot { return g.Slots[i] }
func (g *GenericInventory) SetSlot(i int, s *slots.Slot) error {
	if i < 0 || i >= g.GetSize() {
		return fmt.Errorf("slot index %d out of bounds. maximum index is %d", i, len(g.Slots)-1)
	}

	g.Slots[i] = s
	return nil
}
func (g *GenericInventory) GetItem(i int) item.Item { return g.Slots[i].Item() }
func (g *GenericInventory) GetType() int            { return 0 }
func (g *GenericInventory) GetSize() int            { return len(g.Slots) }

func (g *GenericInventory) GetCraftingOutput() item.Item  { return g.Slots[0].Item() }
func (g *GenericInventory) GetCraftingInput() []item.Item { return g.itemsOf(1, 4) }
func (g *GenericInventory) GetArmor() []item.Item         { return g.itemsOf(4, 8) }
func (g *GenericInventory) GetInventory() []item.Item     { return g.itemsOf(9, 35) }
func (g *GenericInventory) GetHotbar() []item.Item        { return g.itemsOf(36, 44) }
func (g *GenericInventory) GetOffhand() item.Item         { return g.Slots[45].Item() }

func (g *GenericInventory) itemsOf(start, end int) (items []item.Item) {
	for i := start; i < end; i++ {
		items = append(items, g.Slots[i].Item())
	}
	return
}

func (g *GenericInventory) FindItem(predicate func(item.Item) bool) item.Item {
	return g.predicate(1, 0, 45, predicate)
}

func (g *GenericInventory) FindItemNth(nth int, predicate func(item.Item) bool) item.Item {
	return g.predicate(nth, 0, 45, predicate)
}

func (g *GenericInventory) FindItemNthPoint(nth int, predicate func(item.Item) bool, start, end int) item.Item {
	return g.predicate(nth, start, end, predicate)
}

func (g *GenericInventory) predicate(nth, start, end int, predicate func(item.Item) bool) item.Item {
	var item item.Item
	if predicate == nil {
		return item
	}

	for i := start; i < end; i++ {
		item = g.Slots[i].Item()
		if predicate(item) {
			nth--
			if nth == 0 {
				return item
			}
		}
	}
	return item
}
