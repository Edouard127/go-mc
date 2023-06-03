package grids

import (
	"fmt"
	"github.com/Tnze/go-mc/data/item"
	"github.com/Tnze/go-mc/data/slots"
	pk "github.com/Tnze/go-mc/net/packet"
)

type GenericInventory struct {
	Slots [46]slots.Slot
}

func (g *GenericInventory) OnClose() error { return nil }

/* Slot data */

func (g *GenericInventory) GetInventory() []slots.Slot { return g.Slots[8 : len(g.Slots)-9] }
func (g *GenericInventory) GetHotbar() []slots.Slot    { return g.Slots[len(g.Slots)-9:] }

func (g *GenericInventory) GetItem(item item.Item, predicate func(slot slots.Slot) bool) slots.Slot {
	for i := range g.Slots {
		if g.Slots[i].ID == pk.VarInt(item.ID) && predicate(g.Slots[i]) {
			return g.Slots[i]
		}
	}
	return slots.Slot{}
}

func (g *GenericInventory) ApplyData(i []slots.Slot) {
	for _, v := range i {
		fmt.Println(v)
	}
}

/* Getter & Setter */

func (g *GenericInventory) GetSlot(i int) *slots.Slot { return &g.Slots[i] }
func (g *GenericInventory) SetSlot(i int, s slots.Slot) error {
	if i < 0 || i >= len(g.Slots) {
		return fmt.Errorf("slot index %d out of bounds. maximum index is %d", i, len(g.Slots)-1)
	}
	g.Slots[i] = s
	return nil
}

func (g *GenericInventory) GetCraftingOutput() *slots.Slot { return &g.Slots[0] }
func (g *GenericInventory) GetCraftingInput() []slots.Slot { return g.Slots[1:4] }
func (g *GenericInventory) GetArmor() []slots.Slot         { return g.Slots[5:8] }
func (g *GenericInventory) GetOffhand() *slots.Slot        { return &g.Slots[45] }

// TODO: Iterator for slots

func (g *GenericInventory) GetInventorySlot(id int) *slots.Slot {
	for i := range g.GetInventory() {
		if g.Slots[i].ID == pk.VarInt(id) {
			return &g.Slots[i]
		}
	}
	return nil
}
func (g *GenericInventory) GetHotbarSlot(id int) *slots.Slot {
	for i := range g.GetHotbar() {
		if g.Slots[i].ID == pk.VarInt(id) {
			return &g.Slots[i]
		}
	}
	return nil
}
