package grids

type Hopper struct { // Also minecart with hopper
	*Generic
}

func NewHopper(inventory *GenericInventory) *Hopper {
	return &Hopper{InitGenericContainer("minecraft:hopper", 16, 5, inventory)}
}
