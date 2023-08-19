package provider

import (
	pk "github.com/Edouard127/go-mc/net/packet"
)

type Events[T any] struct {
	generic  *handlerHeap[T]           // for every packet
	handlers map[int32]*handlerHeap[T] // for specific packet id only
	tickers  *tickerHeap[T]
}

// AddListener adds a listener to the event.
// The listener will be called when the packet with the same ID is received.
// The listener will be called in the order of priority.
// The listeners cannot have multiple same ID.
func (e *Events[T]) AddListener(listeners ...PacketHandler[T]) {
	for _, l := range listeners {
		var s *handlerHeap[T]
		var ok bool
		if s, ok = e.handlers[l.ID]; !ok {
			s = &handlerHeap[T]{l}
			e.handlers[l.ID] = s
		} else {
			s.Push(l)
		}
	}
}

// AddGeneric adds listeners like AddListener, but the packet ID is ignored.
// Generic listener is always called before specific packet listener.
func (e *Events[T]) AddGeneric(listeners ...PacketHandler[T]) {
	for _, l := range listeners {
		if e.generic == nil {
			e.generic = &handlerHeap[T]{l}
		} else {
			e.generic.Push(l)
		}
	}
}

func (e *Events[T]) HandlePacket(cl *T, p pk.Packet) error {
	if e.generic != nil {
		for _, handler := range *e.generic {
			if err := handler.F(cl, p); err != nil {
				return err
			}
		}
	}

	if h := e.handlers[p.ID]; h != nil {
		for _, handler := range *h {
			if err := handler.F(cl, p); err != nil {
				return err
			}
		}
	}
	return nil
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

func (e *Events[T]) Tick(cl *T) error {
	if e.tickers == nil {
		return nil
	}

	for _, t := range *e.tickers {
		if err := t.F(cl); err != nil {
			return err
		}
	}
	return nil
}

type PacketHandler[T any] struct {
	ID       int32
	Priority int
	F        func(*T, pk.Packet) error
}

// handlerHeap is PriorityQueue<PacketHandlerFunc>
type handlerHeap[T any] []PacketHandler[T]

func (h handlerHeap[T]) Len() int            { return len(h) }
func (h handlerHeap[T]) Less(i, j int) bool  { return h[i].Priority < h[j].Priority }
func (h handlerHeap[T]) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *handlerHeap[T]) Push(x interface{}) { *h = append(*h, x.(PacketHandler[T])) }
func (h *handlerHeap[T]) Pop() interface{} {
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
