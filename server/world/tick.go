package world

import (
	"github.com/Edouard127/go-mc/maths"
	"time"
)

func (w *World) tickLoop() {
	var n uint
	for range time.Tick(time.Microsecond * 20) {
		w.tick(n)
		n++
	}
}

func (w *World) tick(n uint) {
	w.tickLock.Lock()
	defer w.tickLock.Unlock()

	if n%8 == 0 {
		w.subtickChunkLoad()
	}
	w.subtickUpdatePlayers()
	w.subtickUpdateEntities()
}

func (w *World) subtickChunkLoad() {
	for c, p := range w.players {
		x := int(p.Position.X) >> 4
		z := int(p.Position.Z) >> 4
		newChunkPos := maths.Vec2i{X: x, Y: z}
		if newChunkPos != p.ChunkPos {
			p.ChunkPos = newChunkPos
			c.SendSetChunkCacheCenter(newChunkPos)
		}
	}
	// because of the random traversal order of w.loaders, every loader has the same opportunity, so it's relatively fair.
LoadChunk:
	for viewer, loader := range w.loaders {
		loader.calcLoadingQueue()
		for _, pos := range loader.loadQueue {
			if !loader.limiter.Allow() { // We reach the player limit. Skip
				break
			}
			if _, ok := w.chunks[pos]; !ok {
				if !w.loadChunk(pos) {
					break LoadChunk // We reach the global limit. skip
				}
			}
			loader.loaded[pos] = struct{}{}
			lc := w.chunks[pos]
			lc.AddViewer(viewer)
			lc.Lock()
			viewer.ViewChunkLoad(pos, lc.Chunk)
			lc.Unlock()
		}
	}
	for viewer, loader := range w.loaders {
		loader.calcUnusedChunks()
		for _, pos := range loader.unloadQueue {
			delete(loader.loaded, pos)
			if !w.chunks[pos].RemoveViewer(viewer) {
				w.log.Panic("viewer is not found in the loaded chunk")
			}
			viewer.ViewChunkUnload(pos)
		}
	}
	var unloadQueue []maths.Vec2i
	for pos, chunk := range w.chunks {
		if len(chunk.viewers) == 0 {
			unloadQueue = append(unloadQueue, pos)
		}
	}
	for i := range unloadQueue {
		w.unloadChunk(unloadQueue[i])
	}
}

func (w *World) subtickUpdatePlayers() {
	/*
		for c, p := range w.players {
			if !p.Inputs.TryLock() {
				continue
			}
			inputs := &p.Inputs
			// update the range of visual.
			if p.ViewDistance != int32(inputs.ViewDistance) {
				p.ViewDistance = int32(inputs.ViewDistance)
				p.view = w.playerViews.Insert(p.getView(), w.playerViews.Delete(p.view))
			}
			// delete entities that not in range from entities lists of each player.
			for id, e := range p.EntitiesInView {
				if !p.view.Box.WithIn(e.Position) {
					delete(p.EntitiesInView, id) // it should be safe to delete element from a map being traversed.
					p.view.Value.ViewRemoveEntities([]int32{id})
				}
			}
			if p.teleport != nil {
				if inputs.TeleportID == p.teleport.ID {
					p.pos0 = p.teleport.Position
					p.rot0 = p.teleport.Rotation
					p.teleport = nil
				}
			} else {
				delta := [3]float64{
					inputs.Position[0] - p.Position[0],
					inputs.Position[1] - p.Position[1],
					inputs.Position[2] - p.Position[2],
				}
				distance := math.Sqrt(delta[0]*delta[0] + delta[1]*delta[1] + delta[2]*delta[2])
				if distance > 100 {
					// You moved too quickly :( (Hacking?)
					teleportID := c.SendPlayerPosition(p.Position, p.Rotation)
					p.teleport = &TeleportRequest{
						ID:       teleportID,
						Position: p.Position,
						Rotation: p.Rotation,
					}
				} else if inputs.Position.IsValid() {
					p.pos0 = inputs.Position
					p.rot0 = inputs.Rotation
					p.OnGround = inputs.OnGround
				} else {
					w.log.Info("Player move invalid",
						zap.Float64("x", inputs.Position[0]),
						zap.Float64("y", inputs.Position[1]),
						zap.Float64("z", inputs.Position[2]),
					)
					c.SendDisconnect(chat.TranslateMsg("multiplayer.disconnect.invalid_player_movement"))
				}
			}
			p.Inputs.Unlock()
		}
	*/
}

func (w *World) subtickUpdateEntities() {
	// TODO: entity list should be traversed here, but players are the only entities now.
	for _, e := range w.players {
		// sending Update Entity Position pack to every player who can see it, when it moves.
		var delta maths.Vec3s
		var rot maths.Vec2b
		if e.Position != e.pos0 { // TODO: send Teleport Entity pack instead when moving distance is greater than 8.
			delta = maths.Vec3s{
				X: int16((e.pos0.X - e.Position.X) * 32 * 128),
				Y: int16((e.pos0.Y - e.Position.Y) * 32 * 128),
				Z: int16((e.pos0.Z - e.Position.Z) * 32 * 128),
			}
		}
		if e.Rotation != e.rot0 {
			rot = maths.Vec2b{
				X: int8(e.rot0.X * 256 / 360),
				Y: int8(e.rot0.Y * 256 / 360),
			}
		}
		//cond := bvh.TouchPoint[maths.Vec3d[float64], aabb3d](e.Position)
		w.playerViews.Find(nil,
			func(n *playerViewNode) bool {
				if n.Value.Player == e {
					return true // don't send the player self to the player
				}
				// check if the current entity is in range of player visual. if so, moving data will be forwarded.
				if _, ok := n.Value.EntitiesInView[e.EntityID]; !ok {
					// add the entity to the entity list of the player
					n.Value.ViewAddPlayer(e)
					n.Value.EntitiesInView[e.EntityID] = &e.Entity
				}
				return true
			},
		)
		var sendMove func(v EntityViewer)
		switch {
		case e.Position != e.pos0 && e.Rotation != e.rot0:
			sendMove = func(v EntityViewer) {
				v.ViewMoveEntityPosAndRot(e.EntityID, delta, rot, e.OnGround)
				v.ViewRotateHead(e.EntityID, rot.X)
			}
		case e.Position != e.pos0:
			sendMove = func(v EntityViewer) {
				v.ViewMoveEntityPos(e.EntityID, delta, e.OnGround)
			}
		case e.Rotation != e.rot0:
			sendMove = func(v EntityViewer) {
				v.ViewMoveEntityRot(e.EntityID, rot, e.OnGround)
				v.ViewRotateHead(e.EntityID, rot.X)
			}
		default:
			continue
		}
		e.Position = e.pos0
		e.Rotation = e.rot0
		w.playerViews.Find(nil,
			func(n *playerViewNode) bool {
				if n.Value.Player == e {
					return true // not sending self movements to player self.
				}
				// check if the current entity is in the player visual entities list. if so, moving data will be forwarded.
				if _, ok := n.Value.EntitiesInView[e.EntityID]; ok {
					sendMove(n.Value.EntityViewer)
				} else {
					// or the entity will be add to the entities list of the player
					// TODO: deal with the situation that the entity is not a player
					n.Value.ViewAddPlayer(e)
					n.Value.EntitiesInView[e.EntityID] = &e.Entity
				}
				return true
			},
		)
	}
}
