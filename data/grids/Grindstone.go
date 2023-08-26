package grids

type Grindstone struct {
	*Generic
}

func NewGrindstone(inventory *GenericInventory) *Grindstone {
	return &Grindstone{InitGenericContainer("minecraft:grindstone", 15, 3, inventory)}
}
