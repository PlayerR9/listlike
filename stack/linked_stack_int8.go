// Code generated with go generate. DO NOT EDIT.
package stack

import (
	"github.com/PlayerR9/iterators/simple"
	"strconv"
	"strings"
)

// stack_node_int8 is a node in the linked stack.
type stack_node_int8 struct {
	value int8
	next *stack_node_int8
}

// Int8Stack is a stack of int8 values implemented without a maximum capacity
// and using a linked list.
type Int8Stack struct {
	front *stack_node_int8
	size int
}

// NewInt8Stack creates a new linked stack.
//
// Returns:
//   - *Int8Stack: A pointer to the newly created stack. Never returns nil.
func NewInt8Stack() *Int8Stack {
	return &Int8Stack{
		size: 0,
	}
}

// Push implements the stack.Stacker interface.
//
// Always returns true.
func (s *Int8Stack) Push(value int8) bool {
	node := &stack_node_int8{
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
func (s *Int8Stack) PushMany(values []int8) int {
	if len(values) == 0 {
		return 0
	}

	node := &stack_node_int8{
		value: values[0],
	}

	if s.front != nil {
		node.next = s.front
	}

	s.front = node

	for i := 1; i < len(values); i++ {
		node := &stack_node_int8{
			value: values[i],
			next:  s.front,
		}

		s.front = node
	}

	s.size += len(values)
	
	return len(values)
}

// Pop implements the stack.Stacker interface.
func (s *Int8Stack) Pop() (int8, bool) {
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
func (s *Int8Stack) Peek() (int8, bool) {
	if s.front == nil {
		return 0, false
	}

	return s.front.value, true
}

// IsEmpty implements the stack.Stacker interface.
func (s *Int8Stack) IsEmpty() bool {
	return s.front == nil
}

// Size implements the stack.Stacker interface.
func (s *Int8Stack) Size() int {
	return s.size
}

// Iterator implements the stack.Stacker interface.
func (s *Int8Stack) Iterator() simple.Iterater[int8] {
	var builder simple.Builder[int8]

	for node := s.front; node != nil; node = node.next {
		builder.Add(node.value)
	}

	return builder.Build()
}

// Clear implements the stack.Stacker interface.
func (s *Int8Stack) Clear() {
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
func (s *Int8Stack) GoString() string {
	values := make([]string, 0, s.size)
	for node := s.front; node != nil; node = node.next {
		values = append(values, strconv.FormatInt(int64(node.value), 10))
	}

	var builder strings.Builder

	builder.WriteString("Int8Stack[size=")
	builder.WriteString(strconv.Itoa(s.size))
	builder.WriteString(", values=[")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString(" →]]")

	return builder.String()
}

// Slice implements the stack.Stacker interface.
//
// The 0th element is the top of the stack.
func (s *Int8Stack) Slice() []int8 {
	slice := make([]int8, 0, s.size)

	for node := s.front; node != nil; node = node.next {
		slice = append(slice, node.value)
	}

	return slice
}

// Capacity implements the stack.Stacker interface.
//
// Always returns -1.
func (s *Int8Stack) Capacity() int {
	return -1
}

// IsFull implements the stack.Stacker interface.
//
// Always returns false.
func (s *Int8Stack) IsFull() bool {
	return false
}

// Copy is a method that returns a deep copy of the stack.
//
// Returns:
//   - *Int8Stack: A pointer to the newly created stack. Never returns nil.
func (s *Int8Stack) Copy() *Int8Stack {
	if s.front == nil {
		return &Int8Stack{}
	}

	s_copy := &Int8Stack{
		size: s.size,
	}

	node_copy := &stack_node_int8{
		value: s.front.value,
	}

	s_copy.front = node_copy

	prev := node_copy

	for node := s.front.next; node != nil; node = node.next {
		node_copy := &stack_node_int8{
			value: node.value,
		}

		prev.next = node_copy

		prev = node_copy
	}

	return s_copy
}