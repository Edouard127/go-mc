package world

import (
	"fmt"
	"github.com/Edouard127/go-mc/bot/core"
	. "github.com/Edouard127/go-mc/level"
	"github.com/Edouard127/go-mc/level/block"
	"github.com/Edouard127/go-mc/maths"
	"math"
	"sync"
)

type World struct {
	Columns  map[ChunkPos]*Chunk
	Entities map[int32]core.Entity

	worldSync   sync.Mutex
	entityMutex sync.Mutex
}

func NewWorld() *World {
	return &World{
		Columns:  make(map[ChunkPos]*Chunk),
		Entities: make(map[int32]core.Entity),
	}
}

func (w *World) GetBlock(pos maths.Vec3d) (*block.Block, error) {
	w.worldSync.Lock()
	defer w.worldSync.Unlock()
	chunk, ok := w.Columns[ChunkPos{int32(pos.X) >> 4, int32(pos.Z) >> 4}]
	if ok {
		return chunk.GetBlock(pos)
	}
	return block.Air, fmt.Errorf("chunk not found")
}

func (w *World) MustGetBlock(pos maths.Vec3d) *block.Block {
	b, err := w.GetBlock(pos)
	if err != nil {
		panic(err)
	}
	return b
}

func (w *World) SetBlock(d maths.Vec3d, i int) {
	w.worldSync.Lock()
	defer w.worldSync.Unlock()
	chunk, ok := w.Columns[ChunkPos{int32(d.X) >> 4, int32(d.Z) >> 4}]
	if ok {
		chunk.SetBlock(d, i)
	}
}

func (w *World) GetNeighbors(block maths.Vec3d) []maths.Vec3d {
	return []maths.Vec3d{
		{X: block.X + 1, Y: block.Y, Z: block.Z},
		{X: block.X - 1, Y: block.Y, Z: block.Z},
		{X: block.X, Y: block.Y + 1, Z: block.Z},
		{X: block.X, Y: block.Y - 1, Z: block.Z},
		{X: block.X, Y: block.Y, Z: block.Z + 1},
		{X: block.X, Y: block.Y, Z: block.Z - 1},
	}
}

func (w *World) IsBlockLoaded(pos maths.Vec3d) bool {
	w.worldSync.Lock()
	defer w.worldSync.Unlock()
	chunkPos := ChunkPos{int32(pos.X) >> 4, int32(pos.Z) >> 4}
	if chunk, ok := w.Columns[chunkPos]; ok {
		return chunk.IsBlockLoaded(pos)
	}
	return false
}

func (w *World) IsChunkLoaded(pos ChunkPos) bool {
	w.worldSync.Lock()
	defer w.worldSync.Unlock()
	_, ok := w.Columns[pos]
	return ok
}

func (w *World) RayTrace(start, end maths.Vec3d) (maths.RayTraceResult, error) {
	if start.IsZero() && end.IsZero() {
		return maths.RayTraceResult{}, fmt.Errorf("start and end are null vectors")
	}

	for _, pos := range maths.RayTraceBlocks(start, end) {
		result, err := w.GetBlock(pos)
		if result == block.Air || err != nil {
			continue
		}
		return maths.RayTraceResult{pos}, nil
	}

	return maths.RayTraceResult{}, fmt.Errorf("no block found")
}

func (w *World) GetBlockDensity(pos maths.Vec3d, bb maths.AxisAlignedBB[float64]) float64 {
	d0 := 1.0 / ((bb.MaxX-bb.MinX)*2.0 + 1.0)
	d1 := 1.0 / ((bb.MaxY-bb.MinY)*2.0 + 1.0)
	d2 := 1.0 / ((bb.MaxZ-bb.MinZ)*2.0 + 1.0)
	d3 := (1.0 - math.Floor(1.0/d0)) * d0 / 2.0
	d4 := (1.0 - math.Floor(1.0/d2)) * d2 / 2.0

	if d0 >= 0.0 && d1 >= 0.0 && d2 >= 0.0 {
		j2 := 0.0
		k2 := 0.0

		for f := 0.0; f <= 1.0; f += d0 {
			for f1 := 0.0; f1 <= 1.0; f1 += d1 {
				for f2 := 0.0; f2 <= 1.0; f2 += d2 {
					d5 := bb.MinX + (bb.MaxX-bb.MinX)*f
					d6 := bb.MinY + (bb.MaxY-bb.MinY)*f1
					d7 := bb.MinZ + (bb.MaxZ-bb.MinZ)*f2

					if _, err := w.RayTrace(maths.Vec3d{X: d5 + d3, Y: d6, Z: d7 + d4}, pos); err != nil {
						j2++
					}
					k2++
				}
			}
		}

		return j2 / k2
	}
	return 0
}

