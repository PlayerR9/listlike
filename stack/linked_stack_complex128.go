// Code generated with go generate. DO NOT EDIT.
package stack

import (
	"github.com/PlayerR9/iterators/simple"
	"strconv"
	"strings"
)

// stack_node_complex128 is a node in the linked stack.
type stack_node_complex128 struct {
	value complex128
	next *stack_node_complex128
}

// Complex128Stack is a stack of complex128 values implemented without a maximum capacity
// and using a linked list.
type Complex128Stack struct {
	front *stack_node_complex128
	size int
}

// NewComplex128Stack creates a new linked stack.
//
// Returns:
//   - *Complex128Stack: A pointer to the newly created stack. Never returns nil.
func NewComplex128Stack() *Complex128Stack {
	return &Complex128Stack{
		size: 0,
	}
}

// Push implements the stack.Stacker interface.
//
// Always returns true.
func (s *Complex128Stack) Push(value complex128) bool {
	node := &stack_node_complex128{
		value: value,
	}

	if s.front != nil {
		node.next = s.front
	}

	s.front = node
	s.size++

	return true
}

// PushMany implements the stack.Stacker interface.
//
// Always returns the number of values pushed onto the stack.
func (s *Complex128Stack) PushMany(values []complex128) int {
	if len(values) == 0 {
		return 0
	}

	node := &stack_node_complex128{
		value: values[0],
	}

	if s.front != nil {
		node.next = s.front
	}

	s.front = node

	for i := 1; i < len(values); i++ {
		node := &stack_node_complex128{
			value: values[i],
			next:  s.front,
		}

		s.front = node
	}

	s.size += len(values)
	
	return len(values)
}

// Pop implements the stack.Stacker interface.
func (s *Complex128Stack) Pop() (complex128, bool) {
	if s.front == nil {
		return 0, false
	}

	to_remove := s.front
	s.front = s.front.next

	s.size--
	to_remove.next = nil

	return to_remove.value, true
}

// Peek implements the stack.Stacker interface.
func (s *Complex128Stack) Peek() (complex128, bool) {
	if s.front == nil {
		return 0, false
	}

	return s.front.value, true
}

// IsEmpty implements the stack.Stacker interface.
func (s *Complex128Stack) IsEmpty() bool {
	return s.front == nil
}

// Size implements the stack.Stacker interface.
func (s *Complex128Stack) Size() int {
	return s.size
}

// Iterator implements the stack.Stacker interface.
func (s *Complex128Stack) Iterator() simple.Iterater[complex128] {
	var builder simple.Builder[complex128]

	for node := s.front; node != nil; node = node.next {
		builder.Add(node.value)
	}

	return builder.Build()
}

// Clear implements the stack.Stacker interface.
func (s *Complex128Stack) Clear() {
	if s.front == nil {
		return
	}

	prev := s.front

	for node := s.front.next; node != nil; node = node.next {
		prev = node
		prev.next = nil
	}

	prev.next = nil

	s.front = nil
	s.size = 0
}

// GoString implements the stack.Stacker interface.
func (s *Complex128Stack) GoString() string {
	values := make([]string, 0, s.size)
	for node := s.front; node != nil; node = node.next {
		values = append(values, strconv.FormatComplex(node.value, 'f', -1, 128))
	}

	var builder strings.Builder

	builder.WriteString("Complex128Stack[size=")
	builder.WriteString(strconv.Itoa(s.size))
	builder.WriteString(", values=[")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString(" →]]")

	return builder.String()
}

// Slice implements the stack.Stacker interface.
//
// The 0th element is the top of the stack.
func (s *Complex128Stack) Slice() []complex128 {
	slice := make([]complex128, 0, s.size)

	for node := s.front; node != nil; node = node.next {
		slice = append(slice, node.value)
	}

	return slice
}

// Capacity implements the stack.Stacker interface.
//
// Always returns -1.
func (s *Complex128Stack) Capacity() int {
	return -1
}

// IsFull implements the stack.Stacker interface.
//
// Always returns false.
func (s *Complex128Stack) IsFull() bool {
	return false
}

// Copy is a method that returns a deep copy of the stack.
//
// Returns:
//   - *Complex128Stack: A pointer to the newly created stack. Never returns nil.
func (s *Complex128Stack) Copy() *Complex128Stack {
	if s.front == nil {
		return &Complex128Stack{}
	}

	s_copy := &Complex128Stack{
		size: s.size,
	}

	node_copy := &stack_node_complex128{
		value: s.front.value,
	}

	s_copy.front = node_copy

	prev := node_copy

	for node := s.front.next; node != nil; node = node.next {
		node_copy := &stack_node_complex128{
			value: node.value,
		}

		prev.next = node_copy

		prev = node_copy
	}

	return s_copy
}