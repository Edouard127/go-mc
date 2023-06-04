package world

import (
	"github.com/Edouard127/go-mc/maths"
	"github.com/Edouard127/go-mc/server/auth"
	"github.com/google/uuid"
	"time"
)

type Player struct {
	Entity
	Name       string
	UUID       uuid.UUID
	PubKey     *auth.PublicKey
	Properties []auth.Property
	Latency    time.Duration

	Gamemode       int32
	EntitiesInView map[int32]*Entity
	ViewDistance   int32
	view           *playerViewNode
}

func NewPlayer(name string, uuid uuid.UUID, key *auth.PublicKey, properties []auth.Property) *Player {
	return &Player{
		Name:       name,
		UUID:       uuid,
		PubKey:     key,
		Properties: properties,
	}
}

func (p *Player) chunkPosition() maths.Vec2d[int32] { return p.ChunkPos }
func (p *Player) chunkRadius() int32                { return p.ViewDistance }

// getView calculate the visual range enclosure with Position and ViewDistance of a player.
func (p *Player) getView() aabb3d {
	viewDistance := float64(p.ViewDistance) * 16 // the unit of ViewDistance is 1 Chunk（16 Block）
	return aabb3d{
		Upper: vec3d{p.Position.X + viewDistance, p.Position.Y + viewDistance, p.Position.Z + viewDistance},
		Lower: vec3d{p.Position.X - viewDistance, p.Position.Y - viewDistance, p.Position.Z - viewDistance},
	}
}
