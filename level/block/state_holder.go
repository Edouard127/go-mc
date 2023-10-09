package block

import (
	"github.com/Edouard127/go-mc/level/block/states"
)

// StateHolder will hold the neighbors of the states.
// The states are mapped using a map of state id to map of properties
type StateHolder struct {
	Properties map[states.Property]int
	Neighbors  map[StateID]map[states.Property]int

	min, max StateID
	current  StateID
}

func NewStateHolder(properties map[states.Property]int, min StateID) *StateHolder {
	return &StateHolder{
		Properties: properties,
		Neighbors:  make(map[StateID]map[states.Property]int),
		min:        min,
	}
}

// Default will return the default state id
func (s *StateHolder) Default() StateID {
	return s.min
}

// State will return the current state id
func (s *StateHolder) State() StateID {
	return s.current
}

// GetValue will return the value of the property
func (s *StateHolder) GetValue(property states.Property) int {
	return s.Properties[property]
}

// SetValue will set the value of the property and update the current state id
func (s *StateHolder) SetValue(property states.Property, value int) {
	s.Properties[property] = value
}

// PutNeighbors will add a neighbors for a specific map
func (s *StateHolder) PutNeighbors(state StateID, data map[states.Property]int) {
	s.max = max(s.max, state)
	s.Neighbors[state] = data
}
