// Code generated with go generate. DO NOT EDIT.
package stack

import (
	"strconv"
	"strings"

	"github.com/PlayerR9/lib_units/common"
)

// stack_node_int16 is a node in the linked stack.
type stack_node_int16 struct {
	value int16
	next *stack_node_int16
}

// Int16Stack is a stack of int16 values implemented without a maximum capacity
// and using a linked list.
type Int16Stack struct {
	front *stack_node_int16
	size int
}

// NewInt16Stack creates a new linked stack.
//
// Returns:
//   - *Int16Stack: A pointer to the newly created stack. Never returns nil.
func NewInt16Stack() *Int16Stack {
	return &Int16Stack{
		size: 0,
	}
}

// Push implements the stack.Stacker interface.
//
// Always returns true.
func (s *Int16Stack) Push(value int16) bool {
	node := &stack_node_int16{
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
func (s *Int16Stack) PushMany(values []int16) int {
	if len(values) == 0 {
		return 0
	}

	node := &stack_node_int16{
		value: values[0],
	}

	if s.front != nil {
		node.next = s.front
	}

	s.front = node

	for i := 1; i < len(values); i++ {
		node := &stack_node_int16{
			value: values[i],
			next:  s.front,
		}

		s.front = node
	}

	s.size += len(values)
	
	return len(values)
}

// Pop implements the stack.Stacker interface.
func (s *Int16Stack) Pop() (int16, bool) {
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
func (s *Int16Stack) Peek() (int16, bool) {
	if s.front == nil {
		return 0, false
	}

	return s.front.value, true
}

// IsEmpty implements the stack.Stacker interface.
func (s *Int16Stack) IsEmpty() bool {
	return s.front == nil
}

// Size implements the stack.Stacker interface.
func (s *Int16Stack) Size() int {
	return s.size
}

// Iterator implements the stack.Stacker interface.
func (s *Int16Stack) Iterator() common.Iterater[int16] {
	var builder common.Builder[int16]

	for node := s.front; node != nil; node = node.next {
		builder.Add(node.value)
	}

	return builder.Build()
}

// Clear implements the stack.Stacker interface.
func (s *Int16Stack) Clear() {
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
func (s *Int16Stack) GoString() string {
	values := make([]string, 0, s.size)
	for node := s.front; node != nil; node = node.next {
		values = append(values, common.StringOf(node.value))
	}

	var builder strings.Builder

	builder.WriteString("Int16Stack[size=")
	builder.WriteString(strconv.Itoa(s.size))
	builder.WriteString(", values=[")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString(" →]]")

	return builder.String()
}

// Slice implements the stack.Stacker interface.
//
// The 0th element is the top of the stack.
func (s *Int16Stack) Slice() []int16 {
	slice := make([]int16, 0, s.size)

	for node := s.front; node != nil; node = node.next {
		slice = append(slice, node.value)
	}

	return slice
}

// Copy implements the stack.Stacker interface.
//
// The copy is a shallow copy.
func (s *Int16Stack) Copy() common.Copier {
	if s.front == nil {
		return &Int16Stack{}
	}

	s_copy := &Int16Stack{
		size: s.size,
	}

	node_copy := &stack_node_int16{
		value: s.front.value,
	}

	s_copy.front = node_copy

	prev := node_copy

	for node := s.front.next; node != nil; node = node.next {
		node_copy := &stack_node_int16{
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
func (s *Int16Stack) Capacity() int {
	return -1
}

// IsFull implements the stack.Stacker interface.
//
// Always returns false.
func (s *Int16Stack) IsFull() bool {
	return false
}
