package transactions

import (
	"fmt"
	"github.com/Edouard127/go-mc/bot/screen"
	"github.com/Edouard127/go-mc/data/slots"
	pk "github.com/Edouard127/go-mc/net/packet"
	"io"
)

type SlotAction struct {
	Slot    pk.Short
	Button  pk.Byte
	Mode    pk.VarInt
	Changed ChangedSlots
	Item    *slots.Slot
}

func NewSlotAction(button screen.Button, slot int, mode screen.Mode, cursor *slots.Slot, items ...*slots.Slot) *SlotAction {
	return &SlotAction{
		Slot:    pk.Short(slot),
		Button:  pk.Byte(button),
		Mode:    pk.VarInt(mode),
		Changed: items,
		Item:    cursor,
	}
}

func (s *SlotAction) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Tuple{&s.Slot, &s.Button, &s.Mode, &s.Changed, s.Item}.WriteTo(w)
}

func (s *SlotAction) Validate() error {
	if s.Slot < 0 && s.Slot != -999 {
		return fmt.Errorf("slot %d is less than 0", s.Slot)
	}
	if s.Button < 0 || s.Button > 40 {
		return fmt.Errorf("button %d is less than 0", s.Button)
	}
	if s.Mode < 0 || s.Mode > 6 {
		return fmt.Errorf("mode %d is not in range [0, 6]", s.Mode)
	}
	return nil
}

type ChangedSlots []*slots.Slot

func (c ChangedSlots) WriteTo(w io.Writer) (n int64, err error) {
	n0, err := pk.VarInt(len(c)).WriteTo(w)
	if err != nil {
		return n + n0, err
	}
	for _, v := range c {
		n1, err := pk.Short(v.Index).WriteTo(w)
		if err != nil {
			return n + n1, err
		}
		n2, err := v.WriteTo(w)
		if err != nil {
			return n + n1 + n2, err
		}
		n += n1 + n2
	}
	return
}
