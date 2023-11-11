// Code generated by "stringer -type=BellAttachType -output=bell_attach_type_string.go"; DO NOT EDIT.

package properties

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BellAttachTypeFloor-0]
	_ = x[BellAttachTypeCeiling-1]
	_ = x[BellAttachTypeSingleWall-2]
	_ = x[BellAttachTypeDoubleWall-3]
}

const _BellAttachType_name = "BellAttachTypeFloorBellAttachTypeCeilingBellAttachTypeSingleWallBellAttachTypeDoubleWall"

var _BellAttachType_index = [...]uint8{0, 19, 40, 64, 88}

func (i BellAttachType) String() string {
	if i >= BellAttachType(len(_BellAttachType_index)-1) {
		return "BellAttachType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BellAttachType_name[_BellAttachType_index[i]:_BellAttachType_index[i+1]]
}