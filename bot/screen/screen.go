package screen

import (
	"github.com/Edouard127/go-mc/data/grids"
	"github.com/Edouard127/go-mc/data/item"
)

type Manager struct {
	CurrentScreen grids.Container
	HeldItem      item.Item
	Inventory     *grids.GenericInventory
	StateID       int32 // The last state received from the server
}

func NewManager() *Manager {
	return &Manager{Inventory: grids.Containers[0].(*grids.GenericInventory)}
}

func (m *Manager) OpenScreen(id int32) {
	m.CurrentScreen = grids.Containers[int(id)]
}

func (m *Manager) CloseScreen() {
	m.CurrentScreen = nil
}
