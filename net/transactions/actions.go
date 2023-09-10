package transactions

import (
	"github.com/Edouard127/go-mc/bot/screen"
	"github.com/Edouard127/go-mc/data/slots"
)

func LeftClick(item *slots.Slot) []*SlotAction {
	return []*SlotAction{NewSlotAction(screen.LeftClick, item.GetIndex(), 0, &slots.Slot{Index: item.Index})}
}

func RightClick(item *slots.Slot) []*SlotAction {
	return []*SlotAction{NewSlotAction(screen.RightClick, item.GetIndex(), 0, &slots.Slot{Index: item.Index})}
}

func DoubleClick(item *slots.Slot) []*SlotAction {
	return []*SlotAction{NewSlotAction(screen.DoubleClick, item.GetIndex(), 6, &slots.Slot{Index: item.Index})}
}

func Drop(item *slots.Slot) []*SlotAction {
	return []*SlotAction{NewSlotAction(screen.Drop, item.GetIndex(), 4, &slots.Slot{Index: item.Index})}
}

func DropAll(item *slots.Slot) []*SlotAction {
	return []*SlotAction{NewSlotAction(screen.ControlDrop, item.GetIndex(), 4, &slots.Slot{Index: item.Index})}
}

func Swap(item1 *slots.Slot, item2 *slots.Slot) []*SlotAction {
	return []*SlotAction{
		NewSlotAction(screen.LeftClick, item1.GetIndex(), 0, item1, &slots.Slot{Index: item2.Index}),
		NewSlotAction(screen.LeftClick, item2.GetIndex(), 0, item1, item2),
		NewSlotAction(screen.LeftClick, item1.GetIndex(), 0, item2, &slots.Slot{Index: item2.Index}),
		NewSlotAction(screen.LeftClick, -999, 4, &slots.Slot{}, &slots.Slot{}),
	}
}

func SwapWithOffhand(item *slots.Slot) []*SlotAction {
	return []*SlotAction{NewSlotAction(screen.SwapHand, item.GetIndex(), 2, item, &slots.Slot{Index: 40})}
}

func QuickMove(item *slots.Slot) []*SlotAction {
	return []*SlotAction{NewSlotAction(screen.ShiftLeftClick, item.GetIndex(), 1, &slots.Slot{Index: item.Index})}
}
