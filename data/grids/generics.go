package grids

import (
	"github.com/Edouard127/go-mc/data/packetid"
	"github.com/Edouard127/go-mc/data/slots"
	"github.com/Edouard127/go-mc/net"
	pk "github.com/Edouard127/go-mc/net/packet"
)

type Generic3x3 struct {
	*Generic
}

func NewGeneric3x3() *Generic3x3 {
	return &Generic3x3{InitGenericContainer("minecraft:generic_3x3", 7, 9)}
}

type Generic9x1 struct {
	*Generic
}

func NewGeneric9x1() *Generic9x1 {
	return &Generic9x1{InitGenericContainer("minecraft:generic_9x1", 1, 9)}
}

type Generic9x2 struct {
	*Generic
}

func NewGeneric9x2() *Generic9x2 {
	return &Generic9x2{InitGenericContainer("minecraft:generic_9x2", 2, 18)}
}

type Generic9x3 struct {
	*Generic
}

func NewGeneric9x3() *Generic9x3 {
	return &Generic9x3{InitGenericContainer("minecraft:generic_9x3", 3, 27)}
}

type Generic9x4 struct {
	*Generic
}

func NewGeneric9x4() *Generic9x4 {
	return &Generic9x4{InitGenericContainer("minecraft:generic_9x4", 4, 36)}
}

type Generic9x5 struct {
	*Generic
}

func NewGeneric9x5() *Generic9x4 {
	return &Generic9x4{InitGenericContainer("minecraft:generic_9x5", 5, 45)}
}

type Generic9x6 struct {
	*Generic
}

func NewGeneric9x6() *Generic9x6 {
	return &Generic9x6{InitGenericContainer("minecraft:generic_9x6", 6, 54)}
}

type Anvil struct {
	*Generic
}

func NewAnvil() *Anvil {
	return &Anvil{InitGenericContainer("minecraft:anvil", 8, 3)}
}

func (a *Anvil) SetFirstItem(slot *slots.Slot) error {
	return a.SetSlot(0, slot)
}

func (a *Anvil) SetSecondItem(slot *slots.Slot) error {
	return a.SetSlot(1, slot)
}

func (a *Anvil) GetOutputItem() *slots.Slot {
	return a.GetSlot(2)
}

func (a *Anvil) SetOutputName(conn *net.Conn, name string) error {
	return conn.WritePacket(pk.Marshal(packetid.SPacketRenameItem, pk.String(name[:min(len(name), 50)])))
}

func (a *Anvil) getRepairCost() int {
	return 0
}

type Beacon struct {
	*Generic
}

func NewBeacon() *Beacon {
	return &Beacon{InitGenericContainer("minecraft:beacon", 9, 1)}
}

type BlastFurnace struct {
	*Generic
}

func NewBlastFurnace() *BlastFurnace {
	return &BlastFurnace{InitGenericContainer("minecraft:blast_furnace", 10, 3)}
}

type BrewingStand struct {
	*Generic
}

func NewBrewingStand() *BrewingStand {
	return &BrewingStand{InitGenericContainer("minecraft:brewing_stand", 11, 5)}
}

type CartographyTable struct {
	*Generic
}

func NewCartographyTable() *CartographyTable {
	return &CartographyTable{InitGenericContainer("minecraft:cartography", 23, 3)}
}

type CraftingTable struct {
	*Generic
}

func NewCraftingTable() *CraftingTable {
	return &CraftingTable{InitGenericContainer("minecraft:crafting", 12, 10)}
}

type EnchantmentTable struct {
	*Generic
}

func NewEnchantmentTable() *EnchantmentTable {
	return &EnchantmentTable{InitGenericContainer("minecraft:enchantment", 13, 2)}
}

type Furnace struct {
	*Generic
}

func NewFurnace() *Furnace {
	return &Furnace{InitGenericContainer("minecraft:furnace", 13, 4)}
}

type Grindstone struct {
	*Generic
}

func NewGrindstone() *Grindstone {
	return &Grindstone{InitGenericContainer("minecraft:grindstone", 15, 3)}
}

type Hopper struct { // Also minecart with hopper
	*Generic
}

func NewHopper() *Hopper {
	return &Hopper{InitGenericContainer("minecraft:hopper", 16, 5)}
}

type Loom struct {
	*Generic
}

func NewLoom() *Loom {
	return &Loom{InitGenericContainer("minecraft:loom", 18, 4)}
}

type Merchant struct {
	*Generic
}

func NewMerchant() *Merchant {
	return &Merchant{InitGenericContainer("minecraft:merchant", 19, 3)}
}

type ShulkerBox struct {
	*Generic
}

func NewShulkerBox() *ShulkerBox {
	return &ShulkerBox{InitGenericContainer("minecraft:shulker_box", 20, 27)}
}

type SmithingTable struct {
	*Generic
}

func NewSmithingTable() *SmithingTable {
	return &SmithingTable{InitGenericContainer("minecraft:smithing", 21, 4)}
}

type Smoker struct {
	*Generic
}

func NewSmoker() *Smoker {
	return &Smoker{InitGenericContainer("minecraft:smoker", 22, 3)}
}

type Stonecutter struct {
	*Generic
}

func NewStonecutter() *Stonecutter {
	return &Stonecutter{InitGenericContainer("minecraft:stonecutter", 23, 2)}
}

var Containers = map[int]Container{
	0:  new(GenericInventory),
	1:  NewGeneric9x1(),
	2:  NewGeneric9x2(),
	3:  NewGeneric9x3(),
	4:  NewGeneric9x4(),
	5:  NewGeneric9x5(),
	6:  NewGeneric9x6(),
	7:  NewGeneric3x3(),
	8:  NewAnvil(),
	9:  NewBeacon(),
	10: NewBlastFurnace(),
	11: NewBrewingStand(),
	12: NewCraftingTable(),
	13: NewEnchantmentTable(),
	14: NewFurnace(),
	15: NewGrindstone(),
	16: NewHopper(),
	17: InitGenericContainer("nil", 0, 0), // TODO: This is the only one that is not a container, I don't know why mojang did this.
	18: NewLoom(),
	19: NewMerchant(),
	20: NewShulkerBox(),
	21: NewSmithingTable(),
	22: NewSmoker(),
	23: NewCartographyTable(),
	24: NewStonecutter(),
}
