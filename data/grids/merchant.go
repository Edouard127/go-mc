package grids

type Merchant struct {
	*Generic
}

func NewMerchant(inventory *GenericInventory) *Merchant {
	return &Merchant{InitGenericContainer("minecraft:merchant", 19, 3, inventory)}
}
