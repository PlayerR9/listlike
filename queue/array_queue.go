package queue

import (
	"strconv"
	"strings"

	gcstr "github.com/PlayerR9/go-commons/strings"
	itrs "github.com/PlayerR9/iterators/simple"
)

// ArrayQueue is a generic type that represents a queue data structure with
// or without a limited capacity. It is implemented using an array.
type ArrayQueue[T any] struct {
	// values is a slice of type T that stores the elements in the queue.
	values []T
}

// Enqueue implements the Queuer interface.
//
// Always returns true.
func (queue *ArrayQueue[T]) Enqueue(value T) bool {
	queue.values = append(queue.values, value)

	return true
}

// EnqueueMany implements the Queuer interface.
//
// Always returns the number of elements enqueued.
func (queue *ArrayQueue[T]) EnqueueMany(values []T) int {
	if len(values) == 0 {
		return 0
	}

	queue.values = append(queue.values, values...)

	return len(values)
}

// Dequeue implements the Queuer interface.
func (queue *ArrayQueue[T]) Dequeue() (T, bool) {
	if len(queue.values) == 0 {
		return *new(T), false
	}

	toRemove := queue.values[0]
	queue.values = queue.values[1:]
	return toRemove, true
}

// Peek implements the Queuer interface.
func (queue *ArrayQueue[T]) Peek() (T, bool) {
	if len(queue.values) == 0 {
		return *new(T), false
	}

	return queue.values[0], true
}

// IsEmpty implements the Queuer interface.
func (queue *ArrayQueue[T]) IsEmpty() bool {
	return len(queue.values) == 0
}

// Size implements the Queuer interface.
func (queue *ArrayQueue[T]) Size() int {
	return len(queue.values)
}

// Iterator implements the Queuer interface.
func (queue *ArrayQueue[T]) Iterator() itrs.Iterater[T] {
	return itrs.NewSimpleIterator(queue.values)
}

// Clear implements the Queuer interface.
func (queue *ArrayQueue[T]) Clear() {
	queue.values = make([]T, 0)
}

// GoString implements the Queuer interface.
func (queue *ArrayQueue[T]) GoString() string {
	values := make([]string, 0, len(queue.values))
	for _, value := range queue.values {
		values = append(values, gcstr.GoStringOf(value))
	}

	var builder strings.Builder

	builder.WriteString("ArrayQueue{size=")
	builder.WriteString(strconv.Itoa(len(queue.values)))
	builder.WriteString(", values=[← ")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString("]}")

	return builder.String()
}

// Slice implements the Queuer interface.
func (queue *ArrayQueue[T]) Slice() []T {
	slice := make([]T, len(queue.values))
	copy(slice, queue.values)

	return slice
}

// Capacity implements the Queuer interface.
func (queue *ArrayQueue[T]) Capacity() int {
	return -1
}

// IsFull implements the Queuer interface.
func (queue *ArrayQueue[T]) IsFull() bool {
	return false
}

// NewArrayQueue is a function that creates and returns a new instance of a
// ArrayQueue.
//
// Returns:
//   - *ArrayQueue[T]: A pointer to the newly created ArrayQueue. Never returns nil.
func NewArrayQueue[T any]() *ArrayQueue[T] {
	return &ArrayQueue[T]{
		values: make([]T, 0),
	}
}

// Copy is a method of the ArrayQueue type. It is used to create a shallow copy
// of the queue.
//
// Returns:
//   - *ArrayQueue[T]: A shallow copy of the queue.
func (queue *ArrayQueue[T]) Copy() *ArrayQueue[T] {
	queue_copy := &ArrayQueue[T]{
		values: make([]T, len(queue.values)),
	}
	copy(queue_copy.values, queue.values)

	return queue_copy
}
