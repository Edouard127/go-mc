package screen

import (
	"github.com/Edouard127/go-mc/data/grids"
	"github.com/Edouard127/go-mc/data/item"
	"github.com/Edouard127/go-mc/data/slots"
)

type Manager struct {
	Cursor        item.Item
	HeldItem      item.Item
	CurrentScreen *Container
	Screens       map[int]Container
	Inventory     *grids.GenericInventory
	// The last state received from the server
	StateID int32
}

func NewManager() *Manager {
	inventory := new(grids.GenericInventory)
	return &Manager{
		Screens:   fillContainers(inventory),
		Inventory: inventory,
	}
}

func (m *Manager) OpenScreen(id int32) {
	c := m.Screens[int(id)]
	m.CurrentScreen = &c
}

func (m *Manager) CloseScreen() {
	m.CurrentScreen = nil
	m.Cursor = item.Item{}
}

func fillContainers(inventory *grids.GenericInventory) map[int]Container {
	return map[int]Container{
		1:  grids.NewGeneric9x1(inventory),
		2:  grids.NewGeneric9x2(inventory),
		3:  grids.NewGeneric9x3(inventory),
		4:  grids.NewGeneric9x4(inventory),
		5:  grids.NewGeneric9x5(inventory),
		6:  grids.NewGeneric9x6(inventory),
		7:  grids.NewGeneric3x3(inventory),
		8:  grids.NewAnvil(inventory),
		9:  grids.NewBeacon(inventory),
		10: grids.NewBlastFurnace(inventory),
		11: grids.NewBrewingStand(inventory),
		12: grids.NewCraftingTable(inventory),
		13: grids.NewEnchantmentTable(inventory),
		14: grids.NewFurnace(inventory),
		15: grids.NewGrindstone(inventory),
		16: grids.NewHopper(inventory),
		17: grids.InitGenericContainer("nil", 0, 0, inventory), // TODO: This is the only one that is not a container, I don't know why mojang did this.
		18: grids.NewLoom(inventory),
		19: grids.NewMerchant(inventory),
		20: grids.NewShulkerBox(inventory),
		21: grids.NewSmithingTable(inventory),
		22: grids.NewSmoker(inventory),
		23: grids.NewCartographyTable(inventory),
		24: grids.NewStonecutter(inventory),
	}
}

type Container interface {
	GetSlot(int) *slots.Slot
	SetSlot(int, slots.Slot) error
	OnClose() error
}
