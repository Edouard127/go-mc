package states

import (
	"fmt"
	"github.com/Edouard127/go-mc/level/block/states/properties"
	"testing"
)

var (
	chestType = NewEnumProperty("type", map[string]properties.ChestType{
		"single": properties.ChestTypeSingle,
		"left":   properties.ChestTypeLeft,
		"right":  properties.ChestTypeRight,
	})
	moisture          = NewIntegerProperty("moisture", 0, 7)
	stabilityDistance = NewIntegerProperty("distance", 0, 7)
	stairsShape       = NewEnumProperty("shape", map[string]properties.StairsShape{
		"straight":    properties.StairsShapeStraight,
		"inner_left":  properties.StairsShapeInnerLeft,
		"inner_right": properties.StairsShapeInnerRight,
		"outer_left":  properties.StairsShapeOuterLeft,
		"outer_right": properties.StairsShapeOuterRight,
	})
)

func TestNewPropertyEnum(t *testing.T) {
	fmt.Println(chestType, moisture, stabilityDistance, stairsShape)
}
