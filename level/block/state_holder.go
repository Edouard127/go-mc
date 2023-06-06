package block

import (
	"github.com/Edouard127/go-mc/internal/data"
	"github.com/Edouard127/go-mc/level/block/states"
)

type StateHolder struct {
	Properties map[states.Property[any]]uint32
	Neighbors  data.HashTable[uint64, uint32, StateID]
}

func NewStateHolder(properties map[states.Property[any]]uint32) *StateHolder {
	return &StateHolder{
		Properties: properties,
		Neighbors:  *data.NewHashTable[uint64, uint32, StateID](),
	}
}

func (s *StateHolder) GetDefaultValue() StateID {
	// Nil property
	// If you make a customized implementation, make sure that the
	// default value of the block is always at the column 0 of the
	// row 0
	return s.Neighbors.Get(0, 0)
}

func (s *StateHolder) GetValue(property states.Property[any], value uint32) StateID {
	var hashcode uint64 = 0
	if property != nil {
		hashcode = property.HashCode()
	}
	return s.Neighbors.Get(hashcode, value)
}

func (s *StateHolder) SetValue(property states.Property[any], value uint32, id StateID) any {
	s.Properties[property] = value
	var hashcode uint64 = 0
	if property != nil {
		hashcode = property.HashCode()
	}
	s.Neighbors.Put(hashcode, value, id)
	return value
}

func (s *StateHolder) HasProperty(property states.Property[any]) bool {
	_, ok := s.Properties[property]
	return ok
}
