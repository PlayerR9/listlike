// Code generated with go generate. DO NOT EDIT.
package stack

import (
	
	"github.com/PlayerR9/lib_units/common"
	"strconv"
	"strings"
)

// stack_node_uintptr is a node in the linked stack.
type stack_node_uintptr struct {
	value uintptr
	next *stack_node_uintptr
}

// UintptrStack is a stack of uintptr values implemented without a maximum capacity
// and using a linked list.
type UintptrStack struct {
	front *stack_node_uintptr
	size int
}

// NewUintptrStack creates a new linked stack.
//
// Returns:
//   - *UintptrStack: A pointer to the newly created stack. Never returns nil.
func NewUintptrStack() *UintptrStack {
	return &UintptrStack{
		size: 0,
	}
}

// Push implements the stack.Stacker interface.
//
// Always returns true.
func (s *UintptrStack) Push(value uintptr) bool {
	node := &stack_node_uintptr{
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
func (s *UintptrStack) PushMany(values []uintptr) int {
	if len(values) == 0 {
		return 0
	}

	node := &stack_node_uintptr{
		value: values[0],
	}

	if s.front != nil {
		node.next = s.front
	}

	s.front = node

	for i := 1; i < len(values); i++ {
		node := &stack_node_uintptr{
			value: values[i],
			next:  s.front,
		}

		s.front = node
	}

	s.size += len(values)
	
	return len(values)
}

// Pop implements the stack.Stacker interface.
func (s *UintptrStack) Pop() (uintptr, bool) {
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
func (s *UintptrStack) Peek() (uintptr, bool) {
	if s.front == nil {
		return 0, false
	}

	return s.front.value, true
}

// IsEmpty implements the stack.Stacker interface.
func (s *UintptrStack) IsEmpty() bool {
	return s.front == nil
}

// Size implements the stack.Stacker interface.
func (s *UintptrStack) Size() int {
	return s.size
}

// Iterator implements the stack.Stacker interface.
func (s *UintptrStack) Iterator() common.Iterater[uintptr] {
	var builder common.Builder[uintptr]

	for node := s.front; node != nil; node = node.next {
		builder.Add(node.value)
	}

	return builder.Build()
}

// Clear implements the stack.Stacker interface.
func (s *UintptrStack) Clear() {
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
func (s *UintptrStack) GoString() string {
	values := make([]string, 0, s.size)
	for node := s.front; node != nil; node = node.next {
		values = append(values, strconv.FormatUint(uint64(node.value), 10))
	}

	var builder strings.Builder

	builder.WriteString("UintptrStack[size=")
	builder.WriteString(strconv.Itoa(s.size))
	builder.WriteString(", values=[")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString(" →]]")

	return builder.String()
}

// Slice implements the stack.Stacker interface.
//
// The 0th element is the top of the stack.
func (s *UintptrStack) Slice() []uintptr {
	slice := make([]uintptr, 0, s.size)

	for node := s.front; node != nil; node = node.next {
		slice = append(slice, node.value)
	}

	return slice
}

// Copy implements the stack.Stacker interface.
//
// The copy is a shallow copy.
func (s *UintptrStack) Copy() common.Copier {
	if s.front == nil {
		return &UintptrStack{}
	}

	s_copy := &UintptrStack{
		size: s.size,
	}

	node_copy := &stack_node_uintptr{
		value: s.front.value,
	}

	s_copy.front = node_copy

	prev := node_copy

	for node := s.front.next; node != nil; node = node.next {
		node_copy := &stack_node_uintptr{
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
func (s *UintptrStack) Capacity() int {
	return -1
}

// IsFull implements the stack.Stacker interface.
//
// Always returns false.
func (s *UintptrStack) IsFull() bool {
	return false
}
