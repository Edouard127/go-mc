package states

import "github.com/Edouard127/go-mc/level/block/states/properties"

var (
	Attached       = NewBooleanProperty("attached")
	Bottom         = NewBooleanProperty("bottom")
	Conditional    = NewBooleanProperty("conditional")
	Disarmed       = NewBooleanProperty("disarmed")
	Drag           = NewBooleanProperty("drag")
	Enabled        = NewBooleanProperty("enabled")
	Extended       = NewBooleanProperty("extended")
	Eye            = NewBooleanProperty("eye")
	Falling        = NewBooleanProperty("falling")
	Hanging        = NewBooleanProperty("hanging")
	HasBottle0     = NewBooleanProperty("has_bottle_0")
	HasBottle1     = NewBooleanProperty("has_bottle_1")
	HasBottle2     = NewBooleanProperty("has_bottle_2")
	HasRecord      = NewBooleanProperty("has_record")
	HasBook        = NewBooleanProperty("has_book")
	Inverted       = NewBooleanProperty("inverted")
	InWall         = NewBooleanProperty("in_wall")
	Lit            = NewBooleanProperty("lit")
	Locked         = NewBooleanProperty("locked")
	Occupied       = NewBooleanProperty("occupied")
	Open           = NewBooleanProperty("open")
	Persistent     = NewBooleanProperty("persistent")
	Powered        = NewBooleanProperty("powered")
	Short          = NewBooleanProperty("short")
	SignalFire     = NewBooleanProperty("signal_fire")
	Snowy          = NewBooleanProperty("snowy")
	Triggered      = NewBooleanProperty("triggered")
	Unstable       = NewBooleanProperty("unstable")
	Waterlogged    = NewBooleanProperty("waterlogged")
	Berries        = NewBooleanProperty("berries")
	Bloom          = NewBooleanProperty("bloom")
	Shrieking      = NewBooleanProperty("shrieking")
	CanSummon      = NewBooleanProperty("can_summon")
	HorizontalAxis = NewEnumProperty("horizontal_axis", map[string]properties.Axis{
		"x": properties.X,
		"z": properties.Z,
	})
	Axis = NewEnumProperty("axis", map[string]properties.Axis{
		"x": properties.X,
		"y": properties.Y,
		"z": properties.Z,
	})
	Up     = NewBooleanProperty("up")
	Down   = NewBooleanProperty("down")
	North  = NewBooleanProperty("north")
	East   = NewBooleanProperty("east")
	South  = NewBooleanProperty("south")
	West   = NewBooleanProperty("west")
	Facing = NewEnumProperty("facing", map[string]properties.Direction{
		"north": properties.North,
		"east":  properties.East,
		"south": properties.South,
		"west":  properties.West,
		"up":    properties.Up,
		"down":  properties.Down,
	})
	FacingHopper = NewEnumProperty("facing_hopper", map[string]properties.Direction{
		"north": properties.North,
		"east":  properties.East,
		"south": properties.South,
		"west":  properties.West,
		"down":  properties.Down,
	})
	HorizontalFacing = NewEnumProperty("horizontal_facing", map[string]properties.Direction{
		"north": properties.North,
		"east":  properties.East,
		"south": properties.South,
		"west":  properties.West,
	})
	FlowerAmount = NewIntegerProperty("flower_amount", 1, 4)
	Orientation  = NewEnumProperty("orientation", map[string]properties.FrontAndTop{
		"down_east":  properties.DownEast,
		"down_north": properties.DownNorth,
		"down_south": properties.DownSouth,
		"down_west":  properties.DownWest,
		"up_east":    properties.UpEast,
		"up_north":   properties.UpNorth,
		"up_south":   properties.UpSouth,
		"up_west":    properties.UpWest,
		"west_up":    properties.WestUp,
		"east_up":    properties.EastUp,
		"north_up":   properties.NorthUp,
		"south_up":   properties.SouthUp,
	})
	AttachFace = NewEnumProperty("attach_face", map[string]properties.AttachFace{
		"floor":   properties.AttachFaceFloor,
		"wall":    properties.AttachFaceWall,
		"ceiling": properties.AttachFaceCeiling,
	})
	BellAttachment = NewEnumProperty("bell_attachment", map[string]properties.BellAttachType{
		"floor":       properties.BellAttachTypeFloor,
		"ceiling":     properties.BellAttachTypeCeiling,
		"single_wall": properties.BellAttachTypeSingleWall,
		"double_wall": properties.BellAttachTypeDoubleWall,
	})
	EastWall = NewEnumProperty("east_wall", map[string]properties.WallSide{
		"none": properties.WallSideNone,
		"low":  properties.WallSideLow,
		"tall": properties.WallSideTall,
	})
	NorthWall = NewEnumProperty("north_wall", map[string]properties.WallSide{
		"none": properties.WallSideNone,
		"low":  properties.WallSideLow,
		"tall": properties.WallSideTall,
	})
	SouthWall = NewEnumProperty("south_wall", map[string]properties.WallSide{
		"none": properties.WallSideNone,
		"low":  properties.WallSideLow,
		"tall": properties.WallSideTall,
	})
	WestWall = NewEnumProperty("west_wall", map[string]properties.WallSide{
		"none": properties.WallSideNone,
		"low":  properties.WallSideLow,
		"tall": properties.WallSideTall,
	})
	EastRedstone = NewEnumProperty("east_redstone", map[string]properties.RedstoneSide{
		"up":   properties.RedstoneSideUp,
		"side": properties.RedstoneSideSide,
		"none": properties.RedstoneSideNone,
	})
	NorthRedstone = NewEnumProperty("north_redstone", map[string]properties.RedstoneSide{
		"up":   properties.RedstoneSideUp,
		"side": properties.RedstoneSideSide,
		"none": properties.RedstoneSideNone,
	})
	SouthRedstone = NewEnumProperty("south_redstone", map[string]properties.RedstoneSide{
		"up":   properties.RedstoneSideUp,
		"side": properties.RedstoneSideSide,
		"none": properties.RedstoneSideNone,
	})
	WestRedstone = NewEnumProperty("west_redstone", map[string]properties.RedstoneSide{
		"up":   properties.RedstoneSideUp,
		"side": properties.RedstoneSideSide,
		"none": properties.RedstoneSideNone,
	})
	DoubleBlockHalf = NewEnumProperty("double_block_half", map[string]properties.DoubleBlockHalf{
		"upper": properties.DoubleBlockHalfUpper,
		"lower": properties.DoubleBlockHalfLower,
	})
	Half = NewEnumProperty("half", map[string]properties.Half{
		"top":    properties.HalfTop,
		"bottom": properties.HalfBottom,
	})
	RailShape = NewEnumProperty("rail_shape", map[string]properties.RailShape{
		"north_south":     properties.RailShapeNorthSouth,
		"east_west":       properties.RailShapeEastWest,
		"ascending_east":  properties.RailShapeAscendingEast,
		"ascending_west":  properties.RailShapeAscendingWest,
		"ascending_north": properties.RailShapeAscendingNorth,
		"ascending_south": properties.RailShapeAscendingSouth,
		"south_east":      properties.RailShapeSouthEast,
		"south_west":      properties.RailShapeSouthWest,
		"north_west":      properties.RailShapeNorthWest,
		"north_east":      properties.RailShapeNorthEast,
	})
	RailShapeStraight = NewEnumProperty("rail_shape_straight", map[string]properties.RailShape{
		"north_south": properties.RailShapeNorthSouth,
		"east_west":   properties.RailShapeEastWest,
	})
	Age1                 = NewIntegerProperty("age_1", 0, 1)
	Age2                 = NewIntegerProperty("age_2", 0, 2)
	Age3                 = NewIntegerProperty("age_3", 0, 3)
	Age4                 = NewIntegerProperty("age_4", 0, 4)
	Age5                 = NewIntegerProperty("age_5", 0, 5)
	Age7                 = NewIntegerProperty("age_7", 0, 7)
	Age15                = NewIntegerProperty("age_15", 0, 15)
	Age25                = NewIntegerProperty("age_25", 0, 25)
	Bites                = NewIntegerProperty("bites", 0, 6)
	Candles              = NewIntegerProperty("candles", 1, 4)
	Delay                = NewIntegerProperty("delay", 1, 4)
	Distance             = NewIntegerProperty("distance", 1, 7)
	Eggs                 = NewIntegerProperty("eggs", 1, 4)
	Hatch                = NewIntegerProperty("hatch", 0, 2)
	Layers               = NewIntegerProperty("layers", 1, 8)
	LevelCauldron        = NewIntegerProperty("level_cauldron", 0, 3)
	LevelComposter       = NewIntegerProperty("level_composter", 0, 8)
	LevelFlowing         = NewIntegerProperty("level_flowing", 1, 8)
	LevelHoney           = NewIntegerProperty("level_honey", 0, 5)
	Level                = NewIntegerProperty("level", 0, 15)
	Moisture             = NewIntegerProperty("moisture", 0, 7)
	Note                 = NewIntegerProperty("note", 0, 24)
	Pickles              = NewIntegerProperty("pickles", 1, 4)
	Power                = NewIntegerProperty("power", 0, 15)
	Stage                = NewIntegerProperty("stage", 0, 1)
	StabilityDistance    = NewIntegerProperty("stability_distance", 0, 7)
	RespawnAnchorCharges = NewIntegerProperty("respawn_anchor_charges", 0, 4)
	Rotation16           = NewIntegerProperty("rotation_16", 0, 15)
	BedPart              = NewEnumProperty("bed_part", map[string]properties.BedPart{
		"head": properties.BedPartHead,
		"foot": properties.BedPartFoot,
	})
	ChestType = NewEnumProperty("chest_type", map[string]properties.ChestType{
		"single": properties.ChestTypeSingle,
		"left":   properties.ChestTypeLeft,
		"right":  properties.ChestTypeRight,
	})
	ModeComparator = NewEnumProperty("mode_comparator", map[string]properties.ComparatorMode{
		"compare":  properties.ComparatorModeCompare,
		"subtract": properties.ComparatorModeSubtract,
	})
	DoorHinge = NewEnumProperty("door_hinge", map[string]properties.DoorHinge{
		"left":  properties.DoorHingeLeft,
		"right": properties.DoorHingeRight,
	})
	NoteblockInstrument = NewEnumProperty("noteblock_instrument", map[string]properties.NoteBlockInstrument{
		"harp":           properties.NoteBlockInstrumentHarp,
		"basedrum":       properties.NoteBlockInstrumentBasedrum,
		"snare":          properties.NoteBlockInstrumentSnare,
		"hat":            properties.NoteBlockInstrumentHat,
		"bass":           properties.NoteBlockInstrumentBass,
		"flute":          properties.NoteBlockInstrumentFlute,
		"bell":           properties.NoteBlockInstrumentBell,
		"guitar":         properties.NoteBlockInstrumentGuitar,
		"chime":          properties.NoteBlockInstrumentChime,
		"xylophone":      properties.NoteBlockInstrumentXylophone,
		"iron_xylophone": properties.NoteBlockInstrumentIronXylophone,
		"cow_bell":       properties.NoteBlockInstrumentCowBell,
		"didgeridoo":     properties.NoteBlockInstrumentDidgeridoo,
		"bit":            properties.NoteBlockInstrumentBit,
		"banjo":          properties.NoteBlockInstrumentBanjo,
		"pling":          properties.NoteBlockInstrumentPling,
	})
	PistonType = NewEnumProperty("piston_type", map[string]properties.PistonType{
		"normal": properties.PistonTypeDefault,
		"sticky": properties.PistonTypeSticky,
	})
	SlabType = NewEnumProperty("slab_type", map[string]properties.SlabType{
		"bottom": properties.SlabTypeBottom,
		"top":    properties.SlabTypeTop,
		"double": properties.SlabTypeDouble,
	})
	StairsShape = NewEnumProperty("stairs_shape", map[string]properties.StairsShape{
		"straight":    properties.StairsShapeStraight,
		"inner_left":  properties.StairsShapeInnerLeft,
		"inner_right": properties.StairsShapeInnerRight,
		"outer_left":  properties.StairsShapeOuterLeft,
		"outer_right": properties.StairsShapeOuterRight,
	})
	StructureblockMode = NewEnumProperty("structureblock_mode", map[string]properties.StructureMode{
		"save":   properties.StructureModeSave,
		"load":   properties.StructureModeLoad,
		"corner": properties.StructureModeCorner,
		"data":   properties.StructureModeData,
	})
	BambooLeaves = NewEnumProperty("bamboo_leaves", map[string]properties.BambooLeaves{
		"none":  properties.BambooLeavesNone,
		"small": properties.BambooLeavesSmall,
		"large": properties.BambooLeavesLarge,
	})
	Tilt = NewEnumProperty("tilt", map[string]properties.Tilt{
		"unstable": properties.TiltUnstable,
		"partial":  properties.TiltPartial,
		"full":     properties.TiltFull,
	})
	VerticalDirection = NewEnumProperty("vertical_direction", map[string]properties.Direction{
		"up":   properties.Up,
		"down": properties.Down,
	})
	DripstoneThickness = NewEnumProperty("dripstone_thickness", map[string]properties.DripstoneThickness{
		"tip":    properties.DripstoneThicknessTip,
		"middle": properties.DripstoneThicknessMiddle,
		"base":   properties.DripstoneThicknessBase,
	})
	SculkSensorPhase = NewEnumProperty("sculk_sensor_phase", map[string]properties.SculkSensorPhase{
		"cooldown": properties.SculkSensorPhaseCooldown,
		"active":   properties.SculkSensorPhaseActive,
	})
	ChiseledBookshelfSlot0Occupied = NewBooleanProperty("chiseled_bookshelf_slot_0_occupied")
	ChiseledBookshelfSlot1Occupied = NewBooleanProperty("chiseled_bookshelf_slot_1_occupied")
	ChiseledBookshelfSlot2Occupied = NewBooleanProperty("chiseled_bookshelf_slot_2_occupied")
	ChiseledBookshelfSlot3Occupied = NewBooleanProperty("chiseled_bookshelf_slot_3_occupied")
	ChiseledBookshelfSlot4Occupied = NewBooleanProperty("chiseled_bookshelf_slot_4_occupied")
	ChiseledBookshelfSlot5Occupied = NewBooleanProperty("chiseled_bookshelf_slot_5_occupied")
	Dusted                         = NewIntegerProperty("dusted", 0, 3)
	Cracked                        = NewBooleanProperty("cracked")
)

