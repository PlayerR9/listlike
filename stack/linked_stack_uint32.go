// Code generated with go generate. DO NOT EDIT.
package stack

import (
	
	"github.com/PlayerR9/lib_units/common"
	"strconv"
	"strings"
)

// stack_node_uint32 is a node in the linked stack.
type stack_node_uint32 struct {
	value uint32
	next *stack_node_uint32
}

// Uint32Stack is a stack of uint32 values implemented without a maximum capacity
// and using a linked list.
type Uint32Stack struct {
	front *stack_node_uint32
	size int
}

// NewUint32Stack creates a new linked stack.
//
// Returns:
//   - *Uint32Stack: A pointer to the newly created stack. Never returns nil.
func NewUint32Stack() *Uint32Stack {
	return &Uint32Stack{
		size: 0,
	}
}

// Push implements the stack.Stacker interface.
//
// Always returns true.
func (s *Uint32Stack) Push(value uint32) bool {
	node := &stack_node_uint32{
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
func (s *Uint32Stack) PushMany(values []uint32) int {
	if len(values) == 0 {
		return 0
	}

	node := &stack_node_uint32{
		value: values[0],
	}

	if s.front != nil {
		node.next = s.front
	}

	s.front = node

	for i := 1; i < len(values); i++ {
		node := &stack_node_uint32{
			value: values[i],
			next:  s.front,
		}

		s.front = node
	}

	s.size += len(values)
	
	return len(values)
}

// Pop implements the stack.Stacker interface.
func (s *Uint32Stack) Pop() (uint32, bool) {
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
func (s *Uint32Stack) Peek() (uint32, bool) {
	if s.front == nil {
		return 0, false
	}

	return s.front.value, true
}

// IsEmpty implements the stack.Stacker interface.
func (s *Uint32Stack) IsEmpty() bool {
	return s.front == nil
}

// Size implements the stack.Stacker interface.
func (s *Uint32Stack) Size() int {
	return s.size
}

// Iterator implements the stack.Stacker interface.
func (s *Uint32Stack) Iterator() common.Iterater[uint32] {
	var builder common.Builder[uint32]

	for node := s.front; node != nil; node = node.next {
		builder.Add(node.value)
	}

	return builder.Build()
}

// Clear implements the stack.Stacker interface.
func (s *Uint32Stack) Clear() {
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
func (s *Uint32Stack) GoString() string {
	values := make([]string, 0, s.size)
	for node := s.front; node != nil; node = node.next {
		values = append(values, strconv.FormatUint(uint64(node.value), 10))
	}

	var builder strings.Builder

	builder.WriteString("Uint32Stack[size=")
	builder.WriteString(strconv.Itoa(s.size))
	builder.WriteString(", values=[")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString(" →]]")

	return builder.String()
}

// Slice implements the stack.Stacker interface.
//
// The 0th element is the top of the stack.
func (s *Uint32Stack) Slice() []uint32 {
	slice := make([]uint32, 0, s.size)

	for node := s.front; node != nil; node = node.next {
		slice = append(slice, node.value)
	}

	return slice
}

// Copy implements the stack.Stacker interface.
//
// The copy is a shallow copy.
func (s *Uint32Stack) Copy() common.Copier {
	if s.front == nil {
		return &Uint32Stack{}
	}

	s_copy := &Uint32Stack{
		size: s.size,
	}

	node_copy := &stack_node_uint32{
		value: s.front.value,
	}

	s_copy.front = node_copy

	prev := node_copy

	for node := s.front.next; node != nil; node = node.next {
		node_copy := &stack_node_uint32{
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
func (s *Uint32Stack) Capacity() int {
	return -1
}

// IsFull implements the stack.Stacker interface.
//
// Always returns false.
func (s *Uint32Stack) IsFull() bool {
	return false
}
