// Code generated by "stringer -type=Direction -output=direction_string.go"; DO NOT EDIT.

package properties

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Down-0]
	_ = x[Up-1]
	_ = x[North-2]
	_ = x[South-3]
	_ = x[West-4]
	_ = x[East-5]
}

const _Direction_name = "DownUpNorthSouthWestEast"

var _Direction_index = [...]uint8{0, 4, 6, 11, 16, 20, 24}

func (i Direction) String() string {
	if i >= Direction(len(_Direction_index)-1) {
		return "Direction(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Direction_name[_Direction_index[i]:_Direction_index[i+1]]
}
