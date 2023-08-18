package util

import (
	"fmt"
	"sync"
)

type StateTable[T, K comparable, V any] struct {
	data      map[T]map[K]V
	dataMutex sync.Mutex
}

func NewStateTable[T, K comparable, V any]() *StateTable[T, K, V] {
	return &StateTable[T, K, V]{data: make(map[T]map[K]V)}
}

func (h *StateTable[T, K, V]) PutRow(rowKey T, row map[K]V) {
	h.dataMutex.Lock()
	defer h.dataMutex.Unlock()
	h.data[rowKey] = row
}

func (h *StateTable[T, K, V]) PutAt(rowKey T, columnKey K, value V) {
	h.dataMutex.Lock()
	defer h.dataMutex.Unlock()
	if h.data[rowKey] == nil {
		h.data[rowKey] = make(map[K]V)
	}
	h.data[rowKey][columnKey] = value
}

func (h *StateTable[T, K, V]) GetColumn(columnKey T) map[K]V {
	h.dataMutex.Lock()
	defer h.dataMutex.Unlock()
	row, ok := h.data[columnKey]
	if !ok {
		panic(fmt.Errorf("column %v not found", columnKey))
	}
	return row
}

func (h *StateTable[T, K, V]) GetAt(columnKey T, rowKey K) V {
	h.dataMutex.Lock()
	defer h.dataMutex.Unlock()
	row, ok := h.data[columnKey]
	if !ok {
		panic(fmt.Errorf("column %v not found", columnKey))
	}
	value, _ := row[rowKey]
	return value
}

func (h *StateTable[T, K, V]) DeleteColumn(columnKey T) {
	h.dataMutex.Lock()
	defer h.dataMutex.Unlock()
	delete(h.data, columnKey)
}

func (h *StateTable[T, K, V]) Iterator() Iterator[map[K]V] {
	h.dataMutex.Lock()
	defer h.dataMutex.Unlock()
	keys := make([]T, 0, len(h.data))
	for key := range h.data {
		keys = append(keys, key)
	}
	return IteratorFunc[map[K]V](func() (map[K]V, bool) {
		if len(keys) == 0 {
			return nil, false
		}
		key := keys[0]
		keys = keys[1:]
		return h.data[key], true
	})
}
