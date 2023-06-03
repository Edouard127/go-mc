package server

import (
	"container/list"
	"strconv"
	"sync"

	pk "github.com/Tnze/go-mc/net/packet"
)

// Packet758 is a packet in protocol 757.
// We are using type system to force programmers to update packets.
type Packet758 pk.Packet
type Packet757 pk.Packet

type WritePacketError struct {
	Err error
	ID  int32
}

func (s WritePacketError) Error() string {
	return "server: send packet " + strconv.FormatInt(int64(s.ID), 16) + " error: " + s.Err.Error()
}

func (s WritePacketError) Unwrap() error {
	return s.Err
}

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