var FromName = map[string]Property{
	"attached":                           Attached,
	"bottom":                             Bottom,
	"conditional":                        Conditional,
	"disarmed":                           Disarmed,
	"drag":                               Drag,
	"enabled":                            Enabled,
	"extended":                           Extended,
	"eye":                                Eye,
	"falling":                            Falling,
	"hanging":                            Hanging,
	"has_bottle_0":                       HasBottle0,
	"has_bottle_1":                       HasBottle1,
	"has_bottle_2":                       HasBottle2,
	"has_record":                         HasRecord,
	"has_book":                           HasBook,
	"inverted":                           Inverted,
	"in_wall":                            InWall,
	"lit":                                Lit,
	"locked":                             Locked,
	"occupied":                           Occupied,
	"open":                               Open,
	"persistent":                         Persistent,
	"powered":                            Powered,
	"short":                              Short,
	"signal_fire":                        SignalFire,
	"snowy":                              Snowy,
	"triggered":                          Triggered,
	"unstable":                           Unstable,
	"waterlogged":                        Waterlogged,
	"berries":                            Berries,
	"bloom":                              Bloom,
	"shrieking":                          Shrieking,
	"can_summon":                         CanSummon,
	"horizontal_axis":                    HorizontalAxis,
	"axis":                               Axis,
	"up":                                 Up,
	"down":                               Down,
	"north":                              North,
	"east":                               East,
	"south":                              South,
	"west":                               West,
	"facing":                             Facing,
	"facing_hopper":                      FacingHopper,
	"horizontal_facing":                  HorizontalFacing,
	"flower_amount":                      FlowerAmount,
	"orientation":                        Orientation,
	"attach_face":                        AttachFace,
	"bell_attachment":                    BellAttachment,
	"east_wall":                          EastWall,
	"north_wall":                         NorthWall,
	"south_wall":                         SouthWall,
	"west_wall":                          WestWall,
	"east_redstone":                      EastRedstone,
	"north_redstone":                     NorthRedstone,
	"south_redstone":                     SouthRedstone,
	"west_redstone":                      WestRedstone,
	"double_block_half":                  DoubleBlockHalf,
	"half":                               Half,
	"rail_shape":                         RailShape,
	"rail_shape_straight":                RailShapeStraight,
	"age_1":                              Age1,
	"age_2":                              Age2,
	"age_3":                              Age3,
	"age_4":                              Age4,
	"age_5":                              Age5,
	"age_7":                              Age7,
	"age_15":                             Age15,
	"age_25":                             Age25,
	"bites":                              Bites,
	"candles":                            Candles,
	"delay":                              Delay,
	"distance":                           Distance,
	"eggs":                               Eggs,
	"hatch":                              Hatch,
	"layers":                             Layers,
	"level_cauldron":                     LevelCauldron,
	"level_composter":                    LevelComposter,
	"level_flowing":                      LevelFlowing,
	"level_honey":                        LevelHoney,
	"level":                              Level,
	"moisture":                           Moisture,
	"note":                               Note,
	"pickles":                            Pickles,
	"power":                              Power,
	"stage":                              Stage,
	"stability_distance":                 StabilityDistance,
	"respawn_anchor_charges":             RespawnAnchorCharges,
	"rotation_16":                        Rotation16,
	"bed_part":                           BedPart,
	"chest_type":                         ChestType,
	"mode_comparator":                    ModeComparator,
	"door_hinge":                         DoorHinge,
	"noteblock_instrument":               NoteblockInstrument,
	"piston_type":                        PistonType,
	"slab_type":                          SlabType,
	"stairs_shape":                       StairsShape,
	"structureblock_mode":                StructureblockMode,
	"bamboo_leaves":                      BambooLeaves,
	"tilt":                               Tilt,
	"vertical_direction":                 VerticalDirection,
	"dripstone_thickness":                DripstoneThickness,
	"sculk_sensor_phase":                 SculkSensorPhase,
	"chiseled_bookshelf_slot_0_occupied": ChiseledBookshelfSlot0Occupied,
	"chiseled_bookshelf_slot_1_occupied": ChiseledBookshelfSlot1Occupied,
	"chiseled_bookshelf_slot_2_occupied": ChiseledBookshelfSlot2Occupied,
	"chiseled_bookshelf_slot_3_occupied": ChiseledBookshelfSlot3Occupied,
	"chiseled_bookshelf_slot_4_occupied": ChiseledBookshelfSlot4Occupied,
	"chiseled_bookshelf_slot_5_occupied": ChiseledBookshelfSlot5Occupied,
	"dusted":                             Dusted,
	"cracked":                            Cracked,
}
