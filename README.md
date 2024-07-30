# listlike
ListLike is a Go package that contains lists, stacks, and queues. As well as generators for them and some common functions.

# stack
A Go package used for generating linked stacks. It also features some already generated stacks and operations on stacks.


## Table of Contents

1. [Table of Contents](#table-of-contents)
2. [Tool](#tool)
   - [Installation](#installation)
   - [Usage](#usage)
3. [Documentation](#documentation)
4. [Content](#content)


## Tool

### Installation

To install the tool, run the following command:
```
go get -u github.com/PlayerR9/stack/cmd
```


### Usage

Once imported, you can use the tool to generate the linked stack for your own types. Like so:
```go
import _ "github.com/PlayerR9/stack"

//go:generate go run stack/cmd -name=Foo -type=int
```

This will generate a linked stack with the name "Foo" of type "int".

The type generated will be in the same package as the tool. Make sure to read the documentation of the tool before using it.


## Documentation

```markdown
This command generates a linked stack with the specified type.

To use it, run the following command:

//go:generate go run stack/cmd -name=<type_name> -type=<type> [ -g=<generics> ] [ -o=<output_file> ]


**Flag: Type Name**

The "type name" flag is used to specify the name of the linked stack struct. If not set, the default name
of "Linked<DataType>Stack" will be used instead; where <DataType> is the data type of the linked stack. Otherwise,
it must be a valid Go identifier and starting with an upper case letter.


**Flag: Type**

The "type" flag is used to specify the type of the linked stack contains. Because it doesn't make
a lot of sense to have a linked stack without a type, this flag must be set.

For instance, running the following command:

//go:generate go run stack/cmd -name=Stack -type=string

will generate a linked stack with the following fields:

type Stack struct {
	// stack of strings
}

Also, it is possible to specify generics by following the value with the generics between square brackets;
like so: "MyType[T,C]"


**Flag: Generics**

This optional flag is used to specify the type(s) of the generics. However, this only applies if at least one
generic type is specified in the type flag. If none, then this flag is ignored.

As an edge case, if this flag is not specified but the type flag contains generics, then
all generics are set to the default value of "any".

As with the fields flag, its argument is specified as a list of key-value pairs where each pair is separated
by a comma (",") and a slash ("/") is used to separate the key and the value. The key indicates the name of
the generic and the value indicates the type of the generic.

For instance, running the following command:

//go:generate go run stack/cmd -type=Stack -type=MyType[T] -g=T/any

will generate a linked stack with the following fields:

type Stack[T any] struct {
   // stack of MyType[T]
}


**Flag: Output File**

This optional flag is used to specify the output file. If not specified, the output will be written to
standard output, that is, the file "<type_name>_stack.go" in the root of the current directory.
```


## Content

Here are all the pregenerated files:
- [bool](bool.go)
- [byte](byte.go)
- [int](int.go)
- [int8](int8.go)
- [int16](int16.go)
- [int32](int32.go)
- [int64](int64.go)
- [float32](float32.go)
- [float64](float64.go)
- [rune](rune.go)
- [string](string.go)
- [uint](uint.go)
- [uint8](uint8.go)
- [uint16](uint16.go)
- [uint32](uint32.go)
- [uint64](uint64.go)
- [uintptr](uintptr.go)
- [error](error.go)
- [complex128](complex128.go)
- [complex64](complex64.go)
- [generic](generic.go)
