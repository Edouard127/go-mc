package slots

import (
	"github.com/Edouard127/go-mc/data/item"
	pk "github.com/Edouard127/go-mc/net/packet"
	"io"
)

type Slot struct {
	Index int // The index is relative to the position in the current container, this field should not be used
	ID    item.ID
	Count int8
	NBT   SlotTags
}

func (s *Slot) Item() item.Item {
	return item.ByID[s.ID]
}

func (s *Slot) GetIndex() int {
	return s.Index
}

func (s *Slot) WriteTo(w io.Writer) (n int64, err error) {
	var present pk.Boolean = s.ID != 0 && s.Count != 0
	return pk.Tuple{
		present, pk.Opt{
			If:    present,
			Value: pk.Tuple{&s.ID, &s.Count, pk.NBT(&s.NBT)},
		},
	}.WriteTo(w)
}

func (s *Slot) ReadFrom(r io.Reader) (n int64, err error) {
	var present pk.Boolean
	return pk.Tuple{
		&present, pk.Opt{
			If: &present,
			Value: pk.Tuple{
				(*pk.VarInt)(&s.ID), (*pk.Byte)(&s.Count), pk.NBT(&s.NBT),
			},
		},
	}.ReadFrom(r)
}

type Container interface {
	GetInventorySlots() []Slot
	GetHotbarSlots() []Slot
	OnSetSlot(i int, s Slot) error
	OnClose() error
}

type ChangedSlots map[int]*Slot

func (c ChangedSlots) WriteTo(w io.Writer) (n int64, err error) {
	n, err = pk.VarInt(len(c)).WriteTo(w)
	if err != nil {
		return
	}
	for i, v := range c {
		n1, err := pk.Short(i).WriteTo(w)
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

type SlotTags struct {
	Damage      int32     `nbt:"Damage"`
	Unbreakable bool      `nbt:"Unbreakable"`
	CanDestroy  []item.ID `nbt:"CanDestroy"`
}

func (s *SlotTags) WriteTo(w io.Writer) (n int64, err error) {
	return pk.NBT(s).WriteTo(w)
}

func (s *SlotTags) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.NBT(s).ReadFrom(r)
}
