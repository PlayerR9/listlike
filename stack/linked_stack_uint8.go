// Code generated with go generate. DO NOT EDIT.
package stack

import (
	"github.com/PlayerR9/iterators/simple"
	"strconv"
	"strings"
)

// stack_node_uint8 is a node in the linked stack.
type stack_node_uint8 struct {
	value uint8
	next *stack_node_uint8
}

// Uint8Stack is a stack of uint8 values implemented without a maximum capacity
// and using a linked list.
type Uint8Stack struct {
	front *stack_node_uint8
	size int
}

// NewUint8Stack creates a new linked stack.
//
// Returns:
//   - *Uint8Stack: A pointer to the newly created stack. Never returns nil.
func NewUint8Stack() *Uint8Stack {
	return &Uint8Stack{
		size: 0,
	}
}

// Push implements the stack.Stacker interface.
//
// Always returns true.
func (s *Uint8Stack) Push(value uint8) bool {
	node := &stack_node_uint8{
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
func (s *Uint8Stack) PushMany(values []uint8) int {
	if len(values) == 0 {
		return 0
	}

	node := &stack_node_uint8{
		value: values[0],
	}

	if s.front != nil {
		node.next = s.front
	}

	s.front = node

	for i := 1; i < len(values); i++ {
		node := &stack_node_uint8{
			value: values[i],
			next:  s.front,
		}

		s.front = node
	}

	s.size += len(values)
	
	return len(values)
}

// Pop implements the stack.Stacker interface.
func (s *Uint8Stack) Pop() (uint8, bool) {
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
func (s *Uint8Stack) Peek() (uint8, bool) {
	if s.front == nil {
		return 0, false
	}

	return s.front.value, true
}

// IsEmpty implements the stack.Stacker interface.
func (s *Uint8Stack) IsEmpty() bool {
	return s.front == nil
}

// Size implements the stack.Stacker interface.
func (s *Uint8Stack) Size() int {
	return s.size
}

// Iterator implements the stack.Stacker interface.
func (s *Uint8Stack) Iterator() simple.Iterater[uint8] {
	var builder simple.Builder[uint8]

	for node := s.front; node != nil; node = node.next {
		builder.Add(node.value)
	}

	return builder.Build()
}

// Clear implements the stack.Stacker interface.
func (s *Uint8Stack) Clear() {
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
func (s *Uint8Stack) GoString() string {
	values := make([]string, 0, s.size)
	for node := s.front; node != nil; node = node.next {
		values = append(values, strconv.FormatUint(uint64(node.value), 10))
	}

	var builder strings.Builder

	builder.WriteString("Uint8Stack[size=")
	builder.WriteString(strconv.Itoa(s.size))
	builder.WriteString(", values=[")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString(" →]]")

	return builder.String()
}

// Slice implements the stack.Stacker interface.
//
// The 0th element is the top of the stack.
func (s *Uint8Stack) Slice() []uint8 {
	slice := make([]uint8, 0, s.size)

	for node := s.front; node != nil; node = node.next {
		slice = append(slice, node.value)
	}

	return slice
}

// Capacity implements the stack.Stacker interface.
//
// Always returns -1.
func (s *Uint8Stack) Capacity() int {
	return -1
}

// IsFull implements the stack.Stacker interface.
//
// Always returns false.
func (s *Uint8Stack) IsFull() bool {
	return false
}

// Copy is a method that returns a deep copy of the stack.
//
// Returns:
//   - *Uint8Stack: A pointer to the newly created stack. Never returns nil.
func (s *Uint8Stack) Copy() *Uint8Stack {
	if s.front == nil {
		return &Uint8Stack{}
	}

	s_copy := &Uint8Stack{
		size: s.size,
	}

	node_copy := &stack_node_uint8{
		value: s.front.value,
	}

	s_copy.front = node_copy

	prev := node_copy

	for node := s.front.next; node != nil; node = node.next {
		node_copy := &stack_node_uint8{
			value: node.value,
		}

		prev.next = node_copy

		prev = node_copy
	}

	return s_copy
}