package util

type Queue[T any] []T

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Push(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Pop() (v T) {
	h := *q
	l := len(h)
	if l == 0 {
		return
	}
	v, *q = h[0], h[1:l]
	return
}

func (q *Queue[T]) Len() int {
	return len(*q)
}
