package queue

import (
	"strconv"
	"strings"
	"sync"

	uc "github.com/PlayerR9/lib_units/common"
)

// SafeQueue is a generic type that represents a thread-safe queue data
// structure with or without a limited capacity, implemented using a linked list.
type SafeQueue[T any] struct {
	// front and back are pointers to the first and last nodes in the safe queue,
	// respectively.
	front, back *queue_safe_node[T]

	// frontMutex and backMutex are sync.RWMutexes, which are used to ensure that
	// concurrent reads and writes to the front and back nodes are thread-safe.
	mu sync.RWMutex

	// size is the size that observers observe.
	size int
}

// NewSafeQueue is a function that creates and returns a new instance of a
// SafeQueue.
//
// Return:
//   - *SafeQueue[T]: A pointer to the newly created SafeQueue. Never returns nil.
func NewSafeQueue[T any]() *SafeQueue[T] {
	return &SafeQueue[T]{
		size: 0,
	}
}

// Enqueue implements the Queuer interface.
//
// Always returns true.
func (queue *SafeQueue[T]) Enqueue(value T) bool {
	queue.mu.Lock()
	defer queue.mu.Unlock()

	node := &queue_safe_node[T]{
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

// Enqueue implements the Queuer interface.
//
// Always returns true.
func (queue *SafeQueue[T]) EnqueueMany(values []T) int {
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
func (queue *SafeQueue[T]) Dequeue() (T, bool) {
	queue.mu.Lock()
	defer queue.mu.Unlock()

	if queue.front == nil {
		return *new(T), false
	}

	toRemove := queue.front

	if queue.front.next == nil {
		queue.front = nil
		queue.back = nil
	} else {
		queue.front = queue.front.next
	}

	queue.size--

	return toRemove.value, true
}

// Peek implements the Queuer interface.
func (queue *SafeQueue[T]) Peek() (T, bool) {
	queue.mu.RLock()
	defer queue.mu.RUnlock()

	if queue.front == nil {
		return *new(T), false
	}

	return queue.front.value, true
}

// IsEmpty implements the Queuer interface.
func (queue *SafeQueue[T]) IsEmpty() bool {
	queue.mu.RLock()
	defer queue.mu.RUnlock()

	return queue.front == nil
}

// Size implements the Queuer interface.
func (queue *SafeQueue[T]) Size() int {
	queue.mu.RLock()
	defer queue.mu.RUnlock()

	return queue.size
}

// Iterator implements the Queuer interface.
func (queue *SafeQueue[T]) Iterator() uc.Iterater[T] {
	queue.mu.RLock()
	defer queue.mu.RUnlock()

	var builder uc.Builder[T]

	for node := queue.front; node != nil; node = node.next {
		builder.Add(node.value)
	}

	return builder.Build()
}

// Clear implements the Queuer interface.
func (queue *SafeQueue[T]) Clear() {
	queue.mu.Lock()
	defer queue.mu.Unlock()

	if queue.front == nil {
		return // Queue is already empty
	}

	queue.front = nil
	queue.back = nil

	queue.size = 0
}

// GoString implements the fmt.GoStringer interface.
func (queue *SafeQueue[T]) GoString() string {
	queue.mu.RLock()
	defer queue.mu.RUnlock()

	size := queue.size

	values := make([]string, 0, size)
	for node := queue.front; node != nil; node = node.next {
		values = append(values, uc.StringOf(node.value))
	}

	var builder strings.Builder

	builder.WriteString("SafeQueue{size=")
	builder.WriteString(strconv.Itoa(size))
	builder.WriteString(", values=[← ")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString("]}")

	return builder.String()
}

// Slice implements the Queuer interface.
func (queue *SafeQueue[T]) Slice() []T {
	queue.mu.RLock()
	defer queue.mu.RUnlock()

	slice := make([]T, 0, queue.size)

	for node := queue.front; node != nil; node = node.next {
		slice = append(slice, node.value)
	}

	return slice
}

// Copy implements the Queuer interface.
func (queue *SafeQueue[T]) Copy() uc.Copier {
	queue.mu.RLock()
	defer queue.mu.RUnlock()

	queueCopy := &SafeQueue[T]{
		size: queue.size,
	}

	if queue.front == nil {
		return queueCopy
	}

	// First node
	node := &queue_safe_node[T]{
		value: queue.front.value,
	}

	queueCopy.front = node
	queueCopy.back = node

	// Subsequent nodes
	for qNode := queue.front.next; qNode != nil; qNode = qNode.next {
		node = &queue_safe_node[T]{
			value: qNode.value,
		}

		queueCopy.back.next = node
		queueCopy.back = node
	}

	return queueCopy
}

// Capacity implements the Queuer interface.
//
// Always returns -1.
func (queue *SafeQueue[T]) Capacity() int {
	return -1
}

// IsFull implements the Queuer interface.
//
// Always returns false.
func (queue *SafeQueue[T]) IsFull() bool {
	return false
}
