package grids

type EnchantmentTable struct {
	*Generic
}

func NewEnchantmentTable(inventory *GenericInventory) *EnchantmentTable {
	return &EnchantmentTable{InitGenericContainer("minecraft:enchantment", 13, 2, inventory)}
}
