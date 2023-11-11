// Code generated by "stringer -type=StairsShape -output=stairs_shape_string.go"; DO NOT EDIT.

package properties

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StairsShapeStraight-0]
	_ = x[StairsShapeInnerLeft-1]
	_ = x[StairsShapeInnerRight-2]
	_ = x[StairsShapeOuterLeft-3]
	_ = x[StairsShapeOuterRight-4]
}

const _StairsShape_name = "StairsShapeStraightStairsShapeInnerLeftStairsShapeInnerRightStairsShapeOuterLeftStairsShapeOuterRight"

var _StairsShape_index = [...]uint8{0, 19, 39, 60, 80, 101}

func (i StairsShape) String() string {
	if i >= StairsShape(len(_StairsShape_index)-1) {
		return "StairsShape(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StairsShape_name[_StairsShape_index[i]:_StairsShape_index[i+1]]
}