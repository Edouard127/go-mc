package world

import (
	"github.com/Edouard127/go-mc/chat"
	"github.com/Edouard127/go-mc/level"
	"github.com/Edouard127/go-mc/maths"
)

type Client interface {
	ChunkViewer
	EntityViewer
	SendDisconnect(reason chat.Message)
	SendPlayerPosition(pos maths.Vec3d, rot maths.Vec2f) (teleportID int32)
	SendSetChunkCacheCenter(chunkPos maths.Vec2i)
}

type ChunkViewer interface {
	ViewChunkLoad(pos maths.Vec2i, c *level.Chunk)
	ViewChunkUnload(pos maths.Vec2i)
}

type EntityViewer interface {
	ViewAddPlayer(p *Player)
	ViewRemoveEntities(entityIDs []int32)
	ViewMoveEntityPos(id int32, delta maths.Vec3s, onGround bool)
	ViewMoveEntityPosAndRot(id int32, delta maths.Vec3s, rot maths.Vec2b, onGround bool)
	ViewMoveEntityRot(id int32, rot maths.Vec2b, onGround bool)
	ViewRotateHead(id int32, yaw int8)
	ViewTeleportEntity(id int32, pos maths.Vec3d, rot maths.Vec2b, onGround bool)
}
