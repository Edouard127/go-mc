package maths

import (
	"golang.org/x/exp/constraints"
)

type AxisAlignedBB[T constraints.Integer | constraints.Float] struct {
	MinX, MinY, MinZ,
	MaxX, MaxY, MaxZ T
}

func (a AxisAlignedBB[T]) Contract(x, y, z T) AxisAlignedBB[T] {
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
	return AxisAlignedBB[T]{MinX: d0, MinY: d1, MinZ: d2, MaxX: d3, MaxY: d4, MaxZ: d5}
}

func (a AxisAlignedBB[T]) Expand(x, y, z T) AxisAlignedBB[T] {
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
	return AxisAlignedBB[T]{MinX: d0, MinY: d1, MinZ: d2, MaxX: d3, MaxY: d4, MaxZ: d5}
}

func (a AxisAlignedBB[T]) Inflate(x, y, z T) AxisAlignedBB[T] {
	d0 := a.MinX - x
	d1 := a.MinY - y
	d2 := a.MinZ - z
	d3 := a.MaxX + x
	d4 := a.MaxY + y
	d5 := a.MaxZ + z
	return AxisAlignedBB[T]{MinX: d0, MinY: d1, MinZ: d2, MaxX: d3, MaxY: d4, MaxZ: d5}
}

func (a AxisAlignedBB[T]) Intersect(other AxisAlignedBB[T]) AxisAlignedBB[T] {
	d0 := max(a.MinX, other.MinX)
	d1 := max(a.MinY, other.MinY)
	d2 := max(a.MinZ, other.MinZ)
	d3 := min(a.MaxX, other.MaxX)
	d4 := min(a.MaxY, other.MaxY)
	d5 := min(a.MaxZ, other.MaxZ)
	return AxisAlignedBB[T]{MinX: d0, MinY: d1, MinZ: d2, MaxX: d3, MaxY: d4, MaxZ: d5}
}

func (a AxisAlignedBB[T]) MinMax(other AxisAlignedBB[T]) AxisAlignedBB[T] {
	d0 := min(a.MinX, other.MinX)
	d1 := min(a.MinY, other.MinY)
	d2 := min(a.MinZ, other.MinZ)
	d3 := max(a.MaxX, other.MaxX)
	d4 := max(a.MaxY, other.MaxY)
	d5 := max(a.MaxZ, other.MaxZ)
	return AxisAlignedBB[T]{MinX: d0, MinY: d1, MinZ: d2, MaxX: d3, MaxY: d4, MaxZ: d5}
}

func (a AxisAlignedBB[T]) Move(x, y, z T) AxisAlignedBB[T] {
	return AxisAlignedBB[T]{MinX: a.MinX + x, MinY: a.MinY + y, MinZ: a.MinZ + z, MaxX: a.MaxX + x, MaxY: a.MaxY + y, MaxZ: a.MaxZ + z}
}

func (a AxisAlignedBB[T]) IntersectsWith(other AxisAlignedBB[T]) bool {
	return a.MinX < other.MaxX && a.MaxX > other.MinX && a.MinY < other.MaxY && a.MaxY > other.MinY && a.MinZ < other.MaxZ && a.MaxZ > other.MinZ
}

func (a AxisAlignedBB[T]) CollidesHorizontal(other AxisAlignedBB[T]) bool {
	return a.MinX <= other.MaxX && a.MaxX >= other.MinX && a.MinZ <= other.MaxZ && a.MaxZ >= other.MinZ
}

func (a AxisAlignedBB[T]) CollidesVertical(other AxisAlignedBB[T]) bool {
	return a.MinY <= other.MaxY && a.MaxY >= other.MinY
}

func (a AxisAlignedBB[T]) Contains(x, y, z T) bool {
	return x >= a.MinX && x < a.MaxX && y >= a.MinY && y < a.MaxY && z >= a.MinZ && z < a.MaxZ
}

func (a AxisAlignedBB[T]) Deflate(x, y, z T) AxisAlignedBB[T] {
	return a.Inflate(-x, -y, -z)
}
