package stack

import (
	"strconv"
	"strings"

	uc "github.com/PlayerR9/lib_units/common"
	lustr "github.com/PlayerR9/lib_units/strings"
)

// LimitedLinkedStack is a generic type that represents a stack data structure with
// or without a limited capacity, implemented using a linked list.
type LimitedLinkedStack[T any] struct {
	// front is a pointer to the first node in the linked stack.
	front *StackNode[T]

	// size is the current number of elements in the stack.
	size int

	// capacity is the maximum number of elements the stack can hold.
	capacity int
}

// NewLimitedLinkedStack is a function that creates and returns a new instance of a
// LimitedLinkedStack.
//
// Parameters:
//
//   - values: A variadic parameter of type T, which represents the initial values to be
//     stored in the stack.
//
// Returns:
//
//   - *LimitedLinkedStack[T]: A pointer to the newly created LimitedLinkedStack.
func NewLimitedLinkedStack[T any](values ...T) *LimitedLinkedStack[T] {
	stack := new(LimitedLinkedStack[T])
	stack.size = len(values)

	if len(values) == 0 {
		return stack
	}

	// First node
	node := NewStackNode(values[0])

	stack.front = node

	// Subsequent nodes
	for _, element := range values[1:] {
		node = NewStackNode(element)
		node.SetNext(stack.front)

		stack.front = node
	}

	return stack
}

// Push implements the Stacker interface.
func (stack *LimitedLinkedStack[T]) Push(value T) bool {
	if stack.size >= stack.capacity {
		return false
	}

	node := NewStackNode(value)

	if stack.front != nil {
		node.SetNext(stack.front)
	}

	stack.front = node
	stack.size++

	return true
}

// PushMany implements the Stacker interface.
func (stack *LimitedLinkedStack[T]) PushMany(values []T) int {
	if stack.size+len(values) > stack.capacity {
		return 0
	}

	for _, value := range values {
		stack.Push(value)
	}

	return len(values)
}

// Pop implements the Stacker interface.
func (stack *LimitedLinkedStack[T]) Pop() (T, bool) {
	if stack.front == nil {
		return *new(T), false
	}

	toRemove := stack.front
	stack.front = stack.front.Next()

	stack.size--
	toRemove.SetNext(nil)

	return toRemove.Value, true
}

// Peek implements the Stacker interface.
func (stack *LimitedLinkedStack[T]) Peek() (T, bool) {
	if stack.front == nil {
		return *new(T), false
	}

	return stack.front.Value, true
}

// IsEmpty is a method of the LimitedLinkedStack type. It is used to check if the stack
// is empty.
//
// Returns:
//
//   - bool: true if the stack is empty, and false otherwise.
func (stack *LimitedLinkedStack[T]) IsEmpty() bool {
	return stack.front == nil
}

// Size is a method of the LimitedLinkedStack type. It is used to return the number of
// elements in the stack.
//
// Returns:
//
//   - int: The number of elements in the stack.
func (stack *LimitedLinkedStack[T]) Size() int {
	return stack.size
}

// Capacity is a method of the LimitedLinkedStack type. It is used to return the maximum
// number of elements the stack can hold.
//
// Returns:
//
//   - optional.Int: The maximum number of elements the stack can hold.
func (stack *LimitedLinkedStack[T]) Capacity() int {
	return stack.capacity
}

// Iterator is a method of the LimitedLinkedStack type. It is used to return an iterator
// for the elements in the stack.
//
// Returns:
//
//   - uc.Iterater[T]: An iterator for the elements in the stack.
func (stack *LimitedLinkedStack[T]) Iterator() uc.Iterater[T] {
	var builder uc.Builder[T]

	for stack_node := stack.front; stack_node != nil; stack_node = stack_node.Next() {
		builder.Add(stack_node.Value)
	}

	return builder.Build()
}

// Clear is a method of the LimitedLinkedStack type. It is used to remove all elements
// from the stack.
func (stack *LimitedLinkedStack[T]) Clear() {
	if stack.front == nil {
		return // Stack is already empty
	}

	// 1. First node
	prev := stack.front

	// 2. Subsequent nodes
	for node := stack.front.Next(); node != nil; node = node.Next() {
		prev = node
		prev.SetNext(nil)
	}

	prev.SetNext(nil)

	// 3. Reset list fields
	stack.front = nil
	stack.size = 0
}

// IsFull is a method of the LimitedLinkedStack type. It is used to check if the stack is
// full.
//
// Returns:
//
//   - isFull: true if the stack is full, and false otherwise.
func (stack *LimitedLinkedStack[T]) IsFull() bool {
	return stack.size >= stack.capacity
}

// GoString implements the fmt.GoStringer interface.
func (stack *LimitedLinkedStack[T]) GoString() string {
	values := make([]string, 0, stack.size)
	for stack_node := stack.front; stack_node != nil; stack_node = stack_node.Next() {
		values = append(values, lustr.GoStringOf(stack_node.Value))
	}

	var builder strings.Builder

	builder.WriteString("LimitedLinkedStack[capacity=")
	builder.WriteString(strconv.Itoa(stack.capacity))
	builder.WriteString(", size=")
	builder.WriteString(strconv.Itoa(stack.size))
	builder.WriteString(", values=[")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString(" →]]")

	return builder.String()
}

// Slice is a method of the LimitedLinkedStack type. It is used to return a slice of the
// elements in the stack.
//
// Returns:
//
//   - []T: A slice of the elements in the stack.
func (stack *LimitedLinkedStack[T]) Slice() []T {
	slice := make([]T, 0, stack.size)

	for stack_node := stack.front; stack_node != nil; stack_node = stack_node.Next() {
		slice = append(slice, stack_node.Value)
	}

	return slice
}

// Copy is a method of the LimitedLinkedStack type. It is used to create a shallow copy
// of the stack.
//
// Returns:
//   - *LimitedLinkedStack[T]: A copy of the stack.
func (stack *LimitedLinkedStack[T]) Copy() *LimitedLinkedStack[T] {
	stackCopy := &LimitedLinkedStack[T]{
		size:     stack.size,
		capacity: stack.capacity,
	}

	if stack.front == nil {
		return stackCopy
	}

	// First node
	node := NewStackNode(stack.front.Value)

	stackCopy.front = node

	prev := stackCopy.front

	// Subsequent nodes
	for stack_node := stack.front.Next(); stack_node != nil; stack_node = stack_node.Next() {
		node := NewStackNode(stack_node.Value)
		prev.SetNext(node)

		prev = node
	}

	return stackCopy
}
