package world

import (
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/level"
	"github.com/Tnze/go-mc/maths"
)

type Client interface {
	ChunkViewer
	EntityViewer
	SendDisconnect(reason chat.Message)
	SendPlayerPosition(pos maths.Vec3d[float64], rot maths.Vec2d[float32]) (teleportID int32)
	SendSetChunkCacheCenter(chunkPos maths.Vec2d[int32])
}

type ChunkViewer interface {
	ViewChunkLoad(pos maths.Vec2d[int32], c *level.Chunk)
	ViewChunkUnload(pos maths.Vec2d[int32])
}

type EntityViewer interface {
	ViewAddPlayer(p *Player)
	ViewRemoveEntities(entityIDs []int32)
	ViewMoveEntityPos(id int32, delta maths.Vec3d[int16], onGround bool)
	ViewMoveEntityPosAndRot(id int32, delta maths.Vec3d[int16], rot maths.Vec2d[int8], onGround bool)
	ViewMoveEntityRot(id int32, rot maths.Vec2d[int8], onGround bool)
	ViewRotateHead(id int32, yaw int8)
	ViewTeleportEntity(id int32, pos maths.Vec3d[float64], rot maths.Vec2d[int8], onGround bool)
}
