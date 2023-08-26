package grids

type BlastFurnace struct {
	*Generic
}

func NewBlastFurnace(inventory *GenericInventory) *BlastFurnace {
	return &BlastFurnace{InitGenericContainer("minecraft:blast_furnace", 10, 3, inventory)}
}
