package grids

type Anvil struct {
	*Generic
}

func NewAnvil(inventory *GenericInventory) *Anvil {
	return &Anvil{InitGenericContainer("minecraft:anvil", 8, 3, inventory)}
}
