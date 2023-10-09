package maths

type AxisAlignedBB struct {
	MinX, MinY, MinZ,
	MaxX, MaxY, MaxZ float64
}

func (a *AxisAlignedBB) Contract(x, y, z float64) {
	d0 := a.MinX
	d1 := a.MinY
	d2 := a.MinZ
	d3 := a.MaxX
	d4 := a.MaxY
	d5 := a.MaxZ
	if x < 0.0 {
		d0 -= x
	}
	if x > 0.0 {
		d3 -= x
	}
	if y < 0.0 {
		d1 -= y
	}
	if y > 0.0 {
		d4 -= y
	}
	if z < 0.0 {
		d2 -= z
	}
	if z > 0.0 {
		d5 -= z
	}
	a.MinX = d0
	a.MinY = d1
	a.MinZ = d2
	a.MaxX = d3
	a.MaxY = d4
	a.MaxZ = d5
}

func (a *AxisAlignedBB) Expand(x, y, z float64) {
	d0 := a.MinX
	d1 := a.MinY
	d2 := a.MinZ
	d3 := a.MaxX
	d4 := a.MaxY
	d5 := a.MaxZ
	if x < 0.0 {
		d0 += x
	}
	if x > 0.0 {
		d3 += x
	}
	if y < 0.0 {
		d1 += y
	}
	if y > 0.0 {
		d4 += y
	}
	if z < 0.0 {
		d2 += z
	}
	if z > 0.0 {
		d5 += z
	}
	a.MinX = d0
	a.MinY = d1
	a.MinZ = d2
	a.MaxX = d3
	a.MaxY = d4
	a.MaxZ = d5
}

func (a *AxisAlignedBB) Unexpand(x, y, z float64) {
	a.Expand(-x, -y, -z)
}

func (a *AxisAlignedBB) Inflate(x, y, z float64) {
	a.MinX -= x
	a.MinY -= y
	a.MinZ -= z
	a.MaxX += x
	a.MaxY += y
	a.MaxZ += z
}

func (a *AxisAlignedBB) Intersect(other AxisAlignedBB) {
	a.MinX = max(a.MinX, other.MinX)
	a.MinY = max(a.MinY, other.MinY)
	a.MinZ = max(a.MinZ, other.MinZ)
	a.MaxX = min(a.MaxX, other.MaxX)
	a.MaxY = min(a.MaxY, other.MaxY)
	a.MaxZ = min(a.MaxZ, other.MaxZ)
}

func (a *AxisAlignedBB) MinMax(other AxisAlignedBB) {
	a.MinX = min(a.MinX, other.MinX)
	a.MinY = min(a.MinY, other.MinY)
	a.MinZ = min(a.MinZ, other.MinZ)
	a.MaxX = max(a.MaxX, other.MaxX)
	a.MaxY = max(a.MaxY, other.MaxY)
	a.MaxZ = max(a.MaxZ, other.MaxZ)
}

func (a *AxisAlignedBB) Move(x, y, z float64) {
	a.MinX += x
	a.MinY += y
	a.MinZ += z
	a.MaxX += x
	a.MaxY += y
	a.MaxZ += z
}

func (a *AxisAlignedBB) IntersectsWith(other AxisAlignedBB) bool {
	return a.MinX < other.MaxX && a.MaxX > other.MinX && a.MinY < other.MaxY && a.MaxY > other.MinY && a.MinZ < other.MaxZ && a.MaxZ > other.MinZ
}

func (a *AxisAlignedBB) Bottom() Vec3d {
	return Vec3d{a.MinX + (a.MaxX-a.MinX)*0.5, a.MinY, a.MinZ + (a.MaxZ-a.MinZ)*0.5}
}

func (a *AxisAlignedBB) Center() Vec3d {
	return Vec3d{a.MinX + (a.MaxX-a.MinX)*0.5, a.MinY + (a.MaxY-a.MinY)*0.5, a.MinZ + (a.MaxZ-a.MinZ)*0.5}
}

func (a *AxisAlignedBB) Top() Vec3d {
	return Vec3d{a.MinX + (a.MaxX-a.MinX)*0.5, a.MaxY, a.MinZ + (a.MaxZ-a.MinZ)*0.5}
}

func (a *AxisAlignedBB) Contains(x, y, z float64) bool {
	return x >= a.MinX && x < a.MaxX && y >= a.MinY && y < a.MaxY && z >= a.MinZ && z < a.MaxZ
}

func (a *AxisAlignedBB) Deflate(x, y, z float64) {
	a.Inflate(-x, -y, -z)
}
