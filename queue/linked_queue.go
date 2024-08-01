package queue

import (
	"strconv"
	"strings"

	uc "github.com/PlayerR9/lib_units/common"
	lustr "github.com/PlayerR9/lib_units/strings"
)

// LinkedQueue is a generic type that represents a queue data structure with
// or without a limited capacity, implemented using a linked list.
type LinkedQueue[T any] struct {
	// front and back are pointers to the first and last nodes in the linked queue,
	// respectively.
	front, back *queue_node[T]

	// size is the current number of elements in the queue.
	size int
}

// Enqueue implements the Queuer interface.
//
// Always returns true.
func (queue *LinkedQueue[T]) Enqueue(value T) bool {
	node := &queue_node[T]{
		value: value,
	}

	if queue.back == nil {
		queue.front = node
	} else {
		queue.back.next = node
	}

	queue.back = node

	queue.size++

	return true
}

// EnqueueMany implements the Queuer interface.
//
// Always returns true.
func (queue *LinkedQueue[T]) EnqueueMany(values []T) int {
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
func (queue *LinkedQueue[T]) Dequeue() (T, bool) {
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
func (queue *LinkedQueue[T]) Peek() (T, bool) {
	if queue.front == nil {
		return *new(T), false
	}

	return queue.front.value, true
}

// IsEmpty implements the Queuer interface.
func (queue *LinkedQueue[T]) IsEmpty() bool {
	return queue.front == nil
}

// Size implements the Queuer interface.
func (queue *LinkedQueue[T]) Size() int {
	return queue.size
}

// Iterator implements the Queuer interface.
func (queue *LinkedQueue[T]) Iterator() uc.Iterater[T] {
	var builder uc.Builder[T]

	for queue_node := queue.front; queue_node != nil; queue_node = queue_node.next {
		builder.Add(queue_node.value)
	}

	return builder.Build()
}

// Clear implements the Queuer interface.
func (queue *LinkedQueue[T]) Clear() {
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

// GoString implements the fmt.GoStringer interface.
func (queue *LinkedQueue[T]) GoString() string {
	values := make([]string, 0, queue.size)
	for queue_node := queue.front; queue_node != nil; queue_node = queue_node.next {
		values = append(values, lustr.GoStringOf(queue_node.value))
	}

	var builder strings.Builder

	builder.WriteString("LinkedQueue{size=")
	builder.WriteString(strconv.Itoa(queue.size))
	builder.WriteString(", values=[← ")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString("]}")

	return builder.String()
}

// Slice implements the Queuer interface.
func (queue *LinkedQueue[T]) Slice() []T {
	slice := make([]T, 0, queue.size)

	for queue_node := queue.front; queue_node != nil; queue_node = queue_node.next {
		slice = append(slice, queue_node.value)
	}

	return slice
}

// Capacity implements the Queuer interface.
//
// Always returns -1.
func (queue *LinkedQueue[T]) Capacity() int {
	return -1
}

// IsFull implements the Queuer interface.
//
// Always returns false.
func (queue *LinkedQueue[T]) IsFull() bool {
	return false
}

// NewLinkedQueue is a function that creates and returns a new instance of a
// LinkedQueue.
//
// Returns:
//   - *LinkedQueue[T]: A pointer to the newly created LinkedQueue. Never returns nil.
func NewLinkedQueue[T any]() *LinkedQueue[T] {
	return &LinkedQueue[T]{
		size: 0,
	}
}

// Copy is a method that returns a copy of the LinkedQueue.
//
// Returns:
//   - *LinkedQueue[T]: A copy of the LinkedQueue.
func (queue *LinkedQueue[T]) Copy() *LinkedQueue[T] {
	queue_copy := &LinkedQueue[T]{
		size: queue.size,
	}

	if queue.size == 0 {
		return queue_copy
	}

	// First node
	node := &queue_node[T]{
		value: queue.front.value,
	}

	queue_copy.front = node
	queue_copy.back = node

	// Subsequent nodes
	for n := queue.front.next; n != nil; n = n.next {
		node := &queue_node[T]{
			value: n.value,
		}

		queue_copy.back.next = node
		queue_copy.back = node
	}

	return queue_copy
}