func (w *World) Collides(bb maths.AxisAlignedBB[float64]) bool {
	return w.CollidesWithAnyBlock(bb) || w.CollidesWithAnyEntity(bb)
}

func (w *World) CollidesHorizontally(bb maths.AxisAlignedBB[float64]) bool {
	return w.pointCheck(bb, func(p maths.Vec3d) bool {
		return bb.CollidesHorizontal(p.ToAABB())
	})
}

func (w *World) CollidesVertically(bb maths.AxisAlignedBB[float64]) bool {
	return w.pointCheck(bb, func(p maths.Vec3d) bool {
		return bb.CollidesVertical(p.ToAABB())
	})
}

func (w *World) CollidesWithAnyBlock(bb maths.AxisAlignedBB[float64]) bool {
	return w.pointCheck(bb, func(p maths.Vec3d) bool {
		return !w.MustGetBlock(p).IsAir()
	})
}

func (w *World) CollidesWithAnyEntity(bb maths.AxisAlignedBB[float64]) bool {
	return len(w.GetEntitiesInAABB(bb)) > 0
}

func (w *World) pointCheck(bb maths.AxisAlignedBB[float64], predicate func(p maths.Vec3d) bool) bool {
	i := int32(math.Floor(bb.MinX))
	j := int32(math.Floor(bb.MaxX))
	k := int32(math.Floor(bb.MinY))
	l := int32(math.Floor(bb.MaxY))
	i1 := int32(math.Floor(bb.MinZ))
	j1 := int32(math.Floor(bb.MaxZ))

	for x := i; x <= j; x++ {
		for y := k; y <= l; y++ {
			for z := i1; z <= j1; z++ {
				if predicate(maths.Vec3d{X: float64(x), Y: float64(y), Z: float64(z)}) {
					return true
				}
			}
		}
	}

	return false
}

func (w *World) Add(e core.Entity) {
	w.entityMutex.Lock()
	defer w.entityMutex.Unlock()
	w.Entities[e.GetID()] = e
}

func (w *World) Delete(e core.Entity) {
	w.entityMutex.Lock()
	defer w.entityMutex.Unlock()
	delete(w.Entities, e.GetID())
}

func (w *World) GetEntity(id int32) core.Entity {
	result := w.entitySearch(1, 1, func(entity core.Entity) bool { return entity.GetID() == id })
	if len(result) == 0 {
		return nil
	}
	return result[0]
}

func (w *World) SearchEntity(f func(entity core.Entity) bool) core.Entity {
	result := w.entitySearch(1, 1, f)
	if len(result) == 0 {
		return nil
	}
	return result[0]
}

func (w *World) ClosestEntity(pos maths.Vec3d) core.Entity {
	var closestEntity core.Entity
	var closestDistance = math.MaxFloat64
	w.entitySearch(1, 1, func(entity core.Entity) bool {
		distance := pos.DistanceTo(entity.GetPosition())
		if closestEntity == nil || distance < closestDistance {
			closestEntity = entity
			closestDistance = distance
		}
		return false
	})
	return closestEntity
}

func (w *World) GetEntitiesInRange(pos maths.Vec3d, distance float64) []core.Entity {
	return w.entitySearch(1, len(w.Entities), func(entity core.Entity) bool {
		return pos.DistanceTo(entity.GetPosition()) <= distance
	})
}

func (w *World) GetEntitiesInAABB(bb maths.AxisAlignedBB[float64]) []core.Entity {
	return w.entitySearch(1, len(w.Entities), func(entity core.Entity) bool {
		return bb.IntersectsWith(entity.GetBoundingBox())
	})
}

func (w *World) GetEntitiesInAABBExcludingEntity(bb maths.AxisAlignedBB[float64], entity core.Entity) []core.Entity {
	return w.entitySearch(1, len(w.Entities), func(entityPredicate core.Entity) bool {
		return bb.IntersectsWith(entity.GetBoundingBox()) && entityPredicate != entity
	})
}

func (w *World) GetEntitiesNotInAABB(bb maths.AxisAlignedBB[float64]) []core.Entity {
	return w.entitySearch(1, len(w.Entities), func(entity core.Entity) bool {
		return !bb.IntersectsWith(entity.GetBoundingBox())
	})
}

func (w *World) entitySearch(nth, max int, predicate func(core.Entity) bool) []core.Entity {
	w.entityMutex.Lock()
	defer w.entityMutex.Unlock()
	var entities []core.Entity
	for i := range w.Entities {
		if predicate(w.Entities[i]) {
			nth--
			if nth <= 0 {
				if max > 1 {
					nth++
					entities = append(entities, w.Entities[i])
					continue
				}
				return []core.Entity{w.Entities[i]}
			}
		}
	}
	return entities
}
