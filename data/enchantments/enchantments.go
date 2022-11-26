// Code generated by gen_enchantments.go DO NOT EDIT.
// Package enchantments stores information about enchantments in Minecraft.
package enchantments

// ID describes the numeric ID of an enchantment.
type ID uint32

// For more informations about the MinCost and MaxCost fields, see
// https://github.com/PrismarineJS/prismarine-rng/blob/cd4fd9eeda6ea72e428d172a08a99ff4e4ac0394/lib/enchantments.js#L72 and
// https://github.com/PrismarineJS/minecraft-data-generator-server/blob/4395173dd97148f545b3392540d12bc70ebe0bf5/src/main/java/dev/u9g/minecraftdatagenerator/generators/EnchantmentsDataGenerator.java#L83

// MinCost describes the minimum cost of an enchantment.
type MinCost struct {
	Level int32
	Cost  int32
}

// MaxCost describes the maximum cost of an enchantment.
type MaxCost struct {
	Level int32
	Cost  int32
}

// Enchantment describes information about a type of enchantment.
type Enchantment struct {
	ID           ID
	DisplayName  string
	Name         string
	MaxLevel     uint
	MinCost      MinCost
	MaxCost      MaxCost
	Exclude      []string
	Category     string
	Weight       int32
	TreasureOnly bool
	Curse        bool
	Tradeable    bool
	Discoverable bool
}

