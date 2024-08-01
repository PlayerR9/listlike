package stack

import (
	"fmt"
)

// Stacker is an interface that defines methods for a stack data structure.
type Stacker[T any] interface {
	// Push is a method that adds a value of type T to the end of the stack.
	//
	// Parameters:
	//   - value: The value of type T to add to the stack.
	//
	// Returns:
	//   - bool: True if the value was successfully added to the stack, false otherwise.
	Push(value T) bool

	// PushMany is a method that adds multiple values of type T to the end of the stack.
	//
	// Parameters:
	//   - values: The values of type T to add to the stack.
	//
	// Returns:
	//   - int: The number of values that were successfully added to the stack.
	PushMany(values []T) int

	// Pop is a method that pops an element from the stack and returns it.
	//
	// Returns:
	//   - T: The value of type T that was popped.
	//   - bool: True if the value was successfully popped, false otherwise.
	Pop() (T, bool)

	// Peek is a method that returns the value at the front of the stack without removing
	// it.
	//
	// Returns:
	//   - T: The value of type T at the front of the stack.
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
	//   - bool: True if the list is full, false otherwise.
	IsFull() bool

	// Slice is a method that returns a slice of the values in the stack.
	//
	// Returns:
	// 	- []T: A slice of the values in the stack.
	Slice() []T

	fmt.GoStringer
}

// StackNode represents a node in a linked list.
type StackNode[T any] struct {
	// value is the value stored in the node.
	Value T

	// next is a pointer to the next linkedNode in the list.
	next *StackNode[T]
}

// NewStackNode is a constructor function that creates a new linkedNode with the given value.
//
// Parameters:
//   - value: The value to store in the linkedNode.
//
// Returns:
//   - *linkedNode: A pointer to the newly created linkedNode.
func NewStackNode[T any](value T) *StackNode[T] {
	return &StackNode[T]{
		Value: value,
	}
}

// SetNext is a method that sets the next linkedNode in the list.
//
// Parameters:
//   - next: A pointer to the next linkedNode in the list.
func (node *StackNode[T]) SetNext(next *StackNode[T]) {
	node.next = next
}

// Next is a method that returns the next linkedNode in the list.
//
// Returns:
//   - *linkedNode: A pointer to the next linkedNode in the list.
func (node *StackNode[T]) Next() *StackNode[T] {
	return node.next
}
