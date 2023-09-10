package grids

import (
	"fmt"
	"github.com/Edouard127/go-mc/data/item"
	"github.com/Edouard127/go-mc/data/slots"
)

type GenericInventory struct {
	Offset Container
	Slots  [46]*slots.Slot
}

func (g *GenericInventory) OpenWith(c Container) { g.Offset = c }
func (g *GenericInventory) OnClose() error {
	var err error
	if g.Offset != nil {
		err = g.Offset.OnClose()
		g.Offset = nil
	}
	return err
}

func (g *GenericInventory) GetSlot(i int) *slots.Slot { return g.slotOffset(i) }
func (g *GenericInventory) SetSlot(i int, s *slots.Slot) error {
	if i < 0 || g.Offset != nil && i >= g.Offset.GetSize()+g.GetSize() || i >= g.GetSize() {
		return fmt.Errorf("slot index %d out of bounds. maximum index is %d", i, len(g.Slots)-1)
	}
	if g.Offset != nil && i < g.Offset.GetSize() {
		return g.Offset.SetSlot(i, s)
	}
	g.Slots[i] = s
	return nil
}
func (g *GenericInventory) GetItem(i int) item.Item { return g.slotOffset(i).Item() }
func (g *GenericInventory) GetType() int            { return 0 }
func (g *GenericInventory) GetSize() int            { return len(g.Slots) }

func (g *GenericInventory) GetCraftingOutput() item.Item  { return g.slotOffset(0).Item() }
func (g *GenericInventory) GetCraftingInput() []item.Item { return g.itemsOf(1, 4) }
func (g *GenericInventory) GetArmor() []item.Item         { return g.itemsOf(4, 8) }
func (g *GenericInventory) GetInventory() []item.Item     { return g.itemsOf(9, 35) }
func (g *GenericInventory) GetHotbar() []item.Item        { return g.itemsOf(36, 44) }
func (g *GenericInventory) GetOffhand() item.Item         { return g.Slots[45].Item() }

func (g *GenericInventory) itemsOf(start, end int) (items []item.Item) {
	for i := start; i < end; i++ {
		items = append(items, g.slotOffset(i).Item())
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

func (g *GenericInventory) predicate(nth, start, end int, predicate func(item.Item) bool) (item item.Item) {
	if predicate == nil {
		return
	}
	for i := start; i < end; i++ {
		if predicate(g.slotOffset(i).Item()) {
			nth--
			if nth == 0 {
				return g.slotOffset(i).Item()
			}
		}
	}
	return
}

func (g *GenericInventory) slotOffset(index int) *slots.Slot {
	if g.Offset != nil && g.Offset.GetSize() < index+1 {
		return g.Offset.GetSlot(index)
	}

	return g.Slots[index]
}
