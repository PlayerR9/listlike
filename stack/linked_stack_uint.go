// Code generated with go generate. DO NOT EDIT.
package stack

import (
	"github.com/PlayerR9/lib_units/common"
	"strconv"
	"strings"
)

// stack_node_uint is a node in the linked stack.
type stack_node_uint struct {
	value uint
	next *stack_node_uint
}

// UintStack is a stack of uint values implemented without a maximum capacity
// and using a linked list.
type UintStack struct {
	front *stack_node_uint
	size int
}

// NewUintStack creates a new linked stack.
//
// Returns:
//   - *UintStack: A pointer to the newly created stack. Never returns nil.
func NewUintStack() *UintStack {
	return &UintStack{
		size: 0,
	}
}

// Push implements the stack.Stacker interface.
//
// Always returns true.
func (s *UintStack) Push(value uint) bool {
	node := &stack_node_uint{
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
func (s *UintStack) PushMany(values []uint) int {
	if len(values) == 0 {
		return 0
	}

	node := &stack_node_uint{
		value: values[0],
	}

	if s.front != nil {
		node.next = s.front
	}

	s.front = node

	for i := 1; i < len(values); i++ {
		node := &stack_node_uint{
			value: values[i],
			next:  s.front,
		}

		s.front = node
	}

	s.size += len(values)
	
	return len(values)
}

// Pop implements the stack.Stacker interface.
func (s *UintStack) Pop() (uint, bool) {
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
func (s *UintStack) Peek() (uint, bool) {
	if s.front == nil {
		return 0, false
	}

	return s.front.value, true
}

// IsEmpty implements the stack.Stacker interface.
func (s *UintStack) IsEmpty() bool {
	return s.front == nil
}

// Size implements the stack.Stacker interface.
func (s *UintStack) Size() int {
	return s.size
}

// Iterator implements the stack.Stacker interface.
func (s *UintStack) Iterator() common.Iterater[uint] {
	var builder common.Builder[uint]

	for node := s.front; node != nil; node = node.next {
		builder.Add(node.value)
	}

	return builder.Build()
}

// Clear implements the stack.Stacker interface.
func (s *UintStack) Clear() {
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
func (s *UintStack) GoString() string {
	values := make([]string, 0, s.size)
	for node := s.front; node != nil; node = node.next {
		values = append(values, strconv.FormatUint(uint64(node.value), 10))
	}

	var builder strings.Builder

	builder.WriteString("UintStack[size=")
	builder.WriteString(strconv.Itoa(s.size))
	builder.WriteString(", values=[")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString(" →]]")

	return builder.String()
}

// Slice implements the stack.Stacker interface.
//
// The 0th element is the top of the stack.
func (s *UintStack) Slice() []uint {
	slice := make([]uint, 0, s.size)

	for node := s.front; node != nil; node = node.next {
		slice = append(slice, node.value)
	}

	return slice
}

// Capacity implements the stack.Stacker interface.
//
// Always returns -1.
func (s *UintStack) Capacity() int {
	return -1
}

// IsFull implements the stack.Stacker interface.
//
// Always returns false.
func (s *UintStack) IsFull() bool {
	return false
}

// Copy is a method that returns a deep copy of the stack.
//
// Returns:
//   - *UintStack: A pointer to the newly created stack. Never returns nil.
func (s *UintStack) Copy() *UintStack {
	if s.front == nil {
		return &UintStack{}
	}

	s_copy := &UintStack{
		size: s.size,
	}

	node_copy := &stack_node_uint{
		value: s.front.value,
	}

	s_copy.front = node_copy

	prev := node_copy

	for node := s.front.next; node != nil; node = node.next {
		node_copy := &stack_node_uint{
			value: node.value,
		}

		prev.next = node_copy

		prev = node_copy
	}

	return s_copy
}