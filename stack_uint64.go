// Code generated with go generate. DO NOT EDIT.
package stack

import (
	"strconv"
	"strings"

	"github.com/PlayerR9/lib_units/common"
)

// stack_node_uint64 is a node in the linked stack.
type stack_node_uint64 struct {
	value uint64
	next *stack_node_uint64
}

// Uint64Stack is a stack of uint64 values implemented without a maximum capacity
// and using a linked list.
type Uint64Stack struct {
	front *stack_node_uint64
	size int
}

// NewUint64Stack creates a new linked stack.
//
// Returns:
//   - *Uint64Stack: A pointer to the newly created stack. Never returns nil.
func NewUint64Stack() *Uint64Stack {
	return &Uint64Stack{
		size: 0,
	}
}

// Push implements the stack.Stacker interface.
//
// Always returns true.
func (s *Uint64Stack) Push(value uint64) bool {
	node := &stack_node_uint64{
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
func (s *Uint64Stack) PushMany(values []uint64) int {
	if len(values) == 0 {
		return 0
	}

	node := &stack_node_uint64{
		value: values[0],
	}

	if s.front != nil {
		node.next = s.front
	}

	s.front = node

	for i := 1; i < len(values); i++ {
		node := &stack_node_uint64{
			value: values[i],
			next:  s.front,
		}

		s.front = node
	}

	s.size += len(values)
	
	return len(values)
}

// Pop implements the stack.Stacker interface.
func (s *Uint64Stack) Pop() (uint64, bool) {
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
func (s *Uint64Stack) Peek() (uint64, bool) {
	if s.front == nil {
		return 0, false
	}

	return s.front.value, true
}

// IsEmpty implements the stack.Stacker interface.
func (s *Uint64Stack) IsEmpty() bool {
	return s.front == nil
}

// Size implements the stack.Stacker interface.
func (s *Uint64Stack) Size() int {
	return s.size
}

// Iterator implements the stack.Stacker interface.
func (s *Uint64Stack) Iterator() common.Iterater[uint64] {
	var builder common.Builder[uint64]

	for node := s.front; node != nil; node = node.next {
		builder.Add(node.value)
	}

	return builder.Build()
}

// Clear implements the stack.Stacker interface.
func (s *Uint64Stack) Clear() {
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
func (s *Uint64Stack) GoString() string {
	values := make([]string, 0, s.size)
	for node := s.front; node != nil; node = node.next {
		values = append(values, common.StringOf(node.value))
	}

	var builder strings.Builder

	builder.WriteString("Uint64Stack[size=")
	builder.WriteString(strconv.Itoa(s.size))
	builder.WriteString(", values=[")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString(" →]]")

	return builder.String()
}

// Slice implements the stack.Stacker interface.
//
// The 0th element is the top of the stack.
func (s *Uint64Stack) Slice() []uint64 {
	slice := make([]uint64, 0, s.size)

	for node := s.front; node != nil; node = node.next {
		slice = append(slice, node.value)
	}

	return slice
}

// Copy implements the stack.Stacker interface.
//
// The copy is a shallow copy.
func (s *Uint64Stack) Copy() common.Copier {
	if s.front == nil {
		return &Uint64Stack{}
	}

	s_copy := &Uint64Stack{
		size: s.size,
	}

	node_copy := &stack_node_uint64{
		value: s.front.value,
	}

	s_copy.front = node_copy

	prev := node_copy

	for node := s.front.next; node != nil; node = node.next {
		node_copy := &stack_node_uint64{
			value: node.value,
		}

		prev.next = node_copy

		prev = node_copy
	}

	return s_copy
}

// Capacity implements the stack.Stacker interface.
//
// Always returns -1.
func (s *Uint64Stack) Capacity() int {
	return -1
}

// IsFull implements the stack.Stacker interface.
//
// Always returns false.
func (s *Uint64Stack) IsFull() bool {
	return false
}
