package world

import (
	"github.com/Tnze/go-mc/maths"
	"sync/atomic"
)

var entityCounter atomic.Int32

func NewEntityID() int32 {
	return entityCounter.Add(1)
}

type Entity struct {
	EntityID int32
	Position maths.Vec3d[float64]
	Rotation maths.Vec2d[float32]
	ChunkPos maths.Vec2d[int32]
	OnGround bool
	pos0     maths.Vec3d[float64]
	rot0     maths.Vec2d[float32]
}
