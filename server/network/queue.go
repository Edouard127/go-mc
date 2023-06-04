package network

import (
	"container/list"
	"sync"

	pk "github.com/Edouard127/go-mc/net/packet"
)

type PacketQueue struct {
	queue  *list.List
	closed bool
	cond   sync.Cond
}

func NewPacketQueue() (p *PacketQueue) {
	return &PacketQueue{
		queue: list.New(),
		cond:  sync.Cond{L: new(sync.Mutex)},
	}
}

func (p *PacketQueue) Push(packet pk.Packet) {
	p.cond.L.Lock()
	if !p.closed {
		p.queue.PushBack(packet)
	}
	p.cond.Signal()
	p.cond.L.Unlock()
}

func (p *PacketQueue) Pull() (packet pk.Packet, ok bool) {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()
	for p.queue.Front() == nil && !p.closed {
		p.cond.Wait()
	}
	if p.closed {
		return pk.Packet{}, false
	}
	packet = p.queue.Remove(p.queue.Front()).(pk.Packet)
	ok = true
	return
}

func (p *PacketQueue) Close() {
	p.cond.L.Lock()
	p.closed = true
	p.cond.Broadcast()
	p.cond.L.Unlock()
}
