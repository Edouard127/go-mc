package world

import (
	"compress/gzip"
	"errors"
	"fmt"
	"github.com/Tnze/go-mc/level"
	"github.com/Tnze/go-mc/maths"
	"github.com/Tnze/go-mc/save"
	"github.com/Tnze/go-mc/save/region"
	"github.com/Tnze/go-mc/server/auth"
	"github.com/google/uuid"
	"golang.org/x/time/rate"
	"io/fs"
	"os"
	"path/filepath"
)

// ChunkProvider implements chunk storage
type ChunkProvider struct {
	dir     string
	limiter *rate.Limiter
}

func NewProvider(dir string, limiter *rate.Limiter) ChunkProvider {
	return ChunkProvider{dir: dir, limiter: limiter}
}

var ErrReachRateLimit = errors.New("reach rate limit")

func (p *ChunkProvider) GetChunk(pos maths.Vec2d[int32]) (c *level.Chunk, errRet error) {
	if !p.limiter.Allow() {
		return nil, ErrReachRateLimit
	}
	r, err := p.getRegion(pos.X, pos.Y)
	if err != nil {
		return nil, fmt.Errorf("open region fail: %w", err)
	}
	defer func(r *region.Region) {
		err2 := r.Close()
		if errRet == nil && err2 != nil {
			errRet = fmt.Errorf("close region fail: %w", err2)
		}
	}(r)

	x, z := region.In(pos.X, pos.Y)
	if !r.ExistSector(x, z) {
		return nil, errChunkNotExist
	}

	data, err := r.ReadSector(x, z)
	if err != nil {
		return nil, fmt.Errorf("read sector fail: %w", err)
	}

	var chunk save.Chunk
	if err := chunk.Load(data); err != nil {
		return nil, fmt.Errorf("parse chunk data fail: %w", err)
	}

	c, err = level.ChunkFromSave(&chunk)
	if err != nil {
		return nil, fmt.Errorf("load chunk data fail: %w", err)
	}
	return c, nil
}

func (p *ChunkProvider) getRegion(rx, rz int32) (*region.Region, error) {
	filename := fmt.Sprintf("r.%d.%d.mca", rx, rz)
	path := filepath.Join(p.dir, filename)
	r, err := region.Open(path)
	if errors.Is(err, fs.ErrNotExist) {
		r, err = region.Create(path)
	}
	return r, err
}

func (p *ChunkProvider) PutChunk(pos maths.Vec2d[int32], c *level.Chunk) (err error) {
	//var chunk save.Chunk
	//err = level.ChunkToSave(c, &chunk)
	//if err != nil {
	//	return fmt.Errorf("encode chunk data fail: %w", err)
	//}
	//
	//data, err := chunk.Data(1)
	//if err != nil {
	//	return fmt.Errorf("record chunk data fail: %w", err)
	//}
	//
	//r, err := p.getRegion(region.At(int(pos[0]), int(pos[1])))
	//if err != nil {
	//	return fmt.Errorf("open region fail: %w", err)
	//}
	//defer func(r *region.Region) {
	//	err2 := r.Close()
	//	if err == nil && err2 != nil {
	//		err = fmt.Errorf("open region fail: %w", err)
	//	}
	//}(r)
	//
	//x, z := region.In(int(pos[0]), int(pos[1]))
	//err = r.WriteSector(x, z, data)
	//if err != nil {
	//	return fmt.Errorf("write sector fail: %w", err)
	//}

	return nil
}

var errChunkNotExist = errors.New("ErrChunkNotExist")

type PlayerProvider struct {
	dir string
}

func NewPlayerProvider(dir string) *PlayerProvider {
	return &PlayerProvider{dir: dir}
}

func (p *PlayerProvider) readPlayerData(id uuid.UUID) (data save.PlayerData, err error) {
	f, err := os.Open(filepath.Join(p.dir, id.String()+".dat"))
	if err != nil {
		return data, err
	}

	/*
		Most of the time you don't have to worry about that case
		if close fails, there are problems deeper than your application generally
	*/
	defer f.Close()

	r, err := gzip.NewReader(f)
	if err != nil {
		return data, fmt.Errorf("open gzip reader fail: %w", err)
	}

	defer r.Close()

	data, err = save.ReadPlayerData(r)
	if err != nil {
		return data, fmt.Errorf("read player data fail: %w", err)
	}
	return data, nil
}

func (p *PlayerProvider) GetPlayer(name string, id uuid.UUID, pubKey *auth.PublicKey, properties []auth.Property) (player *Player, errRet error) {
	data, err := p.readPlayerData(id)
	if err != nil {
		return nil, fmt.Errorf("read player data fail: %w", err)
	}

	player = &Player{
		Entity: Entity{
			EntityID: NewEntityID(),
			Position: maths.ParseVec3d[float64](data.Pos[0], data.Pos[1], data.Pos[2]),
			Rotation: maths.ParseVec2d[float32](data.Rotation[0], data.Rotation[1]),
			ChunkPos: maths.Vec2d[int32]{X: int32(data.Pos[0]) >> 5, Y: int32(data.Pos[2]) >> 5},
		},
		Name:           name,
		UUID:           id,
		PubKey:         pubKey,
		Properties:     properties,
		Gamemode:       data.PlayerGameType,
		EntitiesInView: make(map[int32]*Entity),
		ViewDistance:   10,
	}
	return
}
