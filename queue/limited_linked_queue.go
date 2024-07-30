package queue

import (
	"strconv"
	"strings"

	uc "github.com/PlayerR9/lib_units/common"
)

// LimitedLinkedQueue is a generic type that represents a queue data structure with
// or without a limited capacity, implemented using a linked list.
type LimitedLinkedQueue[T any] struct {
	// front and back are pointers to the first and last nodes in the linked queue,
	// respectively.
	front, back *queue_node[T]

	// size is the current number of elements in the queue.
	size int

	// capacity is the maximum number of elements the queue can hold.
	capacity int
}

// Enqueue implements the Queuer interface.
func (queue *LimitedLinkedQueue[T]) Enqueue(value T) bool {
	if queue.size >= queue.capacity {
		return false
	}

	queue_node := &queue_node[T]{
		value: value,
	}

	if queue.back == nil {
		queue.front = queue_node
	} else {
		queue.back.next = queue_node
	}

	queue.back = queue_node

	queue.size++

	return true
}

// EnqueueMany implements the Queuer interface.
func (queue *LimitedLinkedQueue[T]) EnqueueMany(values []T) int {
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
func (queue *LimitedLinkedQueue[T]) Dequeue() (T, bool) {
	if queue.front == nil {
		return *new(T), false
	}

	toRemove := queue.front

	queue.front = queue.front.next
	if queue.front == nil {
		queue.back = nil
	}

	queue.size--
	toRemove.next = nil

	return toRemove.value, true
}

// Peek implements the Queuer interface.
func (queue *LimitedLinkedQueue[T]) Peek() (T, bool) {
	if queue.front == nil {
		return *new(T), false
	}

	return queue.front.value, true
}

// IsEmpty implements the Queuer interface.
func (queue *LimitedLinkedQueue[T]) IsEmpty() bool {
	return queue.front == nil
}

// Size implements the Queuer interface.
func (queue *LimitedLinkedQueue[T]) Size() int {
	return queue.size
}

// Capacity implements the Queuer interface.
func (queue *LimitedLinkedQueue[T]) Capacity() int {
	return queue.capacity
}

// Iterator implements the Queuer interface.
func (queue *LimitedLinkedQueue[T]) Iterator() uc.Iterater[T] {
	var builder uc.Builder[T]

	for queue_node := queue.front; queue_node != nil; queue_node = queue_node.next {
		builder.Add(queue_node.value)
	}

	return builder.Build()
}

// Clear implements the Queuer interface.
func (queue *LimitedLinkedQueue[T]) Clear() {
	if queue.front == nil {
		return // Queue is already empty
	}

	// 1. First node
	prev := queue.front

	// 2. Subsequent nodes
	for node := queue.front.next; node != nil; node = node.next {
		prev = node
		prev.next = nil
	}

	prev.next = nil

	// 3. Reset queue fields
	queue.front = nil
	queue.back = nil
	queue.size = 0
}

// IsFull implements the Queuer interface.
func (queue *LimitedLinkedQueue[T]) IsFull() bool {
	return queue.size >= queue.capacity
}

// GoString implements the fmt.GoStringer interface.
func (queue *LimitedLinkedQueue[T]) GoString() string {
	values := make([]string, 0, queue.size)
	for queue_node := queue.front; queue_node != nil; queue_node = queue_node.next {
		values = append(values, uc.StringOf(queue_node.value))
	}

	var builder strings.Builder

	builder.WriteString("LimitedLinkedQueue[capacity=")
	builder.WriteString(strconv.Itoa(queue.capacity))
	builder.WriteString(", size=")
	builder.WriteString(strconv.Itoa(queue.size))
	builder.WriteString(", values=[← ")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString("]]")

	return builder.String()
}

// Slice implements the Queuer interface.
func (queue *LimitedLinkedQueue[T]) Slice() []T {
	slice := make([]T, 0, queue.size)

	for queue_node := queue.front; queue_node != nil; queue_node = queue_node.next {
		slice = append(slice, queue_node.value)
	}

	return slice
}

// Copy implements the Queuer interface.
func (queue *LimitedLinkedQueue[T]) Copy() uc.Copier {
	queueCopy := &LimitedLinkedQueue[T]{
		size:     queue.size,
		capacity: queue.capacity,
	}

	if queue.size == 0 {
		return queueCopy
	}

	// First node
	node := &queue_node[T]{
		value: queue.front.value,
	}

	queueCopy.front = node
	queueCopy.back = node

	// Subsequent nodes
	for n := queue.front.next; n != nil; n = n.next {
		node := &queue_node[T]{
			value: n.value,
		}

		queueCopy.back.next = node
		queueCopy.back = node
	}

	return queueCopy
}

// NewLimitedLinkedQueue is a function that creates and returns a new instance of a
// LimitedLinkedQueue.
//
// Parameters:
//   - capacity: The maximum number of elements the queue can hold.
//
// Returns:
//   - *LimitedLinkedQueue[T]: A pointer to the newly created LimitedLinkedQueue.
//   - error: An error of type *common.ErrInvalidParameter if the capacity is less
//     than 0.
func NewLimitedLinkedQueue[T any](capacity int) (*LimitedLinkedQueue[T], error) {
	if capacity < 0 {
		return nil, uc.NewErrInvalidParameter("capacity", uc.NewErrGTE(0))
	}

	return &LimitedLinkedQueue[T]{
		capacity: capacity,
	}, nil
}
