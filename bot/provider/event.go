package provider

import (
	"container/heap"
	"context"
	pk "github.com/Edouard127/go-mc/net/packet"
)

type Events[T any] struct {
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
		heap.Fix(s, s.Len()-1)
	}
}

func (e *Events[T]) HandlePacket(cl *T, p pk.Packet) error {
	if h := e.handlers[p.ID]; h != nil {
		ctx, cancel := context.WithCancel(context.TODO())
		for _, handler := range *h {
			if err := handler.F(cl, p, cancel); err != nil {
				return err
			}
			// if the context is canceled, stop calling the next handler
			if ctx.Err() != nil {
				break
			}
		}
		cancel()
	}
	return nil
}

type TickHandler[T any] struct {
	Priority int
	F        func(*T, context.CancelFunc) error
}

func (e *Events[T]) AddTicker(tickers ...TickHandler[T]) {
	for _, t := range tickers {
		if e.tickers == nil {
			e.tickers = &tickerHeap[T]{t}
		} else {
			e.tickers.Push(t)
		}
	}
	heap.Fix(e.tickers, e.tickers.Len()-1)
}

func (e *Events[T]) Tick(cl *T) error {
	if e.tickers == nil {
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())

	for _, t := range *e.tickers {
		if err := t.F(cl, cancel); err != nil {
			return err
		}
		// if the context is canceled, stop calling the next handler
		if ctx.Err() != nil {
			break
		}
	}
	cancel()
	return nil
}

type PacketHandler[T any] struct {
	ID       int32
	Priority int
	F        func(*T, pk.Packet, context.CancelFunc) error
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
