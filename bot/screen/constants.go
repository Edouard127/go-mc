package screen

type Mode int
type Button int

// Mode 0
const (
	LeftClick Button = iota
	RightClick
	LeftClickOutside  = -999
	RightClickOutside = -999
)

// Mode 1
const (
	ShiftLeftClick Button = iota
	ShiftRightClick
)

// Mode 2
const (
	KeyOne Button = iota
	KeyTwo
	KeyThree
	KeyFour
	KeyFive
	KeySix
	KeySeven
	KeyEight
	KeyNine
	SwapHand Button = 40
)

// Mode 3
const (
	MiddleClick Button = iota + 2
	_
)

// Mode 4
const (
	Drop Button = iota
	ControlDrop
)

// Mode 5
const (
	LeftClickDrag      Button = 0
	RightClickDrag            = 4
	MiddleClickDrag           = 8
	AddLeftClickDrag          = 1
	AddRightClickDrag         = 5
	AddMiddleClickDrag        = 9
)

// Mode 6
const (
	DoubleClick Button = iota
	_
)
