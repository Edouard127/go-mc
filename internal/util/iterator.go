package util

type Iterator[T any] interface {
	Next() T
	HasNext() bool
}

type IteratorFunc[T any] func() (T, bool)

func (f IteratorFunc[T]) Next() T {
	v, _ := f()
	return v
}

func (f IteratorFunc[T]) HasNext() bool {
	_, ok := f()
	return ok
}

func quickSort[T any](data []T, comparator func(T, T) int) {
	if len(data) <= 1 {
		return
	}
	pivot := data[len(data)-1]
	i := 0
	for j := 0; j < len(data)-1; j++ {
		if comparator(data[j], pivot) <= 0 {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}
	data[i], data[len(data)-1] = data[len(data)-1], data[i]
	quickSort(data[:i], comparator)
	quickSort(data[i+1:], comparator)
}
