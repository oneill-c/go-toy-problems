package datastructures

type Queue[T any] struct {
	data []T
}

func (q *Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.data) == 0 {
		var zero T
		return zero, false
	}
	val := q.data[0]
	q.data = q.data[1:]
	return val, true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}