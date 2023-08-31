package grids

type CartographyTable struct {
	*Generic
}

func NewCartographyTable(inventory *GenericInventory) *CartographyTable {
	return &CartographyTable{InitGenericContainer("minecraft:cartography", 23, 3, inventory)}
}
