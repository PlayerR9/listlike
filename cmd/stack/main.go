// This command generates a linked stack with the specified type.
//
// To use it, run the following command:
//
// //go:generate go run stack/cmd -name=<type_name> -type=<type> [ -g=<generics> ] [ -o=<output_file> ]
//
// **Flag: Type Name**
//
// The "type name" flag is used to specify the name of the linked stack struct. If not set, the default name
// of "Linked<DataType>Stack" will be used instead; where <DataType> is the data type of the linked stack. Otherwise,
// it must be a valid Go identifier and starting with an upper case letter.
//
// **Flag: Type**
//
// The "type" flag is used to specify the type of the linked stack contains. Because it doesn't make
// a lot of sense to have a linked stack without a type, this flag must be set.
//
// For instance, running the following command:
//
// //go:generate go run stack/cmd -name=Stack -type=string
//
// will generate a linked stack with the following fields:
//
//	type Stack struct {
//		// stack of strings
//	}
//
// Also, it is possible to specify generics by following the value with the generics between square brackets;
// like so: "MyType[T,C]"
//
// **Flag: Generics**
//
// This optional flag is used to specify the type(s) of the generics. However, this only applies if at least one
// generic type is specified in the type flag. If none, then this flag is ignored.
//
// As an edge case, if this flag is not specified but the type flag contains generics, then
// all generics are set to the default value of "any".
//
// As with the fields flag, its argument is specified as a list of key-value pairs where each pair is separated
// by a comma (",") and a slash ("/") is used to separate the key and the value. The key indicates the name of
// the generic and the value indicates the type of the generic.
//
// For instance, running the following command:
//
// //go:generate go run stack/cmd -type=Stack -type=MyType[T] -g=T/any
//
// will generate a linked stack with the following fields:
//
//	type Stack[T any] struct {
//	   // stack of MyType[T]
//	}
//
// **Flag: Output File**
//
// This optional flag is used to specify the output file. If not specified, the output will be written to
// standard output, that is, the file "<type_name>_stack.go" in the root of the current directory.
package main

import (
	"flag"
	"log"
	"strings"

	uc "github.com/PlayerR9/lib_units/common"
	ggen "github.com/PlayerR9/lib_units/generator"
)

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
	// Logger is the logger to use.
	Logger *log.Logger

	// generator is the code generator.
	generator *ggen.CodeGenerator[*GenData]
)

func init() {
	Logger = ggen.InitLogger("linked stack")

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

	generator = tmp
}

var (
	// TypeName is the name of the linked stack.
	TypeName *string
)

func init() {
	ggen.SetOutputFlag("<type>__linkedstack.go", false)
	ggen.SetTypeListFlag("type", true, 1, "The data type of the linked stack.")
	ggen.SetGenericsSignFlag("g", false, 1)

	TypeName = flag.String("name", "", "the name of the linked stack. Must be a valid Go identifier. If not set, "+
		"the default name of 'Linked<DataType>Stack' will be used instead.")
}

func main() {
	err := ggen.ParseFlags()
	if err != nil {
		Logger.Fatalf("Invalid flags: %s", err.Error())
	}

	data_type, err := ggen.TypeListFlag.GetType(0)
	if err != nil {
		Logger.Fatalf("Could not get type: %s", err.Error())
	}

	type_name, err := FixTypeName(data_type)
	if err != nil {
		Logger.Fatalf("Could not fix type name: %s", err.Error())
	}

	g := &GenData{
		DataType:  data_type,
		TypeName:  type_name,
		Generics:  ggen.GenericsSigFlag.String(),
		ZeroValue: ggen.ZeroValueOf(data_type, nil),
	}

	err = generator.Generate(type_name, "_linkedstack.go", g)
	if err != nil {
		Logger.Fatalf("Could not generate code: %s", err.Error())
	}

	Logger.Printf("Successfully generated: %q", type_name+"_linkedstack.go")
}

func FixTypeName(data_type string) (string, error) {
	type_name := uc.AssertDerefNil(TypeName, "TypeName")
	if type_name != "" {
		err := ggen.IsValidName(type_name, nil, ggen.Exported)
		if err != nil {
			return "", err
		}

		return type_name, nil
	}

	data_type, err := ggen.FixVariableName(data_type, nil, ggen.Exported)
	if err != nil {
		return "", err
	}

	type_name = "Linked" + data_type + "Stack"

	return type_name, nil
}

const templ = `// Code generated with go generate. DO NOT EDIT.
package {{ .PackageName }}

import (
	{{ range $index, $dep := .Dependencies }}
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
		values = append(values, {{ .StringFunc }}
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

// Copy implements the stack.Stacker interface.
//
// The copy is a shallow copy.
func (s *{{ .TypeSig }}) Copy() common.Copier {
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
`
