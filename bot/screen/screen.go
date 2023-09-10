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
	return &Manager{Inventory: Containers[0].(*grids.GenericInventory)}
}

func (m *Manager) OpenScreen(id int32) {
	m.CurrentScreen = Containers[int(id)]
}

func (m *Manager) CloseScreen() {
	m.CurrentScreen = nil
}

var Containers = map[int]grids.Container{
	0:  new(grids.GenericInventory),
	1:  grids.NewGeneric9x1(),
	2:  grids.NewGeneric9x2(),
	3:  grids.NewGeneric9x3(),
	4:  grids.NewGeneric9x4(),
	5:  grids.NewGeneric9x5(),
	6:  grids.NewGeneric9x6(),
	7:  grids.NewGeneric3x3(),
	8:  grids.NewAnvil(),
	9:  grids.NewBeacon(),
	10: grids.NewBlastFurnace(),
	11: grids.NewBrewingStand(),
	12: grids.NewCraftingTable(),
	13: grids.NewEnchantmentTable(),
	14: grids.NewFurnace(),
	15: grids.NewGrindstone(),
	16: grids.NewHopper(),
	17: grids.InitGenericContainer("nil", 0, 0), // TODO: This is the only one that is not a container, I don't know why mojang did this.
	18: grids.NewLoom(),
	19: grids.NewMerchant(),
	20: grids.NewShulkerBox(),
	21: grids.NewSmithingTable(),
	22: grids.NewSmoker(),
	23: grids.NewCartographyTable(),
	24: grids.NewStonecutter(),
}
