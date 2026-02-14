package main

import (
	"fmt"
)

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)

}
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	item := q.items[0]

	q.items = q.items[1:]
	return item, true
}
func (q *Queue[T]) Peek() T {
	if len(q.items) == 0 {
		var zero T
		return zero
	}
	return q.items[0]
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func main() {
	stringQ := Queue[string]{}
	stringQ.Enqueue("Spark")
	fmt.Println(stringQ.Dequeue())

	intQ := Queue[int]{}
	intQ.Enqueue(4040)
	item, ok := intQ.Dequeue()
	item, ok = intQ.Dequeue()
	if ok {
		fmt.Println(item)
	} else {
		fmt.Println("Queue is empty")
	}
}
