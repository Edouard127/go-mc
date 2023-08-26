package grids

type Beacon struct {
	*Generic
}

func NewBeacon(inventory *GenericInventory) *Beacon {
	return &Beacon{InitGenericContainer("minecraft:beacon", 9, 1, inventory)}
}
