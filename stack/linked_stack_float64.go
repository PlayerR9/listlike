// Code generated with go generate. DO NOT EDIT.
package stack

import (
	"github.com/PlayerR9/lib_units/common"
	"strconv"
	"strings"
)

// stack_node_float64 is a node in the linked stack.
type stack_node_float64 struct {
	value float64
	next *stack_node_float64
}

// Float64Stack is a stack of float64 values implemented without a maximum capacity
// and using a linked list.
type Float64Stack struct {
	front *stack_node_float64
	size int
}

// NewFloat64Stack creates a new linked stack.
//
// Returns:
//   - *Float64Stack: A pointer to the newly created stack. Never returns nil.
func NewFloat64Stack() *Float64Stack {
	return &Float64Stack{
		size: 0,
	}
}

// Push implements the stack.Stacker interface.
//
// Always returns true.
func (s *Float64Stack) Push(value float64) bool {
	node := &stack_node_float64{
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
func (s *Float64Stack) PushMany(values []float64) int {
	if len(values) == 0 {
		return 0
	}

	node := &stack_node_float64{
		value: values[0],
	}

	if s.front != nil {
		node.next = s.front
	}

	s.front = node

	for i := 1; i < len(values); i++ {
		node := &stack_node_float64{
			value: values[i],
			next:  s.front,
		}

		s.front = node
	}

	s.size += len(values)
	
	return len(values)
}

// Pop implements the stack.Stacker interface.
func (s *Float64Stack) Pop() (float64, bool) {
	if s.front == nil {
		return 0.0, false
	}

	to_remove := s.front
	s.front = s.front.next

	s.size--
	to_remove.next = nil

	return to_remove.value, true
}

// Peek implements the stack.Stacker interface.
func (s *Float64Stack) Peek() (float64, bool) {
	if s.front == nil {
		return 0.0, false
	}

	return s.front.value, true
}

// IsEmpty implements the stack.Stacker interface.
func (s *Float64Stack) IsEmpty() bool {
	return s.front == nil
}

// Size implements the stack.Stacker interface.
func (s *Float64Stack) Size() int {
	return s.size
}

// Iterator implements the stack.Stacker interface.
func (s *Float64Stack) Iterator() common.Iterater[float64] {
	var builder common.Builder[float64]

	for node := s.front; node != nil; node = node.next {
		builder.Add(node.value)
	}

	return builder.Build()
}

// Clear implements the stack.Stacker interface.
func (s *Float64Stack) Clear() {
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
func (s *Float64Stack) GoString() string {
	values := make([]string, 0, s.size)
	for node := s.front; node != nil; node = node.next {
		values = append(values, strconv.FormatFloat(node.value, 'f', -1, 64))
	}

	var builder strings.Builder

	builder.WriteString("Float64Stack[size=")
	builder.WriteString(strconv.Itoa(s.size))
	builder.WriteString(", values=[")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString(" →]]")

	return builder.String()
}

// Slice implements the stack.Stacker interface.
//
// The 0th element is the top of the stack.
func (s *Float64Stack) Slice() []float64 {
	slice := make([]float64, 0, s.size)

	for node := s.front; node != nil; node = node.next {
		slice = append(slice, node.value)
	}

	return slice
}

// Capacity implements the stack.Stacker interface.
//
// Always returns -1.
func (s *Float64Stack) Capacity() int {
	return -1
}

// IsFull implements the stack.Stacker interface.
//
// Always returns false.
func (s *Float64Stack) IsFull() bool {
	return false
}

// Copy is a method that returns a deep copy of the stack.
//
// Returns:
//   - *Float64Stack: A pointer to the newly created stack. Never returns nil.
func (s *Float64Stack) Copy() *Float64Stack {
	if s.front == nil {
		return &Float64Stack{}
	}

	s_copy := &Float64Stack{
		size: s.size,
	}

	node_copy := &stack_node_float64{
		value: s.front.value,
	}

	s_copy.front = node_copy

	prev := node_copy

	for node := s.front.next; node != nil; node = node.next {
		node_copy := &stack_node_float64{
			value: node.value,
		}

		prev.next = node_copy

		prev = node_copy
	}

	return s_copy
}