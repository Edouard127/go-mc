package data

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type HashTable[T, K, V constraints.Integer] struct {
	data map[T]map[K]V
}

// NewHashTable returns a new HashTable[T, K, V]
// We allow the constraint Integer because you can have a negative hash
func NewHashTable[T, K, V constraints.Integer]() *HashTable[T, K, V] {
	return &HashTable[T, K, V]{make(map[T]map[K]V)}
}

func (h *HashTable[T, K, V]) Put(rowKey T, columnKey K, value V) {
	if h.data[rowKey] == nil {
		h.data[rowKey] = make(map[K]V)
	}
	h.data[rowKey][columnKey] = value
}

func (h *HashTable[T, K, V]) Get(rowKey T, columnKey K) V {
	row, ok := h.data[rowKey]
	if !ok {
		fmt.Println(fmt.Errorf("rowKey %v not found", rowKey))
		return 0
	}
	value, _ := row[columnKey]
	return value
}
