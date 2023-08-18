package world

import (
	"github.com/Edouard127/go-mc/maths"
	"sync/atomic"
)

var entityCounter atomic.Int32

func NewEntityID() int32 {
	return entityCounter.Add(1)
}

type Entity struct {
	EntityID int32
	Position maths.Vec3d
	Rotation maths.Vec2f
	ChunkPos maths.Vec2i
	OnGround bool
	pos0     maths.Vec3d
	rot0     maths.Vec2f
}
