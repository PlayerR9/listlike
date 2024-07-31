package queue

import (
	"strconv"
	"strings"

	uc "github.com/PlayerR9/lib_units/common"
	lustr "github.com/PlayerR9/lib_units/strings"
)

// LimitedArrayQueue is a generic type that represents a queue data structure with
// or without a limited capacity. It is implemented using an array.
type LimitedArrayQueue[T any] struct {
	// values is a slice of type T that stores the elements in the queue.
	values []T

	// capacity is the maximum number of elements the queue can hold.
	capacity int
}

// Enqueue implements the Queuer interface.
func (queue *LimitedArrayQueue[T]) Enqueue(value T) bool {
	if len(queue.values) >= queue.capacity {
		return false
	}

	queue.values = append(queue.values, value)

	return true
}

// EnqueueMany implements the Queuer interface.
func (queue *LimitedArrayQueue[T]) EnqueueMany(values []T) int {
	if len(values) == 0 {
		return 0
	}

	for i, value := range values {
		ok := queue.Enqueue(value)
		if !ok {
			return i
		}
	}

	return len(values)
}

// Dequeue implements the Queuer interface.
func (queue *LimitedArrayQueue[T]) Dequeue() (T, bool) {
	if len(queue.values) == 0 {
		return *new(T), false
	}

	toRemove := queue.values[0]
	queue.values = queue.values[1:]
	return toRemove, true
}

// Peek implements the Queuer interface.
func (queue *LimitedArrayQueue[T]) Peek() (T, bool) {
	if len(queue.values) == 0 {
		return *new(T), false
	}

	return queue.values[0], true
}

// IsEmpty implements the Queuer interface.
func (queue *LimitedArrayQueue[T]) IsEmpty() bool {
	return len(queue.values) == 0
}

// Size implements the Queuer interface.
func (queue *LimitedArrayQueue[T]) Size() int {
	return len(queue.values)
}

// Capacity implements the Queuer interface.
func (queue *LimitedArrayQueue[T]) Capacity() int {
	return queue.capacity
}

// Iterator implements the Queuer interface.
func (queue *LimitedArrayQueue[T]) Iterator() uc.Iterater[T] {
	return uc.NewSimpleIterator(queue.values)
}

// Clear implements the Queuer interface.
func (queue *LimitedArrayQueue[T]) Clear() {
	queue.values = make([]T, 0, queue.capacity)
}

// IsFull implements the Queuer interface.
func (queue *LimitedArrayQueue[T]) IsFull() bool {
	return len(queue.values) >= queue.capacity
}

// GoString implements the fmt.GoStringer interface.
func (queue *LimitedArrayQueue[T]) GoString() string {
	values := make([]string, 0, len(queue.values))
	for _, value := range queue.values {
		values = append(values, lustr.GoStringOf(value))
	}

	var builder strings.Builder

	builder.WriteString("LimitedArrayQueue[capacity=")
	builder.WriteString(strconv.Itoa(queue.capacity))
	builder.WriteString(", size=")
	builder.WriteString(strconv.Itoa(len(queue.values)))
	builder.WriteString(", values=[← ")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString("]]")

	return builder.String()
}

// Slice implements the Queuer interface.
func (queue *LimitedArrayQueue[T]) Slice() []T {
	slice := make([]T, len(queue.values))
	copy(slice, queue.values)

	return slice
}

// Copy implements the Queuer interface.
func (queue *LimitedArrayQueue[T]) Copy() uc.Copier {
	queueCopy := &LimitedArrayQueue[T]{
		values:   make([]T, len(queue.values)),
		capacity: queue.capacity,
	}
	copy(queueCopy.values, queue.values)

	return queueCopy
}

// NewLimitedArrayQueue is a function that creates and returns a new instance of a
// LimitedArrayQueue.
//
// Parameters:
//   - capacity: The maximum number of elements the queue can hold.
//
// Returns:
//   - *LimitedArrayQueue[T]: A pointer to the newly created LimitedArrayQueue.
//   - error: An error of type *common.ErrInvalidParameter if the capacity is less
//     than 0.
func NewLimitedArrayQueue[T any](capacity int) (*LimitedArrayQueue[T], error) {
	if capacity < 0 {
		return nil, uc.NewErrInvalidParameter("capacity", uc.NewErrGTE(0))
	}

	return &LimitedArrayQueue[T]{
		values: make([]T, 0, capacity),
	}, nil
}
