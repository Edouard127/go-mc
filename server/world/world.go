package world

import (
	"errors"
	"github.com/Edouard127/go-mc/level"
	"github.com/Edouard127/go-mc/level/block"
	"github.com/Edouard127/go-mc/maths"
	"github.com/Edouard127/go-mc/server/internal/bvh"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"sync"
)

type World struct {
	log           *zap.Logger
	config        Config
	chunkProvider ChunkProvider

	chunks   map[maths.Vec2i]*LoadedChunk
	loaders  map[ChunkViewer]*loader
	tickLock sync.Mutex

	// playerViews is a BVH tree，storing the visual range collision boxes of each player.
	// the data structure is used to determine quickly which players to send notify when entity moves.
	playerViews playerViewTree
	players     map[Client]*Player

	Name string
}

type (
	vec3d          = bvh.Vec3[float64]
	aabb3d         = bvh.AABB[float64, vec3d]
	playerViewNode = bvh.Node[float64, aabb3d, playerView]
	playerViewTree = bvh.Tree[float64, aabb3d, playerView]
)

type playerView struct {
	EntityViewer
	*Player
}

type Config struct {
	ViewDistance  int32
	SpawnAngle    float32
	SpawnPosition [3]int32
}

func NewWorld(logger *zap.Logger, name string, provider ChunkProvider, config Config) (w *World) {
	w = &World{
		log:           logger,
		config:        config,
		chunks:        make(map[maths.Vec2i]*LoadedChunk),
		loaders:       make(map[ChunkViewer]*loader),
		players:       make(map[Client]*Player),
		chunkProvider: provider,
		Name:          name,
	}
	go w.tickLoop()
	return
}

func (w *World) SpawnPositionAndAngle() ([3]int32, float32) {
	return w.config.SpawnPosition, w.config.SpawnAngle
}

func (w *World) HashedSeed() [8]byte {
	return [8]byte{} // TODO
}

func (w *World) AddPlayer(c Client, p *Player, limiter *rate.Limiter) {
	w.tickLock.Lock()
	defer w.tickLock.Unlock()
	w.loaders[c] = newLoader(p, limiter)
	w.players[c] = p
	p.view = w.playerViews.Insert(p.getView(), playerView{c, p})
}

func (w *World) RemovePlayer(c Client, p *Player) {
	w.tickLock.Lock()
	defer w.tickLock.Unlock()
	w.log.Debug("Remove Player",
		zap.Int("loader count", len(w.loaders[c].loaded)),
		zap.Int("world count", len(w.chunks)),
	)
	// delete the player from all chunks which load the player.
	for pos := range w.loaders[c].loaded {
		if !w.chunks[pos].RemoveViewer(c) {
			w.log.Panic("viewer is not found in the loaded chunk")
		}
	}
	delete(w.loaders, c)
	delete(w.players, c)
	// delete the player from entity system.
	w.playerViews.Delete(p.view)
	w.playerViews.Find(
		/*bvh.TouchPoint[vec3d, aabb3d](bvh.Vec3d[float64](p.Position))*/ nil,
		func(n *playerViewNode) bool {
			n.Value.ViewRemoveEntities([]int32{p.EntityID})
			delete(n.Value.EntitiesInView, p.EntityID)
			return true
		},
	)
}

func (w *World) loadChunk(pos maths.Vec2i) bool {
	logger := w.log.With(zap.Int("x", pos.X), zap.Int("z", pos.Y))
	logger.Debug("Loading chunk")
	c, err := w.chunkProvider.GetChunk(pos)
	if err != nil {
		if errors.Is(err, errChunkNotExist) {
			logger.Debug("Generate chunk")
			// TODO: because there is no chunk generator，generate an empty chunk and mark it as generated
			c = level.EmptyChunk(24)
			stone := block.ToStateID[block.Stone]
			for s := range c.Sections {
				for i := 0; i < 16*16*16; i++ {
					c.Sections[s].SetBlock(i, stone)
				}
			}
			c.Status = level.StatusFull
		} else if !errors.Is(err, ErrReachRateLimit) {
			logger.Error("GetChunk error", zap.Error(err))
			return false
		}
	}
	w.chunks[pos] = &LoadedChunk{Chunk: c}
	return true
}

func (w *World) unloadChunk(pos maths.Vec2i) {
	logger := w.log.With(zap.Int("x", pos.X), zap.Int("z", pos.Y))
	logger.Debug("Unloading chunk")
	c, ok := w.chunks[pos]
	if !ok {
		logger.Panic("Unloading an non-exist chunk")
	}
	// notify all viewers who are watching the chunk to unload the chunk
	for _, viewer := range c.viewers {
		viewer.ViewChunkUnload(pos)
	}
	// move the chunk to provider and save
	err := w.chunkProvider.PutChunk(pos, c.Chunk)
	if err != nil {
		logger.Error("Store chunk data error", zap.Error(err))
	}
	delete(w.chunks, pos)
}
