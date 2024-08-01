// Code generated with go generate. DO NOT EDIT.
package stack

import (
	"github.com/PlayerR9/lib_units/common"
	"strconv"
	"strings"
)

// stack_node_uint16 is a node in the linked stack.
type stack_node_uint16 struct {
	value uint16
	next *stack_node_uint16
}

// Uint16Stack is a stack of uint16 values implemented without a maximum capacity
// and using a linked list.
type Uint16Stack struct {
	front *stack_node_uint16
	size int
}

// NewUint16Stack creates a new linked stack.
//
// Returns:
//   - *Uint16Stack: A pointer to the newly created stack. Never returns nil.
func NewUint16Stack() *Uint16Stack {
	return &Uint16Stack{
		size: 0,
	}
}

// Push implements the stack.Stacker interface.
//
// Always returns true.
func (s *Uint16Stack) Push(value uint16) bool {
	node := &stack_node_uint16{
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
func (s *Uint16Stack) PushMany(values []uint16) int {
	if len(values) == 0 {
		return 0
	}

	node := &stack_node_uint16{
		value: values[0],
	}

	if s.front != nil {
		node.next = s.front
	}

	s.front = node

	for i := 1; i < len(values); i++ {
		node := &stack_node_uint16{
			value: values[i],
			next:  s.front,
		}

		s.front = node
	}

	s.size += len(values)
	
	return len(values)
}

// Pop implements the stack.Stacker interface.
func (s *Uint16Stack) Pop() (uint16, bool) {
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
func (s *Uint16Stack) Peek() (uint16, bool) {
	if s.front == nil {
		return 0, false
	}

	return s.front.value, true
}

// IsEmpty implements the stack.Stacker interface.
func (s *Uint16Stack) IsEmpty() bool {
	return s.front == nil
}

// Size implements the stack.Stacker interface.
func (s *Uint16Stack) Size() int {
	return s.size
}

// Iterator implements the stack.Stacker interface.
func (s *Uint16Stack) Iterator() common.Iterater[uint16] {
	var builder common.Builder[uint16]

	for node := s.front; node != nil; node = node.next {
		builder.Add(node.value)
	}

	return builder.Build()
}

// Clear implements the stack.Stacker interface.
func (s *Uint16Stack) Clear() {
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
func (s *Uint16Stack) GoString() string {
	values := make([]string, 0, s.size)
	for node := s.front; node != nil; node = node.next {
		values = append(values, strconv.FormatUint(uint64(node.value), 10))
	}

	var builder strings.Builder

	builder.WriteString("Uint16Stack[size=")
	builder.WriteString(strconv.Itoa(s.size))
	builder.WriteString(", values=[")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString(" →]]")

	return builder.String()
}

// Slice implements the stack.Stacker interface.
//
// The 0th element is the top of the stack.
func (s *Uint16Stack) Slice() []uint16 {
	slice := make([]uint16, 0, s.size)

	for node := s.front; node != nil; node = node.next {
		slice = append(slice, node.value)
	}

	return slice
}

// Capacity implements the stack.Stacker interface.
//
// Always returns -1.
func (s *Uint16Stack) Capacity() int {
	return -1
}

// IsFull implements the stack.Stacker interface.
//
// Always returns false.
func (s *Uint16Stack) IsFull() bool {
	return false
}

// Copy is a method that returns a deep copy of the stack.
//
// Returns:
//   - *Uint16Stack: A pointer to the newly created stack. Never returns nil.
func (s *Uint16Stack) Copy() *Uint16Stack {
	if s.front == nil {
		return &Uint16Stack{}
	}

	s_copy := &Uint16Stack{
		size: s.size,
	}

	node_copy := &stack_node_uint16{
		value: s.front.value,
	}

	s_copy.front = node_copy

	prev := node_copy

	for node := s.front.next; node != nil; node = node.next {
		node_copy := &stack_node_uint16{
			value: node.value,
		}

		prev.next = node_copy

		prev = node_copy
	}

	return s_copy
}