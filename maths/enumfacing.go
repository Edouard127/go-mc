package maths

type Facing int

const (
	Down  Facing = 0
	Up    Facing = 1
	North Facing = 2
	South Facing = 3
	West  Facing = 4
	East  Facing = 5
)

func (f Facing) Vec() (v Vec3d) {
	return [...]Vec3d{{Y: -1}, {Y: 1}, {Z: -1}, {Z: 1}, {X: -1}, {X: 1}}[f]
}

func GetClosestFacing(eyePos, blockPos Vec3d) Facing {
	var closest Facing
	var minDiff float64
	for _, side := range GetVisibleSides(eyePos, blockPos) {
		vector := side.Vec()
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

func GetVisibleSides(eyePos, blockPos Vec3d) []Facing {
	var sides []Facing
	blockPos.AddScalar(0.5, 0.5, 0.5)
	axis := checkAxis(eyePos.X-blockPos.X, West)
	if axis != -1 {
		sides = append(sides, axis)
	}
	axis = checkAxis(eyePos.Y-blockPos.Y, Down)
	if axis != -1 {
		sides = append(sides, axis)
	}
	axis = checkAxis(eyePos.Z-blockPos.Z, North)
	if axis != -1 {
		sides = append(sides, axis)
	}
	return sides
}

var EnumFacingValues = []Facing{Down, Up, North, South, West, East}

func (f Facing) Opposite() Facing {
	return EnumFacingValues[(f+3)%6]
}

func checkAxis(diff float64, negativeSide Facing) Facing {
	if diff < -0.5 {
		return negativeSide
	} else if diff > 0.5 {
		return negativeSide.Opposite()
	} else {
		return -1
	}
}
