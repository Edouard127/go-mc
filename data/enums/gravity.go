package enums

const (
	Gravity              = 0.08
	SlowFallGravity      = 0.01
	AirDrag              = 0.98
	YawSpeed             = 3.0
	PitchSpeed           = 3.0
	PlayerSpeed          = 0.1
	SprintSpeed          = 0.3
	SneakSpeed           = 0.3
	StepHeight           = 0.6
	NegligibleVelocity   = 0.003
	SoulSandMultiplier   = 0.4
	HoneyBlockMultiplier = 0.4
	LadderMaxSpeed       = 0.15
	LadderAcceleration   = 0.2
	WaterDrag            = 0.8
	WaterSprintDrag      = 0.9
	WaterDolphinDrag     = 0.96
	LavaDrag             = 0.5
	WaterGravity         = 0.02
	LavaGravity          = 0.02
	LiquidAcceleration   = 0.02
	AirBornInertia       = 0.91
	AirBornAcceleration  = 0.02
	DefaultSlipperiness  = 0.6
	OutOfLiquidImpulse   = 0.3
)

/*func Slipperiness(b level.BlocksState) float64 {
	if t, ok := slipperiness[b]; ok {
		return t
	} else {
		return DefaultSlipperiness
	}
}

var slipperiness = map[level.BlocksState]float64{
	block.ToStateID[block.SoulSand]:   SoulSandMultiplier,
	block.ToStateID[block.HoneyBlock]: HoneyBlockMultiplier,
	block.ToStateID[block.SlimeBlock]: 0.8,
	block.ToStateID[block.Ice]:        0.98,
	block.ToStateID[block.PackedIce]:  0.98,
	block.ToStateID[block.FrostedIce]: 0.98,
}*/
