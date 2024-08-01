package list

import (
	"fmt"
	"strconv"
	"strings"

	uc "github.com/PlayerR9/lib_units/common"
)

// LimitedArrayList is a generic type that represents a list data structure with
// or without a limited capacity. It is implemented using an array.
type LimitedArrayList[T any] struct {
	// values is a slice of type T that stores the elements in the list.
	values []T

	// capacity is the maximum number of elements the list can hold.
	capacity int
}

// NewLimitedArrayList is a function that creates and returns a new instance of a
// LimitedArrayList.
//
// Parameters:
//
//   - capacity: An integer that represents the maximum number of elements the list.
//     can hold. If the capacity is negative, the value is converted to a positive
//     value.
//   - values: A variadic parameter of type T, which represents the initial values to
//     be stored in the list.
//
// Returns:
//
//   - *LimitedArrayList[T]: A pointer to the newly created LimitedArrayList.
func NewLimitedArrayList[T any](capacity int, values ...T) *LimitedArrayList[T] {
	if capacity < 0 {
		capacity *= -1
	}

	list := &LimitedArrayList[T]{
		values:   make([]T, len(values), capacity),
		capacity: capacity,
	}
	copy(list.values, values)

	return list
}

// Append implements the Lister interface.
func (list *LimitedArrayList[T]) Append(value T) bool {
	if len(list.values) >= list.capacity {
		return false
	}

	list.values = append(list.values, value)

	return true
}

// DeleteFirst implements the Lister interface.
func (list *LimitedArrayList[T]) DeleteFirst() (T, bool) {
	if len(list.values) <= 0 {
		return *new(T), false
	}

	toRemove := list.values[0]
	list.values = list.values[1:]
	return toRemove, true
}

// PeekFirst implements the Lister interface.
func (list *LimitedArrayList[T]) PeekFirst() (T, bool) {
	if len(list.values) == 0 {
		return *new(T), false
	}

	elem := list.values[0]

	return elem, true
}

// IsEmpty is a method of the LimitedArrayList type. It checks if the list is empty.
//
// Returns:
//
//   - bool: A boolean value that is true if the list is empty, and false otherwise.
func (list *LimitedArrayList[T]) IsEmpty() bool {
	return len(list.values) == 0
}

// Size is a method of the LimitedArrayList type. It returns the number of elements in
// the list.
//
// Returns:
//
//   - int: An integer that represents the number of elements in the list.
func (list *LimitedArrayList[T]) Size() int {
	return len(list.values)
}

// Capacity is a method of the LimitedArrayList type. It returns the maximum number of
// elements the list can hold.
//
// Returns:
//
//   - optional.Int: An optional integer that represents the maximum number of
//     elements the list can hold.
func (list *LimitedArrayList[T]) Capacity() int {
	return list.capacity
}

// Iterator is a method of the LimitedArrayList type. It returns an iterator for the list.
//
// Returns:
//
//   - uc.Iterater[T]: An iterator for the list.
func (list *LimitedArrayList[T]) Iterator() uc.Iterater[T] {
	return uc.NewSimpleIterator(list.values)
}

// Clear is a method of the LimitedArrayList type. It is used to remove all elements from
// the list.
func (list *LimitedArrayList[T]) Clear() {
	list.values = make([]T, 0, list.capacity)
}

// IsFull is a method of the LimitedArrayList type. It checks if the list is full.
//
// Returns:
//
//   - isFull: A boolean value that is true if the list is full, and false otherwise.
func (list *LimitedArrayList[T]) IsFull() bool {
	return list.capacity <= len(list.values)
}

// GoString implements the fmt.GoStringer interface.
func (list *LimitedArrayList[T]) GoString() string {
	values := make([]string, 0, len(list.values))
	for _, v := range list.values {
		values = append(values, fmt.Sprintf("%v", v))
	}

	var builder strings.Builder

	builder.WriteString("LimitedArrayList[capacity=")
	builder.WriteString(strconv.Itoa(list.capacity))
	builder.WriteString(", size=")
	builder.WriteString(strconv.Itoa(len(list.values)))
	builder.WriteString(", values=[")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString("]]")

	return builder.String()
}

// Prepend implements the Lister interface.
func (list *LimitedArrayList[T]) Prepend(value T) bool {
	if len(list.values) >= list.capacity {
		return false
	}

	list.values = append([]T{value}, list.values...)

	return true
}

// DeleteLast implements the Lister interface.
func (list *LimitedArrayList[T]) DeleteLast() (T, bool) {
	if len(list.values) == 0 {
		return *new(T), false
	}

	toRemove := list.values[len(list.values)-1]
	list.values = list.values[:len(list.values)-1]
	return toRemove, true
}

// PeekLast implements the Lister interface.
func (list *LimitedArrayList[T]) PeekLast() (T, bool) {
	if len(list.values) == 0 {
		return *new(T), false
	}

	return list.values[len(list.values)-1], true
}

// Slice is a method of the LimitedArrayList type that returns a slice of type T
// containing the elements of the list.
//
// Returns:
//
//   - []T: A slice of type T containing the elements of the list.
func (list *LimitedArrayList[T]) Slice() []T {
	slice := make([]T, len(list.values))
	copy(slice, list.values)

	return slice
}

// Copy is a method of the LimitedArrayList type. It is used to create a shallow copy
// of the list.
//
// Returns:
//   - *LimitedArrayList[T]: A copy of the list.
func (list *LimitedArrayList[T]) Copy() *LimitedArrayList[T] {
	list_copy := &LimitedArrayList[T]{
		values:   make([]T, len(list.values)),
		capacity: list.capacity,
	}
	copy(list_copy.values, list.values)

	return list_copy
}
