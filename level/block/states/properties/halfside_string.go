// Code generated by "stringer -type=Half -output=halfside_string.go"; DO NOT EDIT.

package properties

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[HalfTop-0]
	_ = x[HalfBottom-1]
}

const _Half_name = "HalfTopHalfBottom"

var _Half_index = [...]uint8{0, 7, 17}

func (i Half) String() string {
	if i >= Half(len(_Half_index)-1) {
		return "Half(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Half_name[_Half_index[i]:_Half_index[i+1]]
}