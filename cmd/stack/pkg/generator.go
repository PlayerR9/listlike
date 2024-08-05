package pkg

import (
	"log"
	"os"
	"strings"

	ggen "github.com/PlayerR9/lib_units/generator"
)

var (
	// Logger is the logger to use.
	Logger *log.Logger
)

func init() {
	Logger = ggen.InitLogger(os.Stdout, "linked stack")
}

type GenData struct {
	PackageName  string
	Dependencies []string
	StringFunc   string

	TypeName   string
	TypeSig    string
	HelperSig  string
	HelperName string
	Generics   string
	DataType   string
	ZeroValue  string
}

func (g *GenData) SetPackageName(name string) {
	g.PackageName = name
}

var (
	// Generator is the code Generator.
	Generator *ggen.CodeGenerator[*GenData]
)

func init() {
	tmp, err := ggen.NewCodeGeneratorFromTemplate[*GenData]("", templ)
	if err != nil {
		Logger.Fatalf("Could not initialize generator: %s", err.Error())
	}

	tmp.AddDoFunc(func(t *GenData) error {
		sig, err := ggen.MakeTypeSig(t.TypeName, "")
		if err != nil {
			return err
		}

		t.TypeSig = sig

		return nil
	})

	tmp.AddDoFunc(func(gd *GenData) error {
		data_type := strings.TrimPrefix(gd.DataType, "*")

		sig, err := ggen.MakeTypeSig("stack_node_", data_type)
		if err != nil {
			return err
		}

		gd.HelperSig = sig

		return nil
	})

	tmp.AddDoFunc(func(gd *GenData) error {
		data_type := strings.TrimPrefix(gd.DataType, "*")

		gd.HelperName = "stack_node_" + data_type

		return nil
	})

	tmp.AddDoFunc(func(gd *GenData) error {
		f_call, deps := ggen.GetStringFunctionCall("node.value", gd.DataType, nil)

		gd.StringFunc = f_call

		deps = append(deps, "strconv", "strings", "github.com/PlayerR9/lib_units/common")

		gd.Dependencies = ggen.GetPackages(deps)

		return nil
	})

	Generator = tmp
}

const templ = `// Code generated with go generate. DO NOT EDIT.
package {{ .PackageName }}

import ({{ range $index, $dep := .Dependencies }}
	"{{ $dep }}"
	{{- end }}
)

// {{ .HelperName }} is a node in the linked stack.
type {{ .HelperName }}{{ .Generics }} struct {
	value {{ .DataType }}
	next *{{ .HelperSig }}
}

// {{ .TypeName }} is a stack of {{ .DataType }} values implemented without a maximum capacity
// and using a linked list.
type {{ .TypeName }}{{ .Generics }} struct {
	front *{{ .HelperSig }}
	size int
}

// New{{ .TypeName }} creates a new linked stack.
//
// Returns:
//   - *{{ .TypeSig }}: A pointer to the newly created stack. Never returns nil.
func New{{ .TypeName }}{{ .Generics }}() *{{ .TypeSig }} {
	return &{{ .TypeSig }}{
		size: 0,
	}
}

// Push implements the stack.Stacker interface.
//
// Always returns true.
func (s *{{ .TypeSig }}) Push(value {{ .DataType }}) bool {
	node := &{{ .HelperSig }}{
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
func (s *{{ .TypeSig }}) PushMany(values []{{ .DataType }}) int {
	if len(values) == 0 {
		return 0
	}

	node := &{{ .HelperSig }}{
		value: values[0],
	}

	if s.front != nil {
		node.next = s.front
	}

	s.front = node

	for i := 1; i < len(values); i++ {
		node := &{{ .HelperSig }}{
			value: values[i],
			next:  s.front,
		}

		s.front = node
	}

	s.size += len(values)
	
	return len(values)
}

// Pop implements the stack.Stacker interface.
func (s *{{ .TypeSig }}) Pop() ({{ .DataType }}, bool) {
	if s.front == nil {
		return {{ .ZeroValue }}, false
	}

	to_remove := s.front
	s.front = s.front.next

	s.size--
	to_remove.next = nil

	return to_remove.value, true
}

// Peek implements the stack.Stacker interface.
func (s *{{ .TypeSig }}) Peek() ({{ .DataType }}, bool) {
	if s.front == nil {
		return {{ .ZeroValue }}, false
	}

	return s.front.value, true
}

// IsEmpty implements the stack.Stacker interface.
func (s *{{ .TypeSig }}) IsEmpty() bool {
	return s.front == nil
}

// Size implements the stack.Stacker interface.
func (s *{{ .TypeSig }}) Size() int {
	return s.size
}

// Iterator implements the stack.Stacker interface.
func (s *{{ .TypeSig }}) Iterator() common.Iterater[{{ .DataType }}] {
	var builder common.Builder[{{ .DataType }}]

	for node := s.front; node != nil; node = node.next {
		builder.Add(node.value)
	}

	return builder.Build()
}

// Clear implements the stack.Stacker interface.
func (s *{{ .TypeSig }}) Clear() {
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
func (s *{{ .TypeSig }}) GoString() string {
	values := make([]string, 0, s.size)
	for node := s.front; node != nil; node = node.next {
		values = append(values, {{ .StringFunc }})
	}

	var builder strings.Builder

	builder.WriteString("{{ .TypeSig }}[size=")
	builder.WriteString(strconv.Itoa(s.size))
	builder.WriteString(", values=[")
	builder.WriteString(strings.Join(values, ", "))
	builder.WriteString(" →]]")

	return builder.String()
}

// Slice implements the stack.Stacker interface.
//
// The 0th element is the top of the stack.
func (s *{{ .TypeSig }}) Slice() []{{ .DataType }} {
	slice := make([]{{ .DataType }}, 0, s.size)

	for node := s.front; node != nil; node = node.next {
		slice = append(slice, node.value)
	}

	return slice
}

// Capacity implements the stack.Stacker interface.
//
// Always returns -1.
func (s *{{ .TypeSig }}) Capacity() int {
	return -1
}

// IsFull implements the stack.Stacker interface.
//
// Always returns false.
func (s *{{ .TypeSig }}) IsFull() bool {
	return false
}

// Copy is a method that returns a deep copy of the stack.
//
// Returns:
//   - *{{ .TypeSig }}: A pointer to the newly created stack. Never returns nil.
func (s *{{ .TypeSig }}) Copy() *{{ .TypeSig }} {
	if s.front == nil {
		return &{{ .TypeSig }}{}
	}

	s_copy := &{{ .TypeSig }}{
		size: s.size,
	}

	node_copy := &{{ .HelperSig }}{
		value: s.front.value,
	}

	s_copy.front = node_copy

	prev := node_copy

	for node := s.front.next; node != nil; node = node.next {
		node_copy := &{{ .HelperSig }}{
			value: node.value,
		}

		prev.next = node_copy

		prev = node_copy
	}

	return s_copy
}`
