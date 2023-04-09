package world

import (
	"fmt"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/bot/core"
	"github.com/Tnze/go-mc/bot/maths"
	. "github.com/Tnze/go-mc/level"
	"github.com/Tnze/go-mc/level/block"
	"math"
	"reflect"
)

type World struct {
	Columns  map[ChunkPos]*Chunk
	entities []*interface{}
}

func NewWorld() (w *World) {
	w = &World{
		Columns: make(map[ChunkPos]*Chunk),
	}
	return
}

func (w *World) AddEntity(e interface{}) error {
	if w.isValidEntity(&e) {
		w.entities = append(w.entities, &e)
		return nil
	}
	return fmt.Errorf("invalid entity")
}

func (w *World) RemoveEntity(e interface{}) error {
	if w.isValidEntity(&e) {
		if i, _, err := w.GetEntityByID(e.(*core.Entity).ID); err == nil {
			w.entities = append(w.entities[:i], w.entities[i+1:]...)
			return nil
		} else {
			return err
		}
	}
	return fmt.Errorf("invalid entity")
}

func (w *World) GetEntities() []*interface{} {
	return w.entities
}

func (w *World) GetEntitiesByType(t interface{}) []*interface{} {
	var entities []*interface{}
	for _, e := range w.entities {
		if reflect.TypeOf(*e) == reflect.TypeOf(t) {
			entities = append(entities, e)
		}
	}
	return entities
}

func (w *World) GetEntityByID(id int32) (int, interface{}, error) {
	for i, e := range w.entities {
		if t, ok := (*e).(*core.Entity); ok {
			if t.ID == id {
				return i, *e, nil
			}
		} else if t, ok := (*e).(*core.EntityLiving); ok {
			if t.ID == id {
				return i, *e, nil
			}
		} else if t, ok := (*e).(*core.EntityPlayer); ok {
			if t.ID == id {
				return i, *e, nil
			}
		}
	}
	return -1, nil, fmt.Errorf("entity not found")
}

/*
isValidEntity

	@param interface{} - The entity to check
	@return bool - Whether the entity is valid

	@description
		If the return value is true, then we do not need to do more type assertions because we know that the entity is valid.
		So it will always at least have the properties of core.Entity.
*/
func (w *World) isValidEntity(e *interface{}) bool {
	if _, ok := (*e).(core.Entity); ok {
		return true
	} else if _, ok := (*e).(*core.EntityLiving); ok {
		return true
	} else if _, ok := (*e).(*core.EntityPlayer); ok {
		return true
	}
	return false
}

func (w *World) GetBlock(pos maths.Vec3d[float64]) (block.Block, basic.Error) {
	chunkPos := ChunkPos{int32(pos.X) >> 4, int32(pos.Z) >> 4}
	if chunk, ok := w.Columns[chunkPos]; ok {
		return chunk.GetBlock(pos)
	} else {
		return block.StateList[block.ToStateID[block.Air{}]], basic.Error{Err: basic.InvalidChunk, Info: fmt.Errorf("chunk not found")}
	}
}

func (w *World) SetBlock(d maths.Vec3d[float64], i int) {
	chunkPos := ChunkPos{int32(d.X) >> 4, int32(d.Z) >> 4}
	if chunk, ok := w.Columns[chunkPos]; ok {
		chunk.SetBlock(d, i)
	}
}

func (w *World) GetNeighbors(block maths.Vec3d[float64]) []maths.Vec3d[float64] {
	return []maths.Vec3d[float64]{
		{X: block.X + 1, Y: block.Y, Z: block.Z},
		{X: block.X - 1, Y: block.Y, Z: block.Z},
		{X: block.X, Y: block.Y + 1, Z: block.Z},
		{X: block.X, Y: block.Y - 1, Z: block.Z},
		{X: block.X, Y: block.Y, Z: block.Z + 1},
		{X: block.X, Y: block.Y, Z: block.Z - 1},
	}
}

func (w *World) isChunkLoaded(pos ChunkPos) bool {
	_, ok := w.Columns[pos]
	return ok
}

func (w *World) RayTrace(start, end maths.Vec3d[float64]) (maths.RayTraceResult, basic.Error) {
	if start == maths.NullVec3d && end == maths.NullVec3d {
		return maths.RayTraceResult{}, basic.Error{Err: basic.NullValue, Info: fmt.Errorf("start and end cannot be null")}
	}

	for _, pos := range maths.RayTraceBlocks(start, end) {
		block, _ := w.GetBlock(pos)
		if block.IsAir() {
			continue
		} else {
			return maths.RayTraceResult{
				Position: pos,
			}, basic.Error{Err: basic.NoError, Info: nil}
		}
	}

	return maths.RayTraceResult{}, basic.Error{Err: basic.NoValue, Info: fmt.Errorf("no block found")}
}

func (w *World) GetBlockDensity(pos maths.Vec3d[float64], bb maths.AxisAlignedBB[float64]) float64 {
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

					if _, err := w.RayTrace(maths.Vec3d[float64]{X: d5 + d3, Y: d6, Z: d7 + d4}, pos); err.Is(basic.NoValue) {
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

func (w *World) IsAABBInMaterial(bb maths.AxisAlignedBB[float64]) bool {
	i := int32(math.Floor(bb.MinX))
	j := int32(math.Floor(bb.MaxX))
	k := int32(math.Floor(bb.MinY))
	l := int32(math.Floor(bb.MaxY))
	i1 := int32(math.Floor(bb.MinZ))
	j1 := int32(math.Floor(bb.MaxZ))

	for x := i; x <= j; x++ {
		for y := k; y <= l; y++ {
			for z := i1; z <= j1; z++ {
				if getBlock, err := w.GetBlock(maths.Vec3d[float64]{X: float64(x), Y: float64(y), Z: float64(z)}); !err.Is(basic.InvalidChunk) && !getBlock.IsAir() {
					if getBlock.IsLiquid() {
						return false
					}
				}
			}
		}
	}
	return true
}
