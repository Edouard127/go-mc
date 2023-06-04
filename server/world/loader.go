package world

import (
	"github.com/Tnze/go-mc/maths"
	"golang.org/x/time/rate"
	"math"
	"sort"
)

// loader takes part in chunk loading，each loader contains a position 'pos' and a radius 'r'
// chunks pointed by the position, and the radius of loader will be load。
type loader struct {
	loaderSource
	loaded      map[maths.Vec2d[int32]]struct{}
	loadQueue   []maths.Vec2d[int32]
	unloadQueue []maths.Vec2d[int32]
	limiter     *rate.Limiter
}

type loaderSource interface {
	chunkPosition() maths.Vec2d[int32]
	chunkRadius() int32
}

func newLoader(source loaderSource, limiter *rate.Limiter) (l *loader) {
	l = &loader{
		loaderSource: source,
		loaded:       make(map[maths.Vec2d[int32]]struct{}),
		limiter:      limiter,
	}
	l.calcLoadingQueue()
	return
}

// calcLoadingQueue calculate the chunks which loader point.
// The result is stored in l.loadQueue and the previous will be removed.
func (l *loader) calcLoadingQueue() {
	l.loadQueue = l.loadQueue[:0]
	for _, v := range loadList[:radiusIdx[l.chunkRadius()]] {
		pos := l.chunkPosition()
		pos.X, pos.Y = pos.X+v.X, pos.Y+v.Y
		if _, ok := l.loaded[pos]; !ok {
			l.loadQueue = append(l.loadQueue, pos)
		}
	}
}

// calcUnusedChunks calculate the chunks the loader wants to remove.
// Behaviour is same with calcLoadingQueue.
func (l *loader) calcUnusedChunks() {
	l.unloadQueue = l.unloadQueue[:0]
	for chunk := range l.loaded {
		player := l.chunkPosition()
		r := l.chunkRadius()
		if distance2i(maths.Vec2d[int32]{X: chunk.X - player.X, Y: chunk.Y - player.Y}) > float64(r) {
			l.loadQueue = append(l.loadQueue, chunk)
		}
	}
}

// loadList is chunks in a certain distance of (0, 0), order by Euclidean distance
// the more forward the chunk is, the closer it to (0, 0)
var loadList []maths.Vec2d[int32]

// radiusIdx[i] is the count of chunks in loadList and the distance of i
var radiusIdx []int

func init() {
	const maxR int32 = 32

	// calculate loadList
	for x := -maxR; x <= maxR; x++ {
		for z := -maxR; z <= maxR; z++ {
			pos := maths.Vec2d[int32]{X: x, Y: z}
			if distance2i(pos) < float64(maxR) {
				loadList = append(loadList, pos)
			}
		}
	}
	sort.Slice(loadList, func(i, j int) bool {
		return distance2i(loadList[i]) < distance2i(loadList[j])
	})

	// calculate radiusIdx
	radiusIdx = make([]int, maxR+1)
	for i, v := range loadList {
		r := int32(math.Ceil(distance2i(v)))
		if r > maxR {
			break
		}
		radiusIdx[r] = i
	}
}

// distance calculates the Euclidean distance that a position to the origin point
func distance2i(pos maths.Vec2d[int32]) float64 {
	return math.Sqrt(float64(pos.X*pos.X) + float64(pos.Y*pos.Y))
}
