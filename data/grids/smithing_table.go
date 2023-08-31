package grids

type SmithingTable struct {
	*Generic
}

func NewSmithingTable(inventory *GenericInventory) *SmithingTable {
	return &SmithingTable{InitGenericContainer("minecraft:smithing", 21, 4, inventory)}
}
