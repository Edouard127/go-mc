package block

import (
	"github.com/Edouard127/go-mc/level/block/states"
)

type StateFeeder[T any] struct {
	index int
	max   int
}

func NewStateFeeder[T any]() *StateFeeder[T] {
	return &StateFeeder[T]{0, 0}
}

func (s *StateFeeder[T]) FeedState(state *StateHolder, properties []states.Property[any]) *StateHolder {
	length := 1

	for _, p := range properties {
		length *= len(p.GetValues())
	}

	s.max += length

	state.SetValue(nil, 0, StateID(s.index)) // default state
	s.index++

	if len(properties) > 0 {
		for _, p := range properties {
			for _, v := range p.GetValues() {
				state.SetValue(p, parseState(v), StateID(s.index))
				s.index++
			}
		}
	}

	return state
}