var (
	Protection = Enchantment{
		ID:          0,
		DisplayName: "Protection",
		Name:        "protection",
		MaxLevel:    4,
		MinCost: MinCost{
			Level: 11,
			Cost:  -10,
		},
		MaxCost: MaxCost{
			Level: 11,
			Cost:  1,
		},
		Exclude: []string{
			"fire_protection",
			"blast_protection",
			"projectile_protection",
		},
		Category:     "armor",
		Weight:       10,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	FireProtection = Enchantment{
		ID:          1,
		DisplayName: "Fire Protection",
		Name:        "fire_protection",
		MaxLevel:    4,
		MinCost: MinCost{
			Level: 8,
			Cost:  2,
		},
		MaxCost: MaxCost{
			Level: 8,
			Cost:  10,
		},
		Exclude: []string{
			"protection",
			"blast_protection",
			"projectile_protection",
		},
		Category:     "armor",
		Weight:       5,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	FeatherFalling = Enchantment{
		ID:          2,
		DisplayName: "Feather Falling",
		Name:        "feather_falling",
		MaxLevel:    4,
		MinCost: MinCost{
			Level: 6,
			Cost:  -1,
		},
		MaxCost: MaxCost{
			Level: 6,
			Cost:  5,
		},
		Exclude:      []string{},
		Category:     "armor_feet",
		Weight:       5,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	BlastProtection = Enchantment{
		ID:          3,
		DisplayName: "Blast Protection",
		Name:        "blast_protection",
		MaxLevel:    4,
		MinCost: MinCost{
			Level: 8,
			Cost:  -3,
		},
		MaxCost: MaxCost{
			Level: 8,
			Cost:  5,
		},
		Exclude: []string{
			"protection",
			"fire_protection",
			"projectile_protection",
		},
		Category:     "armor",
		Weight:       2,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	ProjectileProtection = Enchantment{
		ID:          4,
		DisplayName: "Projectile Protection",
		Name:        "projectile_protection",
		MaxLevel:    4,
		MinCost: MinCost{
			Level: 6,
			Cost:  -3,
		},
		MaxCost: MaxCost{
			Level: 6,
			Cost:  3,
		},
		Exclude: []string{
			"protection",
			"fire_protection",
			"blast_protection",
		},
		Category:     "armor",
		Weight:       5,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Respiration = Enchantment{
		ID:          5,
		DisplayName: "Respiration",
		Name:        "respiration",
		MaxLevel:    3,
		MinCost: MinCost{
			Level: 10,
			Cost:  0,
		},
		MaxCost: MaxCost{
			Level: 10,
			Cost:  30,
		},
		Exclude:      []string{},
		Category:     "armor_head",
		Weight:       2,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	AquaAffinity = Enchantment{
		ID:          6,
		DisplayName: "Aqua Affinity",
		Name:        "aqua_affinity",
		MaxLevel:    1,
		MinCost: MinCost{
			Level: 0,
			Cost:  1,
		},
		MaxCost: MaxCost{
			Level: 0,
			Cost:  41,
		},
		Exclude:      []string{},
		Category:     "armor_head",
		Weight:       2,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Thorns = Enchantment{
		ID:          7,
		DisplayName: "Thorns",
		Name:        "thorns",
		MaxLevel:    3,
		MinCost: MinCost{
			Level: 20,
			Cost:  -10,
		},
		MaxCost: MaxCost{
			Level: 10,
			Cost:  51,
		},
		Exclude:      []string{},
		Category:     "armor_chest",
		Weight:       1,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	DepthStrider = Enchantment{
		ID:          8,
		DisplayName: "Depth Strider",
		Name:        "depth_strider",
		MaxLevel:    3,
		MinCost: MinCost{
			Level: 10,
			Cost:  0,
		},
		MaxCost: MaxCost{
			Level: 10,
			Cost:  15,
		},
		Exclude: []string{
			"frost_walker",
		},
		Category:     "armor_feet",
		Weight:       2,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	FrostWalker = Enchantment{
		ID:          9,
		DisplayName: "Frost Walker",
		Name:        "frost_walker",
		MaxLevel:    2,
		MinCost: MinCost{
			Level: 10,
			Cost:  0,
		},
		MaxCost: MaxCost{
			Level: 10,
			Cost:  15,
		},
		Exclude: []string{
			"depth_strider",
		},
		Category:     "armor_feet",
		Weight:       2,
		TreasureOnly: true,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	BindingCurse = Enchantment{
		ID:          10,
		DisplayName: "Curse of Binding",
		Name:        "binding_curse",
		MaxLevel:    1,
		MinCost: MinCost{
			Level: 0,
			Cost:  25,
		},
		MaxCost: MaxCost{
			Level: 0,
			Cost:  50,
		},
		Exclude:      []string{},
		Category:     "wearable",
		Weight:       1,
		TreasureOnly: true,
		Curse:        true,
		Tradeable:    true,
		Discoverable: true,
	}
	SoulSpeed = Enchantment{
		ID:          11,
		DisplayName: "Soul Speed",
		Name:        "soul_speed",
		MaxLevel:    3,
		MinCost: MinCost{
			Level: 10,
			Cost:  0,
		},
		MaxCost: MaxCost{
			Level: 10,
			Cost:  15,
		},
		Exclude:      []string{},
		Category:     "armor_feet",
		Weight:       1,
		TreasureOnly: true,
		Curse:        false,
		Tradeable:    false,
		Discoverable: false,
	}
	SwiftSneak = Enchantment{
		ID:          12,
		DisplayName: "Swift Sneak",
		Name:        "swift_sneak",
		MaxLevel:    3,
		MinCost: MinCost{
			Level: 25,
			Cost:  0,
		},
		MaxCost: MaxCost{
			Level: 25,
			Cost:  50,
		},
		Exclude:      []string{},
		Category:     "armor_legs",
		Weight:       1,
		TreasureOnly: true,
		Curse:        false,
		Tradeable:    false,
		Discoverable: false,
	}
	Sharpness = Enchantment{
		ID:          13,
		DisplayName: "Sharpness",
		Name:        "sharpness",
		MaxLevel:    5,
		MinCost: MinCost{
			Level: 11,
			Cost:  -10,
		},
		MaxCost: MaxCost{
			Level: 11,
			Cost:  10,
		},
		Exclude: []string{
			"smite",
			"bane_of_arthropods",
		},
		Category:     "weapon",
		Weight:       10,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Smite = Enchantment{
		ID:          14,
		DisplayName: "Smite",
		Name:        "smite",
		MaxLevel:    5,
		MinCost: MinCost{
			Level: 8,
			Cost:  -3,
		},
		MaxCost: MaxCost{
			Level: 8,
			Cost:  17,
		},
		Exclude: []string{
			"sharpness",
			"bane_of_arthropods",
		},
		Category:     "weapon",
		Weight:       5,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	BaneOfArthropods = Enchantment{
		ID:          15,
		DisplayName: "Bane of Arthropods",
		Name:        "bane_of_arthropods",
		MaxLevel:    5,
		MinCost: MinCost{
			Level: 8,
			Cost:  -3,
		},
		MaxCost: MaxCost{
			Level: 8,
			Cost:  17,
		},
		Exclude: []string{
			"sharpness",
			"smite",
		},
		Category:     "weapon",
		Weight:       5,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Knockback = Enchantment{
		ID:          16,
		DisplayName: "Knockback",
		Name:        "knockback",
		MaxLevel:    2,
		MinCost: MinCost{
			Level: 20,
			Cost:  -15,
		},
		MaxCost: MaxCost{
			Level: 10,
			Cost:  51,
		},
		Exclude:      []string{},
		Category:     "weapon",
		Weight:       5,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	FireAspect = Enchantment{
		ID:          17,
		DisplayName: "Fire Aspect",
		Name:        "fire_aspect",
		MaxLevel:    2,
		MinCost: MinCost{
			Level: 20,
			Cost:  -10,
		},
		MaxCost: MaxCost{
			Level: 10,
			Cost:  51,
		},
		Exclude:      []string{},
		Category:     "weapon",
		Weight:       2,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Looting = Enchantment{
		ID:          18,
		DisplayName: "Looting",
		Name:        "looting",
		MaxLevel:    3,
		MinCost: MinCost{
			Level: 9,
			Cost:  6,
		},
		MaxCost: MaxCost{
			Level: 10,
			Cost:  51,
		},
		Exclude: []string{
			"silk_touch",
		},
		Category:     "weapon",
		Weight:       2,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Sweeping = Enchantment{
		ID:          19,
		DisplayName: "Sweeping Edge",
		Name:        "sweeping",
		MaxLevel:    3,
		MinCost: MinCost{
			Level: 9,
			Cost:  -4,
		},
		MaxCost: MaxCost{
			Level: 9,
			Cost:  11,
		},
		Exclude:      []string{},
		Category:     "weapon",
		Weight:       2,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Efficiency = Enchantment{
		ID:          20,
		DisplayName: "Efficiency",
		Name:        "efficiency",
		MaxLevel:    5,
		MinCost: MinCost{
			Level: 10,
			Cost:  -9,
		},
		MaxCost: MaxCost{
			Level: 10,
			Cost:  51,
		},
		Exclude:      []string{},
		Category:     "digger",
		Weight:       10,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	SilkTouch = Enchantment{
		ID:          21,
		DisplayName: "Silk Touch",
		Name:        "silk_touch",
		MaxLevel:    1,
		MinCost: MinCost{
			Level: 0,
			Cost:  15,
		},
		MaxCost: MaxCost{
			Level: 10,
			Cost:  51,
		},
		Exclude: []string{
			"looting",
			"fortune",
			"luck_of_the_sea",
		},
		Category:     "digger",
		Weight:       1,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Unbreaking = Enchantment{
		ID:          22,
		DisplayName: "Unbreaking",
		Name:        "unbreaking",
		MaxLevel:    3,
		MinCost: MinCost{
			Level: 8,
			Cost:  -3,
		},
		MaxCost: MaxCost{
			Level: 10,
			Cost:  51,
		},
		Exclude:      []string{},
		Category:     "breakable",
		Weight:       5,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Fortune = Enchantment{
		ID:          23,
		DisplayName: "Fortune",
		Name:        "fortune",
		MaxLevel:    3,
		MinCost: MinCost{
			Level: 9,
			Cost:  6,
		},
		MaxCost: MaxCost{
			Level: 10,
			Cost:  51,
		},
		Exclude: []string{
			"silk_touch",
		},
		Category:     "digger",
		Weight:       2,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Power = Enchantment{
		ID:          24,
		DisplayName: "Power",
		Name:        "power",
		MaxLevel:    5,
		MinCost: MinCost{
			Level: 10,
			Cost:  -9,
		},
		MaxCost: MaxCost{
			Level: 10,
			Cost:  6,
		},
		Exclude:      []string{},
		Category:     "bow",
		Weight:       10,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Punch = Enchantment{
		ID:          25,
		DisplayName: "Punch",
		Name:        "punch",
		MaxLevel:    2,
		MinCost: MinCost{
			Level: 20,
			Cost:  -8,
		},
		MaxCost: MaxCost{
			Level: 20,
			Cost:  17,
		},
		Exclude:      []string{},
		Category:     "bow",
		Weight:       2,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Flame = Enchantment{
		ID:          26,
		DisplayName: "Flame",
		Name:        "flame",
		MaxLevel:    1,
		MinCost: MinCost{
			Level: 0,
			Cost:  20,
		},
		MaxCost: MaxCost{
			Level: 0,
			Cost:  50,
		},
		Exclude:      []string{},
		Category:     "bow",
		Weight:       2,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Infinity = Enchantment{
		ID:          27,
		DisplayName: "Infinity",
		Name:        "infinity",
		MaxLevel:    1,
		MinCost: MinCost{
			Level: 0,
			Cost:  20,
		},
		MaxCost: MaxCost{
			Level: 0,
			Cost:  50,
		},
		Exclude: []string{
			"mending",
		},
		Category:     "bow",
		Weight:       1,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	LuckOfTheSea = Enchantment{
		ID:          28,
		DisplayName: "Luck of the Sea",
		Name:        "luck_of_the_sea",
		MaxLevel:    3,
		MinCost: MinCost{
			Level: 9,
			Cost:  6,
		},
		MaxCost: MaxCost{
			Level: 10,
			Cost:  51,
		},
		Exclude: []string{
			"silk_touch",
		},
		Category:     "fishing_rod",
		Weight:       2,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Lure = Enchantment{
		ID:          29,
		DisplayName: "Lure",
		Name:        "lure",
		MaxLevel:    3,
		MinCost: MinCost{
			Level: 9,
			Cost:  6,
		},
		MaxCost: MaxCost{
			Level: 10,
			Cost:  51,
		},
		Exclude:      []string{},
		Category:     "fishing_rod",
		Weight:       2,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Loyalty = Enchantment{
		ID:          30,
		DisplayName: "Loyalty",
		Name:        "loyalty",
		MaxLevel:    3,
		MinCost: MinCost{
			Level: 7,
			Cost:  5,
		},
		MaxCost: MaxCost{
			Level: 0,
			Cost:  50,
		},
		Exclude: []string{
			"riptide",
		},
		Category:     "trident",
		Weight:       5,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Impaling = Enchantment{
		ID:          31,
		DisplayName: "Impaling",
		Name:        "impaling",
		MaxLevel:    5,
		MinCost: MinCost{
			Level: 8,
			Cost:  -7,
		},
		MaxCost: MaxCost{
			Level: 8,
			Cost:  13,
		},
		Exclude:      []string{},
		Category:     "trident",
		Weight:       2,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Riptide = Enchantment{
		ID:          32,
		DisplayName: "Riptide",
		Name:        "riptide",
		MaxLevel:    3,
		MinCost: MinCost{
			Level: 7,
			Cost:  10,
		},
		MaxCost: MaxCost{
			Level: 0,
			Cost:  50,
		},
		Exclude: []string{
			"loyalty",
			"channeling",
		},
		Category:     "trident",
		Weight:       2,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Channeling = Enchantment{
		ID:          33,
		DisplayName: "Channeling",
		Name:        "channeling",
		MaxLevel:    1,
		MinCost: MinCost{
			Level: 0,
			Cost:  25,
		},
		MaxCost: MaxCost{
			Level: 0,
			Cost:  50,
		},
		Exclude: []string{
			"riptide",
		},
		Category:     "trident",
		Weight:       1,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Multishot = Enchantment{
		ID:          34,
		DisplayName: "Multishot",
		Name:        "multishot",
		MaxLevel:    1,
		MinCost: MinCost{
			Level: 0,
			Cost:  20,
		},
		MaxCost: MaxCost{
			Level: 0,
			Cost:  50,
		},
		Exclude: []string{
			"piercing",
		},
		Category:     "crossbow",
		Weight:       2,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	QuickCharge = Enchantment{
		ID:          35,
		DisplayName: "Quick Charge",
		Name:        "quick_charge",
		MaxLevel:    3,
		MinCost: MinCost{
			Level: 20,
			Cost:  -8,
		},
		MaxCost: MaxCost{
			Level: 0,
			Cost:  50,
		},
		Exclude:      []string{},
		Category:     "crossbow",
		Weight:       5,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Piercing = Enchantment{
		ID:          36,
		DisplayName: "Piercing",
		Name:        "piercing",
		MaxLevel:    4,
		MinCost: MinCost{
			Level: 10,
			Cost:  -9,
		},
		MaxCost: MaxCost{
			Level: 0,
			Cost:  50,
		},
		Exclude: []string{
			"multishot",
		},
		Category:     "crossbow",
		Weight:       10,
		TreasureOnly: false,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	Mending = Enchantment{
		ID:          37,
		DisplayName: "Mending",
		Name:        "mending",
		MaxLevel:    1,
		MinCost: MinCost{
			Level: 25,
			Cost:  0,
		},
		MaxCost: MaxCost{
			Level: 25,
			Cost:  50,
		},
		Exclude: []string{
			"infinity",
		},
		Category:     "breakable",
		Weight:       2,
		TreasureOnly: true,
		Curse:        false,
		Tradeable:    true,
		Discoverable: true,
	}
	VanishingCurse = Enchantment{
		ID:          38,
		DisplayName: "Curse of Vanishing",
		Name:        "vanishing_curse",
		MaxLevel:    1,
		MinCost: MinCost{
			Level: 0,
			Cost:  25,
		},
		MaxCost: MaxCost{
			Level: 0,
			Cost:  50,
		},
		Exclude:      []string{},
		Category:     "vanishable",
		Weight:       1,
		TreasureOnly: true,
		Curse:        true,
		Tradeable:    true,
		Discoverable: true,
	}
)

// ByID is an index of minecraft enchantments by their ID.
var ByID = map[ID]*Enchantment{
	0:  &Protection,
	1:  &FireProtection,
	2:  &FeatherFalling,
	3:  &BlastProtection,
	4:  &ProjectileProtection,
	5:  &Respiration,
	6:  &AquaAffinity,
	7:  &Thorns,
	8:  &DepthStrider,
	9:  &FrostWalker,
	10: &BindingCurse,
	11: &SoulSpeed,
	12: &SwiftSneak,
	13: &Sharpness,
	14: &Smite,
	15: &BaneOfArthropods,
	16: &Knockback,
	17: &FireAspect,
	18: &Looting,
	19: &Sweeping,
	20: &Efficiency,
	21: &SilkTouch,
	22: &Unbreaking,
	23: &Fortune,
	24: &Power,
	25: &Punch,
	26: &Flame,
	27: &Infinity,
	28: &LuckOfTheSea,
	29: &Lure,
	30: &Loyalty,
	31: &Impaling,
	32: &Riptide,
	33: &Channeling,
	34: &Multishot,
	35: &QuickCharge,
	36: &Piercing,
	37: &Mending,
	38: &VanishingCurse,
}

// ByName is an index of minecraft enchantments by their name.
var ByName = map[string]*Enchantment{
	"protection":            &Protection,
	"fire_protection":       &FireProtection,
	"feather_falling":       &FeatherFalling,
	"blast_protection":      &BlastProtection,
	"projectile_protection": &ProjectileProtection,
	"respiration":           &Respiration,
	"aqua_affinity":         &AquaAffinity,
	"thorns":                &Thorns,
	"depth_strider":         &DepthStrider,
	"frost_walker":          &FrostWalker,
	"binding_curse":         &BindingCurse,
	"soul_speed":            &SoulSpeed,
	"swift_sneak":           &SwiftSneak,
	"sharpness":             &Sharpness,
	"smite":                 &Smite,
	"bane_of_arthropods":    &BaneOfArthropods,
	"knockback":             &Knockback,
	"fire_aspect":           &FireAspect,
	"looting":               &Looting,
	"sweeping":              &Sweeping,
	"efficiency":            &Efficiency,
	"silk_touch":            &SilkTouch,
	"unbreaking":            &Unbreaking,
	"fortune":               &Fortune,
	"power":                 &Power,
	"punch":                 &Punch,
	"flame":                 &Flame,
	"infinity":              &Infinity,
	"luck_of_the_sea":       &LuckOfTheSea,
	"lure":                  &Lure,
	"loyalty":               &Loyalty,
	"impaling":              &Impaling,
	"riptide":               &Riptide,
	"channeling":            &Channeling,
	"multishot":             &Multishot,
	"quick_charge":          &QuickCharge,
	"piercing":              &Piercing,
	"mending":               &Mending,
	"vanishing_curse":       &VanishingCurse,
}

// ByDisplayName is an index of minecraft enchantments by their display name.
var ByDisplayName = map[string]*Enchantment{
	"Protection":            &Protection,
	"Fire Protection":       &FireProtection,
	"Feather Falling":       &FeatherFalling,
	"Blast Protection":      &BlastProtection,
	"Projectile Protection": &ProjectileProtection,
	"Respiration":           &Respiration,
	"Aqua Affinity":         &AquaAffinity,
	"Thorns":                &Thorns,
	"Depth Strider":         &DepthStrider,
	"Frost Walker":          &FrostWalker,
	"Curse of Binding":      &BindingCurse,
	"Soul Speed":            &SoulSpeed,
	"Swift Sneak":           &SwiftSneak,
	"Sharpness":             &Sharpness,
	"Smite":                 &Smite,
	"Bane of Arthropods":    &BaneOfArthropods,
	"Knockback":             &Knockback,
	"Fire Aspect":           &FireAspect,
	"Looting":               &Looting,
	"Sweeping Edge":         &Sweeping,
	"Efficiency":            &Efficiency,
	"Silk Touch":            &SilkTouch,
	"Unbreaking":            &Unbreaking,
	"Fortune":               &Fortune,
	"Power":                 &Power,
	"Punch":                 &Punch,
	"Flame":                 &Flame,
	"Infinity":              &Infinity,
	"Luck of the Sea":       &LuckOfTheSea,
	"Lure":                  &Lure,
	"Loyalty":               &Loyalty,
	"Impaling":              &Impaling,
	"Riptide":               &Riptide,
	"Channeling":            &Channeling,
	"Multishot":             &Multishot,
	"Quick Charge":          &QuickCharge,
	"Piercing":              &Piercing,
	"Mending":               &Mending,
	"Curse of Vanishing":    &VanishingCurse,
}