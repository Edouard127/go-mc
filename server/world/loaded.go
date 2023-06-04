package world

import (
	"github.com/Edouard127/go-mc/level"
	"sync"
)

type LoadedChunk struct {
	sync.Mutex
	viewers []ChunkViewer
	*level.Chunk
}

func (lc *LoadedChunk) AddViewer(v ChunkViewer) {
	lc.Lock()
	defer lc.Unlock()
	for _, v2 := range lc.viewers {
		if v2 == v {
			panic("append an exist viewer")
		}
	}
	lc.viewers = append(lc.viewers, v)
}

func (lc *LoadedChunk) RemoveViewer(v ChunkViewer) bool {
	lc.Lock()
	defer lc.Unlock()
	for i, v2 := range lc.viewers {
		if v2 == v {
			last := len(lc.viewers) - 1
			lc.viewers[i] = lc.viewers[last]
			lc.viewers = lc.viewers[:last]
			return true
		}
	}
	return false
}
