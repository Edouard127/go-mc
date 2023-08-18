package maths

type RayTraceResult struct {
	Position Vec3d
}

func RayTraceBlocks(start, end Vec3d) []Vec3d {
	var result []Vec3d
	diff := end.Sub(start)
	distance := diff.Length()
	if distance == 0 {
		return result
	}
	for i := 0; i < int(distance); i++ {
		result = append(result, start.Add(diff.MulScalar(float64(i)/distance)))
	}
	return result
}
