// Code generated with go generate. DO NOT EDIT.
package stack

import (
	"github.com/PlayerR9/iterators/simple"
	"strconv"
	"strings"
)

// stack_node_string is a node in the linked stack.
type stack_node_string struct {
	value string
	next *stack_node_string
}

// StringStack is a stack of string values implemented without a maximum capacity
// and using a linked list.
type StringStack struct {
	front *stack_node_string
	size int
}

// NewStringStack creates a new linked stack.
//
// Returns:
//   - *StringStack: A pointer to the newly created stack. Never returns nil.
func NewStringStack() *StringStack {
	return &StringStack{
		size: 0,
	}
}

// Push implements the stack.Stacker interface.
//
// Always returns true.
func (s *StringStack) Push(value string) bool {
	node := &stack_node_string{
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
func (s *StringStack) PushMany(values []string) int {
	if len(values) == 0 {
		return 0
	}

	node := &stack_node_string{
		value: values[0],
	}

	if s.front != nil {
		node.next = s.front
	}

	s.front = node

	for i := 1; i < len(values); i++ {
		node := &stack_node_string{
			value: values[i],
			next:  s.front,
		}

		s.front = node
	}

	s.size += len(values)
	
	return len(values)
}

// Pop implements the stack.Stacker interface.
func (s *StringStack) Pop() (string, bool) {
	if s.front == nil {
		return "", false
	}

	to_remove := s.front
	s.front = s.front.next

	s.size--
	to_remove.next = nil

	return to_remove.value, true
}

// Peek implements the stack.Stacker interface.
func (s *StringStack) Peek() (string, bool) {
	if s.front == nil {
		return "", false
	}

	return s.front.value, true
}

// IsEmpty implements the stack.Stacker interface.
func (s *StringStack) IsEmpty() bool {
	return s.front == nil
}

// Size implements the stack.Stacker interface.
func (s *StringStack) Size() int {
	return s.size
}

// Iterator implements the stack.Stacker interface.
func (s *StringStack) Iterator() simple.Iterater[string] {
	var builder simple.Builder[string]

	for node := s.front; node != nil; node = node.next {
		builder.Add(node.value)
	}

	return builder.Build()
}

// Clear implements the stack.Stacker interface.
func (s *StringStack) Clear() {
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
func (s *StringStack) GoString() string {
	values := make([]string, 0, s.size)
	for node := s.front; node != nil; node = node.next {
		values = append(values, node.value)
	}

	var builder strings.Builder

	builder.WriteString("StringStack[size=")
	builder.WriteString(strconv.Itoa(s.size))
	builder.WriteString(", values=[")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString(" →]]")

	return builder.String()
}

// Slice implements the stack.Stacker interface.
//
// The 0th element is the top of the stack.
func (s *StringStack) Slice() []string {
	slice := make([]string, 0, s.size)

	for node := s.front; node != nil; node = node.next {
		slice = append(slice, node.value)
	}

	return slice
}

// Capacity implements the stack.Stacker interface.
//
// Always returns -1.
func (s *StringStack) Capacity() int {
	return -1
}

// IsFull implements the stack.Stacker interface.
//
// Always returns false.
func (s *StringStack) IsFull() bool {
	return false
}

// Copy is a method that returns a deep copy of the stack.
//
// Returns:
//   - *StringStack: A pointer to the newly created stack. Never returns nil.
func (s *StringStack) Copy() *StringStack {
	if s.front == nil {
		return &StringStack{}
	}

	s_copy := &StringStack{
		size: s.size,
	}

	node_copy := &stack_node_string{
		value: s.front.value,
	}

	s_copy.front = node_copy

	prev := node_copy

	for node := s.front.next; node != nil; node = node.next {
		node_copy := &stack_node_string{
			value: node.value,
		}

		prev.next = node_copy

		prev = node_copy
	}

	return s_copy
}