package block

import (
	"fmt"
	"github.com/Edouard127/go-mc/level/block/states"
)

// StateHolder will hold the neighbors of the states.
// The states are mapped using a map of state id to map of properties
type StateHolder struct {
	Properties map[states.Property[any]]uint32
	Neighbors  map[StateID]map[states.Property[any]]uint32

	min, max StateID
	current  StateID
}

func NewStateHolder(properties map[states.Property[any]]uint32, min StateID) *StateHolder {
	return &StateHolder{
		Properties: properties,
		Neighbors:  make(map[StateID]map[states.Property[any]]uint32),
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
func (s *StateHolder) GetValue(property states.Property[any]) uint32 {
	return s.Properties[property]
}

// SetValue will set the value of the property and update the current state id
func (s *StateHolder) SetValue(property states.Property[any], value uint32) any {
	if property.CanUpdate(value) {
		s.Properties[property] = value
		//s.update()
		return value
	}

	panic("invalid value for property")
}

// PutNeighbors will add a neighbors for a specific map
func (s *StateHolder) PutNeighbors(state StateID, data map[states.Property[any]]uint32) {
	s.max = max(s.max, state)
	s.Neighbors[state] = data
}

func (s *StateHolder) update() {
	for key, value := range s.Neighbors {
		if fmt.Sprintf("%v", value) == fmt.Sprintf("%v", s.Properties) {
			s.current = key
			return
		}
	}

	panic(fmt.Errorf("invalid properties for current state %d", s.current))
}
