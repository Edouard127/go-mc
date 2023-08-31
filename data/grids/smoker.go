package grids

type Smoker struct {
	*Generic
}

func NewSmoker(inventory *GenericInventory) *Smoker {
	return &Smoker{InitGenericContainer("minecraft:smoker", 22, 3, inventory)}
}
