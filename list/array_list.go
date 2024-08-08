package list

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

// ArrayIterator is the iterator for the Lister interface.
type ArrayIterator[T any] struct {
	// values is a slice of type T, which represents the values to be iterated over.
	values []T

	// pos is an integer that represents the current position in the list.
	pos int
}

// Consume implements the Iterater interface.
func (it *ArrayIterator[T]) Consume() (T, error) {
	if it.pos >= len(it.values) {
		return *new(T), io.EOF
	}

	val := it.values[it.pos]

	it.pos++

	return val, nil
}

// Restart implements the Iterater interface.
func (it *ArrayIterator[T]) Restart() {
	it.pos = 0
}

// ArrayList is a generic type that represents a list data structure with
// or without a limited capacity. It is implemented using an array.
type ArrayList[T any] struct {
	// values is a slice of type T that stores the elements in the list.
	values []T

	// capacity is the maximum number of elements the list can hold.
	capacity int
}

// NewArrayList is a function that creates and returns a new instance of a
// ArrayList.
//
// Parameters:
//
//   - values: A variadic parameter of type T, which represents the initial values to
//     be stored in the list.
//
// Returns:
//   - *ArrayList[T]: A pointer to the newly created ArrayList.
func NewArrayList[T any](values ...T) *ArrayList[T] {
	list := &ArrayList[T]{
		values:   make([]T, len(values)),
		capacity: -1,
	}
	copy(list.values, values)

	return list
}

// NewArrayList is a function that creates and returns a new instance of a
// ArrayList.
//
// Parameters:
//   - capacity: An integer that represents the maximum number of elements the list.
//     can hold. If the capacity is negative, the value is converted to a positive
//     value.
//   - values: A variadic parameter of type T, which represents the initial values to
//     be stored in the list.
//
// Returns:
//   - *ArrayList[T]: A pointer to the newly created ArrayList.
func NewLimitedArrayList[T any](capacity int, values ...T) *ArrayList[T] {
	if capacity < 0 {
		capacity *= -1
	}

	list := &ArrayList[T]{
		values:   make([]T, len(values), capacity),
		capacity: capacity,
	}
	copy(list.values, values)

	return list
}

// Append implements the Lister interface.
func (list *ArrayList[T]) Append(value T) bool {
	if list.capacity != -1 && len(list.values) >= list.capacity {
		return false
	}

	list.values = append(list.values, value)

	return true
}

// DeleteFirst implements the Lister interface.
func (list *ArrayList[T]) DeleteFirst() (T, bool) {
	if len(list.values) <= 0 {
		return *new(T), false
	}

	toRemove := list.values[0]
	list.values = list.values[1:]
	return toRemove, true
}

// PeekFirst implements the Lister interface.
func (list *ArrayList[T]) PeekFirst() (T, bool) {
	if len(list.values) == 0 {
		return *new(T), false
	}

	elem := list.values[0]

	return elem, true
}

// IsEmpty is a method of the ArrayList type. It checks if the list is empty.
//
// Returns:
//
//   - bool: A boolean value that is true if the list is empty, and false otherwise.
func (list *ArrayList[T]) IsEmpty() bool {
	return len(list.values) == 0
}

// Size is a method of the ArrayList type. It returns the number of elements in
// the list.
//
// Returns:
//
//   - int: An integer that represents the number of elements in the list.
func (list *ArrayList[T]) Size() int {
	return len(list.values)
}

// Capacity is a method of the ArrayList type. It returns the maximum number of
// elements the list can hold.
//
// Returns:
//
//   - optional.Int: An optional integer that represents the maximum number of
//     elements the list can hold.
func (list *ArrayList[T]) Capacity() int {
	return list.capacity
}

// Iterator is a method of the ArrayList type. It returns an iterator for the list.
//
// Returns:
//   - *ArrayIterator[T]: An iterator for the list.
func (list *ArrayList[T]) Iterator() *ArrayIterator[T] {
	return &ArrayIterator[T]{
		values: list.values,
		pos:    0,
	}
}

// Clear is a method of the ArrayList type. It is used to remove all elements from
// the list.
func (list *ArrayList[T]) Clear() {
	if list.capacity == -1 {
		list.values = make([]T, 0)
	} else {
		list.values = make([]T, 0, list.capacity)
	}
}

// IsFull is a method of the ArrayList type. It checks if the list is full.
//
// Returns:
//   - isFull: A boolean value that is true if the list is full, and false otherwise.
func (list *ArrayList[T]) IsFull() bool {
	return list.capacity != -1 && len(list.values) >= list.capacity
}

// GoString implements the fmt.GoStringer interface.
func (list *ArrayList[T]) GoString() string {
	values := make([]string, 0, len(list.values))
	for _, v := range list.values {
		values = append(values, fmt.Sprintf("%v", v))
	}

	var builder strings.Builder

	builder.WriteString("ArrayList[")

	if list.capacity != -1 {
		builder.WriteString("capacity=")
		builder.WriteString(strconv.Itoa(list.capacity))
		builder.WriteString(", size=")
	} else {
		builder.WriteString("size=")
	}

	builder.WriteString(strconv.Itoa(len(list.values)))
	builder.WriteString(", values=[")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString("]]")

	return builder.String()
}

// Prepend implements the Lister interface.
func (list *ArrayList[T]) Prepend(value T) bool {
	if list.capacity != -1 && len(list.values) >= list.capacity {
		return false
	}

	list.values = append([]T{value}, list.values...)

	return true
}

// DeleteLast implements the Lister interface.
func (list *ArrayList[T]) DeleteLast() (T, bool) {
	if len(list.values) == 0 {
		return *new(T), false
	}

	toRemove := list.values[len(list.values)-1]
	list.values = list.values[:len(list.values)-1]
	return toRemove, true
}

// PeekLast implements the Lister interface.
func (list *ArrayList[T]) PeekLast() (T, bool) {
	if len(list.values) == 0 {
		return *new(T), false
	}

	return list.values[len(list.values)-1], true
}

// Slice is a method of the ArrayList type that returns a slice of type T
// containing the elements of the list.
//
// Returns:
//
//   - []T: A slice of type T containing the elements of the list.
func (list *ArrayList[T]) Slice() []T {
	slice := make([]T, len(list.values))
	copy(slice, list.values)

	return slice
}

// Copy is a method of the ArrayList type. It is used to create a shallow copy
// of the list.
//
// Returns:
//   - *ArrayList[T]: A copy of the list.
func (list *ArrayList[T]) Copy() *ArrayList[T] {
	if list.capacity == -1 {
		l := &ArrayList[T]{
			values: make([]T, len(list.values)),
		}

		copy(l.values, list.values)

		return l
	}

	l := &ArrayList[T]{
		values:   make([]T, len(list.values), list.capacity),
		capacity: list.capacity,
	}

	copy(l.values, list.values)

	return l
}
