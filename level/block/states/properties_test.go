package states

import (
	"fmt"
	"testing"
)

var (
	chestType = NewEnumProperty[ChestType]("type", map[string]ChestType{
		"single": ChestTypeSingle,
		"left":   ChestTypeLeft,
		"right":  ChestTypeRight,
	})
	moisture          = NewIntegerProperty("moisture", 0, 7)
	stabilityDistance = NewIntegerProperty("distance", 0, 7)
	stairsShape       = NewEnumProperty[StairsShape]("shape", map[string]StairsShape{
		"straight":    StairsShapeStraight,
		"inner_left":  StairsShapeInnerLeft,
		"inner_right": StairsShapeInnerRight,
		"outer_left":  StairsShapeOuterLeft,
		"outer_right": StairsShapeOuterRight,
	})
)

func TestNewPropertyEnum(t *testing.T) {
	fmt.Println(chestType, moisture, stabilityDistance, stairsShape)
}
