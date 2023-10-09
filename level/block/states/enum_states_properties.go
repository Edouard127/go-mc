package states

var (
	HorizontalAxisPropertyProperty = NewEnumProperty("axis", map[string]Axis{
		"x": X,
		"z": Z,
	})
	AxisProperty = NewEnumProperty("axis", map[string]Axis{
		"x": X,
		"y": Y,
		"z": Z,
	})
	FacingProperty = NewEnumProperty("facing", map[string]Direction{
		"north": DirectionNorth,
		"east":  DirectionEast,
		"south": DirectionSouth,
		"west":  DirectionWest,
	})
	HorizontalFacingProperty = NewEnumProperty("facing", map[string]Direction{
		"north": DirectionNorth,
		"east":  DirectionEast,
		"south": DirectionSouth,
		"west":  DirectionWest,
	})
	OrientationProperty = NewEnumProperty("orientation", map[string]FrontAndTop{
		"down_east":  DownEast,
		"down_north": DownNorth,
		"down_south": DownSouth,
		"down_west":  DownWest,
		"up_east":    UpEast,
		"up_north":   UpNorth,
		"up_south":   UpSouth,
		"up_west":    UpWest,
		"west_up":    WestUp,
		"east_up":    EastUp,
		"north_up":   NorthUp,
		"south_up":   SouthUp,
	})
	AttachFaceProperty = NewEnumProperty("face", map[string]AttachFace{
		"floor":   AttachFaceFloor,
		"wall":    AttachFaceWall,
		"ceiling": AttachFaceCeiling,
	})
	BellAttachmentProperty = NewEnumProperty("attachment", map[string]BellAttachType{
		"floor":       BellAttachTypeFloor,
		"ceiling":     BellAttachTypeCeiling,
		"single_wall": BellAttachTypeSingleWall,
		"double_wall": BellAttachTypeDoubleWall,
	})
	EastWallProperty = NewEnumProperty("east", map[string]WallSide{
		"none": WallSideNone,
		"low":  WallSideLow,
		"tall": WallSideTall,
	})
	NorthWallProperty = NewEnumProperty("north", map[string]WallSide{
		"none": WallSideNone,
		"low":  WallSideLow,
		"tall": WallSideTall,
	})
	SouthWallProperty = NewEnumProperty("south", map[string]WallSide{
		"none": WallSideNone,
		"low":  WallSideLow,
		"tall": WallSideTall,
	})
	WestWallProperty = NewEnumProperty("west", map[string]WallSide{
		"none": WallSideNone,
		"low":  WallSideLow,
		"tall": WallSideTall,
	})
	EastRedstoneProperty = NewEnumProperty("east", map[string]RedstoneSide{
		"none": RedstoneSideNone,
		"side": RedstoneSideSide,
		"up":   RedstoneSideUp,
	})
	NorthRedstoneProperty = NewEnumProperty("north", map[string]RedstoneSide{
		"none": RedstoneSideNone,
		"side": RedstoneSideSide,
		"up":   RedstoneSideUp,
	})
	SouthRedstoneProperty = NewEnumProperty("south", map[string]RedstoneSide{
		"none": RedstoneSideNone,
		"side": RedstoneSideSide,
		"up":   RedstoneSideUp,
	})
	WestRedstoneProperty = NewEnumProperty("west", map[string]RedstoneSide{
		"none": RedstoneSideNone,
		"side": RedstoneSideSide,
		"up":   RedstoneSideUp,
	})
	DoubleBlockHalfProperty = NewEnumProperty("half", map[string]DoubleBlockHalf{
		"lower": DoubleBlockHalfLower,
		"upper": DoubleBlockHalfUpper,
	})
	HalfProperty = NewEnumProperty("half", map[string]Half{
		"top":    HalfTop,
		"bottom": HalfBottom,
	})
	RailShapeProperty = NewEnumProperty("shape", map[string]RailShape{
		"north_south":     RailShapeNorthSouth,
		"east_west":       RailShapeEastWest,
		"ascending_east":  RailShapeAscendingEast,
		"ascending_west":  RailShapeAscendingWest,
		"ascending_north": RailShapeAscendingNorth,
		"ascending_south": RailShapeAscendingSouth,
		"south_east":      RailShapeSouthEast,
		"south_west":      RailShapeSouthWest,
		"north_west":      RailShapeNorthWest,
		"north_east":      RailShapeNorthEast,
	})
	RailShapeStraightProperty = NewEnumProperty("shape", map[string]RailShape{
		"north_south":     RailShapeNorthSouth,
		"east_west":       RailShapeEastWest,
		"ascending_east":  RailShapeAscendingEast,
		"ascending_west":  RailShapeAscendingWest,
		"ascending_north": RailShapeAscendingNorth,
		"ascending_south": RailShapeAscendingSouth,
	})
	/*LevelCauldronProperty        = block.NewIntegerProperty("level", 1, 3)
	LevelComposterProperty       = block.NewIntegerProperty("level", 0, 8)
	LevelFlowingProperty         = block.NewIntegerProperty("level", 1, 8)
	LevelHoneyProperty           = block.NewIntegerProperty("level", 0, 5)
	RedstoneSignalProperty       = block.NewIntegerProperty("signal", 0, 15)
	StagesProperty               = block.NewIntegerProperty("stage", 0, 1)
	StabilityProperty            = block.NewIntegerProperty("distance", 0, 2)
	RespawnAnchorChargesProperty = block.NewIntegerProperty("charges", 0, 4)
	Rotation16Property           = block.NewIntegerProperty("rotation", 0, 15)*/
	BedPartProperty = NewEnumProperty("part", map[string]BedPart{
		"head": BedPartHead,
		"foot": BedPartFoot,
	})
	ChestTypeProperty = NewEnumProperty("type", map[string]ChestType{
		"single": ChestTypeSingle,
		"left":   ChestTypeLeft,
		"right":  ChestTypeRight,
	})
	ComparatorModeProperty = NewEnumProperty("mode", map[string]ComparatorMode{
		"compare":  ComparatorModeCompare,
		"subtract": ComparatorModeSubtract,
	})
	DoorHingeProperty = NewEnumProperty("hinge", map[string]DoorHingeSide{
		"left":  DoorHingeSideLeft,
		"right": DoorHingeSideRight,
	})
	InstrumentProperty = NewEnumProperty("instrument", map[string]NoteBlockInstrument{
		"harp":           NoteBlockInstrumentHarp,
		"basedrum":       NoteBlockInstrumentBasedrum,
		"snare":          NoteBlockInstrumentSnare,
		"hat":            NoteBlockInstrumentHat,
		"bass":           NoteBlockInstrumentBass,
		"flute":          NoteBlockInstrumentFlute,
		"bell":           NoteBlockInstrumentBell,
		"guitar":         NoteBlockInstrumentGuitar,
		"chime":          NoteBlockInstrumentChime,
		"xylophone":      NoteBlockInstrumentXylophone,
		"iron_xylophone": NoteBlockInstrumentIronXylophone,
		"cow_bell":       NoteBlockInstrumentCowBell,
		"didgeridoo":     NoteBlockInstrumentDidgeridoo,
		"bit":            NoteBlockInstrumentBit,
		"banjo":          NoteBlockInstrumentBanjo,
		"pling":          NoteBlockInstrumentPling,
	})
	PistonTypeProperty = NewEnumProperty("type", map[string]PistonType{
		"normal": PistonTypeDefault,
		"sticky": PistonTypeSticky,
	})
	SlabTypeProperty = NewEnumProperty("type", map[string]SlabType{
		"bottom": SlabTypeBottom,
		"top":    SlabTypeTop,
		"double": SlabTypeDouble,
	})
	StairsShapeProperty = NewEnumProperty("shape", map[string]StairsShape{
		"straight":    StairsShapeStraight,
		"inner_left":  StairsShapeInnerLeft,
		"inner_right": StairsShapeInnerRight,
		"outer_left":  StairsShapeOuterLeft,
		"outer_right": StairsShapeOuterRight,
	})
	StructureModeProperty = NewEnumProperty("mode", map[string]StructureMode{
		"save":   StructureModeSave,
		"load":   StructureModeLoad,
		"corner": StructureModeCorner,
		"data":   StructureModeData,
	})
	BambooLeavesProperty = NewEnumProperty("leaves", map[string]BambooLeaves{
		"none":  BambooLeavesNone,
		"small": BambooLeavesSmall,
		"large": BambooLeavesLarge,
	})
	TiltProperty = NewEnumProperty("tilt", map[string]Tilt{
		"none":     TiltNone,
		"unstable": TiltUnstable,
		"partial":  TiltPartial,
		"full":     TiltFull,
	})
	VerticalDirectionProperty = NewEnumProperty("direction", map[string]Direction{
		"up":   DirectionUp,
		"down": DirectionDown,
	})
	DripstoneThicknessProperty = NewEnumProperty("thickness", map[string]DripstoneThickness{
		"tip_merge": DripstoneThicknessTipMerge,
		"tip":       DripstoneThicknessTip,
		"frustum":   DripstoneThicknessFrustum,
		"middle":    DripstoneThicknessMiddle,
		"base":      DripstoneThicknessBase,
	})
	SculkSensorPhaseProperty = NewEnumProperty("phase", map[string]SculkSensorPhase{
		"inactive": SculkSensorPhaseInactive,
		"active":   SculkSensorPhaseActive,
		"cooldown": SculkSensorPhaseCooldown,
	})
)
