package list

import (
	"strconv"
	"strings"
	"sync"

	gcstr "github.com/PlayerR9/go-commons/strings"
	itrs "github.com/PlayerR9/iterators/simple"
)

// LimitedSafeList is a generic type that represents a thread-safe list data
// structure with or without a maximum capacity, implemented using a linked list.
type LimitedSafeList[T any] struct {
	// front and back are pointers to the first and last nodes in the safe list,
	// respectively.
	front, back *ListSafeNode[T]

	// frontMutex and backMutex are sync.RWMutexes, which are used to ensure that
	// concurrent reads and writes to the front and back nodes are thread-safe.
	frontMutex, backMutex sync.RWMutex

	// size is the current number of elements in the list.
	size int

	// capacity is the maximum number of elements that the list can hold.
	capacity int
}

// NewLimitedSafeList is a function that creates and returns a new instance of a
// LimitedSafeList.
//
// Parameters:
//
//   - values: A variadic parameter of type T, which represents the initial values to
//     be stored in the list.
//
// Returns:
//
//   - *LimitedSafeList[T]: A pointer to the newly created LimitedSafeList.
func NewLimitedSafeList[T any](values ...T) *LimitedSafeList[T] {
	list := new(LimitedSafeList[T])

	if len(values) == 0 {
		return list
	}

	list.size = len(values)

	// First node
	list_node := NewListSafeNode(values[0])

	list.front = list_node
	list.back = list_node

	// Subsequent nodes
	for _, element := range values {
		list_node := NewListSafeNode(element)

		list_node.SetPrev(list.back)

		list.back.SetNext(list_node)
		list.back = list_node
	}

	return list
}

// Append implements the Lister interface.
func (list *LimitedSafeList[T]) Append(value T) bool {
	list.backMutex.Lock()
	defer list.backMutex.Unlock()

	if list.size >= list.capacity {
		return false
	}

	node := NewListSafeNode(value)

	if list.back != nil {
		list.back.SetNext(node)
		node.SetPrev(list.back)
	} else {
		// The list is empty
		list.frontMutex.Lock()
		list.front = node
		list.frontMutex.Unlock()
	}

	list.back = node

	list.size++

	return true
}

// DeleteFirst implements the Lister interface.
func (list *LimitedSafeList[T]) DeleteFirst() (T, bool) {
	list.frontMutex.Lock()
	defer list.frontMutex.Unlock()

	if list.front == nil {
		return *new(T), false
	}

	toRemove := list.front

	list.backMutex.Lock()

	list.front = list.front.Next()

	if list.front == nil {
		list.back = nil
	} else {
		list.front.SetPrev(nil)
	}

	list.backMutex.Unlock()

	list.size--

	toRemove.SetNext(nil)

	return toRemove.Value, true
}

// PeekFirst implements the Lister interface.
func (list *LimitedSafeList[T]) PeekFirst() (T, bool) {
	list.frontMutex.RLock()
	defer list.frontMutex.RUnlock()

	if list.front == nil {
		return *new(T), false
	}

	return list.front.Value, true
}

// IsEmpty is a method of the LimitedSafeList type. It checks if the list is empty.
//
// Returns:
//
//   - bool: A boolean value that is true if the list is empty, and false otherwise.
func (list *LimitedSafeList[T]) IsEmpty() bool {
	list.frontMutex.RLock()
	defer list.frontMutex.RUnlock()

	return list.front == nil
}

// Size is a method of the LimitedSafeList type. It returns the number of elements in the
// list.
//
// Returns:
//
//   - int: An integer that represents the number of elements in the list.
func (list *LimitedSafeList[T]) Size() int {
	list.frontMutex.RLock()
	defer list.frontMutex.RUnlock()

	list.backMutex.RLock()
	defer list.backMutex.RUnlock()

	return list.size
}

// Capacity is a method of the LimitedSafeList type. It returns the maximum number of
// elements that the list can hold.
//
// Returns:
//
//   - optional.Int: An optional integer that represents the maximum number of
//     elements the list can hold.
func (list *LimitedSafeList[T]) Capacity() int {
	return list.capacity
}

// Iterator is a method of the LimitedSafeList type. It is used to return an iterator
// for the list.
// However, the iterator does not share the list's thread safety.
//
// Returns:
//
//   - uc.Iterater[T]: An iterator for the list.
func (list *LimitedSafeList[T]) Iterator() itrs.Iterater[T] {
	list.frontMutex.RLock()
	defer list.frontMutex.RUnlock()

	list.backMutex.RLock()
	defer list.backMutex.RUnlock()

	var builder itrs.Builder[T]

	for node := list.front; node != nil; node = node.Next() {
		builder.Add(node.Value)
	}

	return builder.Build()
}

