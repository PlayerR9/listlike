package queue

import (
	"strconv"
	"strings"
	"sync"

	gcers "github.com/PlayerR9/go-commons/errors"
	gcstr "github.com/PlayerR9/go-commons/strings"
	uc "github.com/PlayerR9/lib_units/common"
)

// LimitedSafeQueue is a generic type that represents a thread-safe queue data
// structure with or without a limited capacity, implemented using a linked list.
type LimitedSafeQueue[T any] struct {
	// front and back are pointers to the first and last nodes in the safe queue,
	// respectively.
	front, back *queue_safe_node[T]

	// frontMutex and backMutex are sync.RWMutexes, which are used to ensure that
	// concurrent reads and writes to the front and back nodes are thread-safe.
	frontMutex, backMutex sync.RWMutex

	// size is the current number of elements in the queue.
	size int

	// capacity is the maximum number of elements that the queue can hold.
	capacity int
}

// Enqueue implements the Queuer interface.
func (queue *LimitedSafeQueue[T]) Enqueue(value T) bool {
	queue.backMutex.Lock()
	defer queue.backMutex.Unlock()

	if queue.size >= queue.capacity {
		return false
	}

	node := &queue_safe_node[T]{
		value: value,
	}

	if queue.back == nil {
		queue.frontMutex.Lock()
		queue.front = node
		queue.frontMutex.Unlock()
	} else {
		queue.back.next = node
	}

	queue.back = node
	queue.size++

	return true
}

// EnqueueMany implements the Queuer interface.
func (queue *LimitedSafeQueue[T]) EnqueueMany(values []T) int {
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
func (queue *LimitedSafeQueue[T]) Dequeue() (T, bool) {
	queue.frontMutex.Lock()
	defer queue.frontMutex.Unlock()

	if queue.front == nil {
		return *new(T), false
	}

	toRemove := queue.front

	if queue.front.next == nil {
		queue.front = nil

		queue.backMutex.Lock()
		queue.back = nil
		queue.backMutex.Unlock()
	} else {
		queue.front = queue.front.next
	}

	queue.size--
	toRemove.next = nil

	return toRemove.value, true
}

// Peek implements the Queuer interface.
func (queue *LimitedSafeQueue[T]) Peek() (T, bool) {
	queue.frontMutex.RLock()
	defer queue.frontMutex.RUnlock()

	if queue.front == nil {
		return *new(T), false
	}

	return queue.front.value, true
}

// IsEmpty implements the Queuer interface.
func (queue *LimitedSafeQueue[T]) IsEmpty() bool {
	queue.frontMutex.RLock()
	defer queue.frontMutex.RUnlock()

	return queue.front == nil
}

// Size implements the Queuer interface.
func (queue *LimitedSafeQueue[T]) Size() int {
	queue.frontMutex.RLock()
	defer queue.frontMutex.RUnlock()

	queue.backMutex.RLock()
	defer queue.backMutex.RUnlock()

	return queue.size
}

// Capacity implements the Queuer interface.
func (queue *LimitedSafeQueue[T]) Capacity() int {
	return queue.capacity
}

// Iterator implements the Queuer interface.
func (queue *LimitedSafeQueue[T]) Iterator() uc.Iterater[T] {
	queue.frontMutex.RLock()
	defer queue.frontMutex.RUnlock()

	queue.backMutex.RLock()
	defer queue.backMutex.RUnlock()

	var builder uc.Builder[T]

	for node := queue.front; node != nil; node = node.next {
		builder.Add(node.value)
	}

	return builder.Build()
}

// Clear implements the Queuer interface.
func (queue *LimitedSafeQueue[T]) Clear() {
	queue.frontMutex.Lock()
	defer queue.frontMutex.Unlock()

	queue.backMutex.Lock()
	defer queue.backMutex.Unlock()

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
func (queue *LimitedSafeQueue[T]) IsFull() (isFull bool) {
	queue.backMutex.RLock()
	defer queue.backMutex.RUnlock()

	return queue.size >= queue.capacity
}

// GoString implements the fmt.GoStringer interface.
func (queue *LimitedSafeQueue[T]) GoString() string {
	queue.frontMutex.RLock()
	defer queue.frontMutex.RUnlock()

	queue.backMutex.RLock()
	defer queue.backMutex.RUnlock()

	values := make([]string, 0, queue.size)
	for node := queue.front; node != nil; node = node.next {
		values = append(values, gcstr.GoStringOf(node.value))
	}

	var builder strings.Builder

	builder.WriteString("LimitedSafeQueue[capacity=")
	builder.WriteString(strconv.Itoa(queue.capacity))
	builder.WriteString(", size=")
	builder.WriteString(strconv.Itoa(queue.size))
	builder.WriteString(", values=[← ")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString("]]")

	return builder.String()
}

// Slice implements the Queuer interface.
func (queue *LimitedSafeQueue[T]) Slice() []T {
	queue.frontMutex.RLock()
	defer queue.frontMutex.RUnlock()

	queue.backMutex.RLock()
	defer queue.backMutex.RUnlock()

	slice := make([]T, 0, queue.size)

	for node := queue.front; node != nil; node = node.next {
		slice = append(slice, node.value)
	}

	return slice
}

// NewLimitedSafeQueue is a function that creates and returns a new instance of a
// LimitedSafeQueue.
//
// Parameters:
//   - capacity: The capacity of the queue.
//
// Return:
//   - *LimitedSafeQueue[T]: A pointer to the newly created LimitedSafeQueue.
func NewLimitedSafeQueue[T any](capacity int) (*LimitedSafeQueue[T], error) {
	if capacity < 0 {
		return nil, gcers.NewErrInvalidParameter("capacity", uc.NewErrGTE(0))
	}

	return &LimitedSafeQueue[T]{
		capacity: capacity,
	}, nil
}

// Copy is a method of the LimitedSafeQueue type. It is used to create a shallow
// copy of the queue.
//
// Returns:
//   - *LimitedSafeQueue[T]: A shallow copy of the queue.
func (queue *LimitedSafeQueue[T]) Copy() *LimitedSafeQueue[T] {
	queue.frontMutex.RLock()
	defer queue.frontMutex.RUnlock()

	queue.backMutex.RLock()
	defer queue.backMutex.RUnlock()

	queue_copy := &LimitedSafeQueue[T]{
		size: queue.size,
	}

	if queue.front == nil {
		return queue_copy
	}

	// First node
	node := &queue_safe_node[T]{
		value: queue.front.value,
	}

	queue_copy.front = node
	queue_copy.back = node

	// Subsequent nodes
	for qNode := queue.front.next; qNode != nil; qNode = qNode.next {
		node := &queue_safe_node[T]{
			value: qNode.value,
		}

		queue_copy.back.next = node
		queue_copy.back = node
	}

	return queue_copy
}
