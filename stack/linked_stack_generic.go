// Code generated with go generate. DO NOT EDIT.
package stack

import (
	"fmt"
	"github.com/PlayerR9/iterators/simple"
	"strconv"
	"strings"
)

// stack_node_T is a node in the linked stack.
type stack_node_T[T any] struct {
	value T
	next *stack_node_T[T]
}

// LinkedStack is a stack of T values implemented without a maximum capacity
// and using a linked list.
type LinkedStack[T any] struct {
	front *stack_node_T[T]
	size int
}

// NewLinkedStack creates a new linked stack.
//
// Returns:
//   - *LinkedStack[T]: A pointer to the newly created stack. Never returns nil.
func NewLinkedStack[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{
		size: 0,
	}
}

// Push implements the stack.Stacker interface.
//
// Always returns true.
func (s *LinkedStack[T]) Push(value T) bool {
	node := &stack_node_T[T]{
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
func (s *LinkedStack[T]) PushMany(values []T) int {
	if len(values) == 0 {
		return 0
	}

	node := &stack_node_T[T]{
		value: values[0],
	}

	if s.front != nil {
		node.next = s.front
	}

	s.front = node

	for i := 1; i < len(values); i++ {
		node := &stack_node_T[T]{
			value: values[i],
			next:  s.front,
		}

		s.front = node
	}

	s.size += len(values)
	
	return len(values)
}

// Pop implements the stack.Stacker interface.
func (s *LinkedStack[T]) Pop() (T, bool) {
	if s.front == nil {
		return *new(T), false
	}

	to_remove := s.front
	s.front = s.front.next

	s.size--
	to_remove.next = nil

	return to_remove.value, true
}

// Peek implements the stack.Stacker interface.
func (s *LinkedStack[T]) Peek() (T, bool) {
	if s.front == nil {
		return *new(T), false
	}

	return s.front.value, true
}

// IsEmpty implements the stack.Stacker interface.
func (s *LinkedStack[T]) IsEmpty() bool {
	return s.front == nil
}

// Size implements the stack.Stacker interface.
func (s *LinkedStack[T]) Size() int {
	return s.size
}

// Iterator implements the stack.Stacker interface.
func (s *LinkedStack[T]) Iterator() simple.Iterater[T] {
	var builder simple.Builder[T]

	for node := s.front; node != nil; node = node.next {
		builder.Add(node.value)
	}

	return builder.Build()
}

// Clear implements the stack.Stacker interface.
func (s *LinkedStack[T]) Clear() {
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
func (s *LinkedStack[T]) GoString() string {
	values := make([]string, 0, s.size)
	for node := s.front; node != nil; node = node.next {
		values = append(values, fmt.Sprintf("%v", node.value))
	}

	var builder strings.Builder

	builder.WriteString("LinkedStack[T][size=")
	builder.WriteString(strconv.Itoa(s.size))
	builder.WriteString(", values=[")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString(" →]]")

	return builder.String()
}

// Slice implements the stack.Stacker interface.
//
// The 0th element is the top of the stack.
func (s *LinkedStack[T]) Slice() []T {
	slice := make([]T, 0, s.size)

	for node := s.front; node != nil; node = node.next {
		slice = append(slice, node.value)
	}

	return slice
}

// Capacity implements the stack.Stacker interface.
//
// Always returns -1.
func (s *LinkedStack[T]) Capacity() int {
	return -1
}

// IsFull implements the stack.Stacker interface.
//
// Always returns false.
func (s *LinkedStack[T]) IsFull() bool {
	return false
}

// Copy is a method that returns a deep copy of the stack.
//
// Returns:
//   - *LinkedStack[T]: A pointer to the newly created stack. Never returns nil.
func (s *LinkedStack[T]) Copy() *LinkedStack[T] {
	if s.front == nil {
		return &LinkedStack[T]{}
	}

	s_copy := &LinkedStack[T]{
		size: s.size,
	}

	node_copy := &stack_node_T[T]{
		value: s.front.value,
	}

	s_copy.front = node_copy

	prev := node_copy

	for node := s.front.next; node != nil; node = node.next {
		node_copy := &stack_node_T[T]{
			value: node.value,
		}

		prev.next = node_copy

		prev = node_copy
	}

	return s_copy
}