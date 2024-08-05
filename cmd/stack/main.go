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
	ggen "github.com/PlayerR9/lib_units/generator"
	pkg "github.com/PlayerR9/listlike/cmd/stack/pkg"
)

func main() {
	data_type, type_name, err := pkg.ParseFlags()
	if err != nil {
		pkg.Logger.Fatalf("Could not parse flags: %s", err.Error())
	}

	g := &pkg.GenData{
		DataType:  data_type,
		TypeName:  type_name,
		Generics:  ggen.GenericsSigFlag.String(),
		ZeroValue: ggen.ZeroValueOf(data_type, nil),
	}

	dest, err := pkg.Generator.Generate(type_name, "_linkedstack.go", g)
	if err != nil {
		pkg.Logger.Fatalf("Could not generate code: %s", err.Error())
	}

	err = dest.WriteFile()
	if err != nil {
		pkg.Logger.Fatal(err.Error())
	} else {
		pkg.Logger.Printf("Successfully generated: %q", dest.DestLoc)
	}
}
