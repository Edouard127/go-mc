package grids

type Loom struct {
	*Generic
}

func NewLoom(inventory *GenericInventory) *Loom {
	return &Loom{InitGenericContainer("minecraft:loom", 18, 4, inventory)}
}
