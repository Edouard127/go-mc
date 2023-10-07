package provider

import (
	"context"
	pk "github.com/Edouard127/go-mc/net/packet"
	"sort"
)

type Events[T any] struct {
	handlers map[int32][]PacketHandler[T] // for packets
	tickers  []TickHandler[T]             // for tickers
}

// AddListener adds a listener to the event.
// The listener will be called when the packet with the same ID is received.
// The listener will be called in the order of priority.
// The listeners cannot have multiple same ID.
func (e *Events[T]) AddListener(listeners ...PacketHandler[T]) {
	for _, l := range listeners {
		if e.handlers == nil {
			e.handlers = make(map[int32][]PacketHandler[T])
		}
		if e.handlers[l.ID] == nil {
			e.handlers[l.ID] = []PacketHandler[T]{l}
		} else {
			e.handlers[l.ID] = append(e.handlers[l.ID], l)
			sortPacket[T](e.handlers[l.ID])
		}
	}
}

func (e *Events[T]) HandlePacket(cl *T, p pk.Packet) error {
	if h := e.handlers[p.ID]; h != nil {
		ctx, cancel := context.WithCancel(context.TODO())
		for _, handler := range h {
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

func (e *Events[T]) AddTicker(tickers ...TickHandler[T]) {
	if e.tickers == nil {
		e.tickers = []TickHandler[T]{}
	}
	e.tickers = append(e.tickers, tickers...)

}

func (e *Events[T]) Tick(cl *T) error {
	if e.tickers == nil {
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	for _, t := range e.tickers {
		if err := t.F(cl, cancel); err != nil {
			return err
		}
		// if the context is canceled, stop calling the next handler
		if ctx.Err() != nil {
			break
		}
	}
	sortTick[T](e.tickers)
	cancel()
	return nil
}

type PacketHandler[T any] struct {
	ID       int32
	Priority int
	F        func(*T, pk.Packet, context.CancelFunc) error
}

type TickHandler[T any] struct {
	Priority int
	F        func(*T, context.CancelFunc) error
}

func sortPacket[T any](slice []PacketHandler[T]) {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i].Priority > slice[j].Priority
	})
}

func sortTick[T any](slice []TickHandler[T]) {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i].Priority > slice[j].Priority
	})
}
