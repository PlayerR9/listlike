package queue

import (
	"fmt"

	uc "github.com/PlayerR9/lib_units/common"
)

// Queuer is an interface that defines methods for a queue data structure.
type Queuer[T any] interface {
	// Enqueue is a method that adds a value of type T to the end of the queue.
	//
	// Parameters:
	//   - value: The value of type T to add to the queue.
	//
	// Returns:
	//   - bool: True if the value was successfully added to the queue, false otherwise.
	Enqueue(value T) bool

	// EnqueueMany is a method that adds multiple values of type T to the end of the queue.
	//
	// Parameters:
	//   - values: The values of type T to add to the queue.
	//
	// Returns:
	//   - int: The number of values successfully added to the queue.
	EnqueueMany(values []T) int

	// Dequeue is a method that dequeues an element from the queue and returns it.
	//
	// Returns:
	//   - T: The value of type T that was dequeued.
	//   - bool: True if the value was successfully dequeued, false otherwise.
	Dequeue() (T, bool)

	// Peek is a method that returns the value at the front of the queue without
	// removing it.
	//
	// Returns:
	//   - T: The value of type T at the front of the queue.
	//   - bool: True if the value was successfully peeked, false otherwise.
	Peek() (T, bool)

	// IsEmpty is a method that checks whether the list is empty.
	//
	// Returns:
	//
	//   - bool: True if the list is empty, false otherwise.
	IsEmpty() bool

	// Size method returns the number of elements currently in the list.
	//
	// Returns:
	//
	//   - int: The number of elements in the list.
	Size() int

	// Clear method is used to remove all elements from the list, making it empty.
	Clear()

	// Capacity is a method that returns the maximum number of elements that the list can hold.
	//
	// Returns:
	//
	//   - int: The maximum number of elements that the list can hold. -1 if there is no limit.
	Capacity() int

	// IsFull is a method that checks whether the list is full.
	//
	// Returns:
	//
	//   - bool: True if the list is full, false otherwise.
	IsFull() bool

	// Slice is a method that returns a slice of the elements in the list.
	//
	// Returns:
	//  	- []T: A slice of the elements in the list.
	Slice() []T

	uc.Iterable[T]

	uc.Copier
	fmt.GoStringer
}

// queue_node represents a node in a linked queue.
type queue_node[T any] struct {
	// value is the value stored in the node.
	value T

	// next is a pointer to the next linkedNode in the queue.
	next *queue_node[T]
}

// queue_safe_node represents a node in a linked queue.
type queue_safe_node[T any] struct {
	// value is the value stored in the node.
	value T

	// next is a pointer to the next queueLinkedNode in the queue.
	next *queue_safe_node[T]
}
