// Code generated with go generate. DO NOT EDIT.
package stack

import (
	"strconv"
	"strings"

	"github.com/PlayerR9/lib_units/common"
)

// stack_node_error is a node in the linked stack.
type stack_node_error struct {
	value error
	next *stack_node_error
}

// ErrorStack is a stack of error values implemented without a maximum capacity
// and using a linked list.
type ErrorStack struct {
	front *stack_node_error
	size int
}

// NewErrorStack creates a new linked stack.
//
// Returns:
//   - *ErrorStack: A pointer to the newly created stack. Never returns nil.
func NewErrorStack() *ErrorStack {
	return &ErrorStack{
		size: 0,
	}
}

// Push implements the stack.Stacker interface.
//
// Always returns true.
func (s *ErrorStack) Push(value error) bool {
	node := &stack_node_error{
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
func (s *ErrorStack) PushMany(values []error) int {
	if len(values) == 0 {
		return 0
	}

	node := &stack_node_error{
		value: values[0],
	}

	if s.front != nil {
		node.next = s.front
	}

	s.front = node

	for i := 1; i < len(values); i++ {
		node := &stack_node_error{
			value: values[i],
			next:  s.front,
		}

		s.front = node
	}

	s.size += len(values)
	
	return len(values)
}

// Pop implements the stack.Stacker interface.
func (s *ErrorStack) Pop() (error, bool) {
	if s.front == nil {
		return nil, false
	}

	to_remove := s.front
	s.front = s.front.next

	s.size--
	to_remove.next = nil

	return to_remove.value, true
}

// Peek implements the stack.Stacker interface.
func (s *ErrorStack) Peek() (error, bool) {
	if s.front == nil {
		return nil, false
	}

	return s.front.value, true
}

// IsEmpty implements the stack.Stacker interface.
func (s *ErrorStack) IsEmpty() bool {
	return s.front == nil
}

// Size implements the stack.Stacker interface.
func (s *ErrorStack) Size() int {
	return s.size
}

// Iterator implements the stack.Stacker interface.
func (s *ErrorStack) Iterator() common.Iterater[error] {
	var builder common.Builder[error]

	for node := s.front; node != nil; node = node.next {
		builder.Add(node.value)
	}

	return builder.Build()
}

// Clear implements the stack.Stacker interface.
func (s *ErrorStack) Clear() {
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
func (s *ErrorStack) GoString() string {
	values := make([]string, 0, s.size)
	for node := s.front; node != nil; node = node.next {
		values = append(values, common.StringOf(node.value))
	}

	var builder strings.Builder

	builder.WriteString("ErrorStack[size=")
	builder.WriteString(strconv.Itoa(s.size))
	builder.WriteString(", values=[")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString(" →]]")

	return builder.String()
}

// Slice implements the stack.Stacker interface.
//
// The 0th element is the top of the stack.
func (s *ErrorStack) Slice() []error {
	slice := make([]error, 0, s.size)

	for node := s.front; node != nil; node = node.next {
		slice = append(slice, node.value)
	}

	return slice
}

// Copy implements the stack.Stacker interface.
//
// The copy is a shallow copy.
func (s *ErrorStack) Copy() common.Copier {
	if s.front == nil {
		return &ErrorStack{}
	}

	s_copy := &ErrorStack{
		size: s.size,
	}

	node_copy := &stack_node_error{
		value: s.front.value,
	}

	s_copy.front = node_copy

	prev := node_copy

	for node := s.front.next; node != nil; node = node.next {
		node_copy := &stack_node_error{
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
func (s *ErrorStack) Capacity() int {
	return -1
}

// IsFull implements the stack.Stacker interface.
//
// Always returns false.
func (s *ErrorStack) IsFull() bool {
	return false
}