// Clear is a method of the LimitedSafeList type. It is used to remove all elements from
// the list.
func (list *LimitedSafeList[T]) Clear() {
	list.frontMutex.Lock()
	defer list.frontMutex.Unlock()

	list.backMutex.Lock()
	defer list.backMutex.Unlock()

	if list.front == nil {
		return // List is already empty
	}

	// 1. First node
	list.front.SetPrev(nil)
	prev := list.front

	// 2. Subsequent nodes
	for node := list.front.Next(); node != nil; node = node.Next() {
		node.SetPrev(nil)

		prev = node
		prev.SetNext(nil)
	}

	prev.SetNext(nil)

	// 3. Reset list fields
	list.front = nil
	list.back = nil
	list.size = 0
}

// IsFull is a method of the LimitedSafeList type. It checks if the list is fu
//
// Returns:
//
//   - isFull: A boolean value that is true if the list is full, and false otherwise.
func (list *LimitedSafeList[T]) IsFull() (isFull bool) {
	return list.capacity <= list.size
}

// GoString implements the fmt.GoStringer interface.
func (list *LimitedSafeList[T]) GoString() string {
	list.frontMutex.RLock()
	defer list.frontMutex.RUnlock()

	list.backMutex.RLock()
	defer list.backMutex.RUnlock()

	values := make([]string, 0, list.size)
	for node := list.front; node != nil; node = node.Next() {
		values = append(values, gcstr.GoStringOf(node.Value))
	}

	var builder strings.Builder

	builder.WriteString("LimitedSafeList[capacity=")
	builder.WriteString(strconv.Itoa(list.capacity))
	builder.WriteString(", size=")
	builder.WriteString(strconv.Itoa(list.size))
	builder.WriteString(", values=[")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString("]]")

	return builder.String()
}

// Prepend implements the Lister interface.
func (list *LimitedSafeList[T]) Prepend(value T) bool {
	list.frontMutex.Lock()
	defer list.frontMutex.Unlock()

	if list.size >= list.capacity {
		return false
	}

	node := NewListSafeNode(value)

	if list.front == nil {
		// The list is empty
		list.backMutex.Lock()
		list.back = node
		list.backMutex.Unlock()
	} else {
		node.SetNext(list.front)
		list.front.SetPrev(node)
	}

	list.front = node

	list.size++

	return true
}

// DeleteLast implements the Lister interface.
func (list *LimitedSafeList[T]) DeleteLast() (T, bool) {
	list.backMutex.Lock()
	defer list.backMutex.Unlock()

	if list.back == nil {
		return *new(T), false
	}

	toRemove := list.back

	list.frontMutex.Lock()

	list.back = list.back.Prev()

	if list.back == nil {
		list.front = nil
	} else {
		list.back.SetNext(nil)
	}

	list.frontMutex.Unlock()

	list.size--

	toRemove.SetPrev(nil)

	return toRemove.Value, true
}

// PeekLast implements the Lister interface.
func (list *LimitedSafeList[T]) PeekLast() (T, bool) {
	list.backMutex.RLock()
	defer list.backMutex.RUnlock()

	if list.back == nil {
		return *new(T), false
	}

	return list.back.Value, true
}

// Slice is a method of the LimitedSafeList type. It is used to return a slice of the
// elements in the list.
//
// Returns:
//
//   - []T: A slice of type T containing the elements of the list.
func (list *LimitedSafeList[T]) Slice() []T {
	list.frontMutex.RLock()
	defer list.frontMutex.RUnlock()

	list.backMutex.RLock()
	defer list.backMutex.RUnlock()

	slice := make([]T, 0, list.size)

	for node := list.front; node != nil; node = node.Next() {
		slice = append(slice, node.Value)
	}

	return slice
}

// Copy is a method of the LimitedSafeList type. It is used to create a shallow copy of
// the list.
//
// Returns:
//   - *LimitedSafeList[T]: A copy of the list.
func (list *LimitedSafeList[T]) Copy() *LimitedSafeList[T] {
	list.frontMutex.RLock()
	defer list.frontMutex.RUnlock()

	list.backMutex.RLock()
	defer list.backMutex.RUnlock()

	list_copy := &LimitedSafeList[T]{
		size:     list.size,
		capacity: list.capacity,
	}

	if list.front == nil {
		return list_copy
	}

	// First node
	node := NewListSafeNode(list.front.Value)

	list_copy.front = node

	prev := list_copy.front

	// Subsequent nodes
	for node := list.front.Next(); node != nil; node = node.Next() {
		nodeCopy := NewListSafeNode(node.Value)
		nodeCopy.SetPrev(prev)

		prev.SetNext(nodeCopy)
		prev = nodeCopy
	}

	if list_copy.front.Next() != nil {
		list_copy.front.Next().SetPrev(list_copy.front)
	}

	list_copy.back = prev

	return list_copy
}
