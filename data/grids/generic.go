package grids

import (
	"github.com/Edouard127/go-mc/data/slots"
)

type Container interface {
	GetSlot(int) *slots.Slot
	SetSlot(int, *slots.Slot) error
	OnClose() error
	GetType() int
	GetSize() int
}

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

	return g.Data[i]
}

func (g *Generic) SetSlot(i int, s *slots.Slot) error {
	g.Data[i] = s
	return nil
}

func (g *Generic) GetType() int {
	return g.Type
}

func (g *Generic) GetSize() int {
	return g.Size
}
