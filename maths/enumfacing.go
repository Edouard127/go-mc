package maths

type EnumFacing int8

const (
	DOWN  EnumFacing = 0
	UP    EnumFacing = 1
	NORTH EnumFacing = 2
	SOUTH EnumFacing = 3
	WEST  EnumFacing = 4
	EAST  EnumFacing = 5
)

var EnumFacingValues = []EnumFacing{DOWN, UP, NORTH, SOUTH, WEST, EAST}

func (f EnumFacing) Vector() (v Vec3d) {
	switch f {
	case DOWN:
		v = Vec3d{Y: -1}
	case UP:
		v = Vec3d{Y: 1}
	case NORTH:
		v = Vec3d{Z: -1}
	case SOUTH:
		v = Vec3d{Z: 1}
	case WEST:
		v = Vec3d{X: -1}
	case EAST:
		v = Vec3d{X: 1}
	}
	return
}

func GetClosestFacing(eyePos, blockPos Vec3d) EnumFacing {
	var closest EnumFacing
	var minDiff float64
	for _, side := range GetVisibleSides(eyePos, blockPos) {
		vector := side.Vector()
		vector.AddScalar(0.5, 0.5, 0.5)
		blockPos.Add(vector)
		diff := eyePos.DistanceTo(blockPos)
		if minDiff == 0 || diff < minDiff {
			minDiff = diff
			closest = side
		}
	}
	return closest
}

func GetVisibleSides(eyePos, blockPos Vec3d) []EnumFacing {
	var sides []EnumFacing
	blockPos.AddScalar(0.5, 0.5, 0.5)
	axis := checkAxis(eyePos.X-blockPos.X, WEST)
	if axis != -1 {
		sides = append(sides, axis)
	}
	axis = checkAxis(eyePos.Y-blockPos.Y, DOWN)
	if axis != -1 {
		sides = append(sides, axis)
	}
	axis = checkAxis(eyePos.Z-blockPos.Z, NORTH)
	if axis != -1 {
		sides = append(sides, axis)
	}
	return sides
}

func (f EnumFacing) GetOpposite() EnumFacing {
	return EnumFacingValues[(f+3)%6]
}

func checkAxis(diff float64, negativeSide EnumFacing) EnumFacing {
	if diff < -0.5 {
		return negativeSide
	} else if diff > 0.5 {
		return negativeSide.GetOpposite()
	} else {
		return -1
	}
}
