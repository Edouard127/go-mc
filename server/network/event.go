package network

import (
	pk "github.com/Edouard127/go-mc/net/packet"
)

type Events[T any] struct {
	handlers map[int32]*HandlerHeap[T] // for specific packet id only
	tickers  *tickerHeap[T]
}

func NewEvents[T any]() *Events[T] {
	return &Events[T]{
		handlers: make(map[int32]*HandlerHeap[T]),
	}
}

func (e *Events[T]) GetHandlers(id int32) *HandlerHeap[T] {
	return e.handlers[id]
}

// AddListener adds a listener to the event.
// The listener will be called when the packet with the same ID is received.
// The listener will be called in the order of priority.
// The listeners cannot have multiple same ID.
func (e *Events[T]) AddListener(listeners ...PacketHandler[T]) {
	for _, l := range listeners {
		var s *HandlerHeap[T]
		var ok bool
		if s, ok = e.handlers[l.ID]; !ok {
			s = &HandlerHeap[T]{l}
			e.handlers[l.ID] = s
		} else {
			s.Push(l)
		}
	}
}

type TickHandler[T any] struct {
	Priority int
	F        func(*T) error
}

func (e *Events[T]) AddTicker(tickers ...TickHandler[T]) {
	for _, t := range tickers {
		if e.tickers == nil {
			e.tickers = &tickerHeap[T]{t}
		} else {
			e.tickers.Push(t)
		}
	}
}

type PacketHandler[T any] struct {
	ID       int32
	Priority int
	F        func(*T, pk.Packet) error
}

// HandlerHeap is PriorityQueue<PacketHandlerFunc>
type HandlerHeap[T any] []PacketHandler[T]

func (h HandlerHeap[T]) Len() int            { return len(h) }
func (h HandlerHeap[T]) Less(i, j int) bool  { return h[i].Priority < h[j].Priority }
func (h HandlerHeap[T]) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HandlerHeap[T]) Push(x interface{}) { *h = append(*h, x.(PacketHandler[T])) }
func (h *HandlerHeap[T]) Pop() interface{} {
	old := *h
	n := len(old)
	*h = old[0 : n-1]
	return old[n-1]
}

// tickerHeap is PriorityQueue<TickHandlerFunc>
type tickerHeap[T any] []TickHandler[T]

func (h tickerHeap[T]) Len() int { return len(h) }
func (h tickerHeap[T]) Less(i, j int) bool {
	return h[i].Priority < h[j].Priority
}
func (h tickerHeap[T]) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *tickerHeap[T]) Push(x interface{}) {
	*h = append(*h, x.(TickHandler[T]))
}
func (h *tickerHeap[T]) Pop() interface{} {
	old := *h
	n := len(old)
	*h = old[0 : n-1]
	return old[n-1]
}
