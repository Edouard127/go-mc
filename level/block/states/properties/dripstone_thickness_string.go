// Code generated by "stringer -type=DripstoneThickness -output=dripstone_thickness_string.go"; DO NOT EDIT.

package properties

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DripstoneThicknessTipMerge-0]
	_ = x[DripstoneThicknessTip-1]
	_ = x[DripstoneThicknessFrustum-2]
	_ = x[DripstoneThicknessMiddle-3]
	_ = x[DripstoneThicknessBase-4]
}

const _DripstoneThickness_name = "DripstoneThicknessTipMergeDripstoneThicknessTipDripstoneThicknessFrustumDripstoneThicknessMiddleDripstoneThicknessBase"

var _DripstoneThickness_index = [...]uint8{0, 26, 47, 72, 96, 118}

func (i DripstoneThickness) String() string {
	if i >= DripstoneThickness(len(_DripstoneThickness_index)-1) {
		return "DripstoneThickness(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DripstoneThickness_name[_DripstoneThickness_index[i]:_DripstoneThickness_index[i+1]]
}
