package maths

type RayTraceResult struct {
	Position Vec3d
}

func RayTraceBlocks(start, end Vec3d) []Vec3d {
	var result []Vec3d
	end.Sub(start)
	distance := end.Length()
	if distance == 0 {
		return result
	}
	for i := 0; i < int(distance); i++ {
		idk := float64(i) / distance
		result = append(result, Vec3d{
			X: start.X + end.X*idk,
			Y: start.Y + end.Y*idk,
			Z: start.Z + end.Z*idk,
		})
	}
	return result
}
