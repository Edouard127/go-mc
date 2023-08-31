package enums

type PlayerAction int

const (
	PressShiftKey PlayerAction = iota
	ReleaseShiftKey
	StopSleeping
	StartSprinting
	StopSprinting
	StartJumpWithHorse
	StopJumpWithHorse
	OpenHorseInventory
	StartFlyingWithElytra
)
