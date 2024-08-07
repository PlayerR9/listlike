package pkg

import (
	"errors"
	"flag"

	ggen "github.com/PlayerR9/go-generator/generator"
)

var (
	OutputLocFlag *ggen.OutputLocVal

	TypeListFlag *ggen.TypeListVal

	GenericsFlag *ggen.GenericsSignVal

	// TypeName is the name of the linked stack.
	TypeName *string
)

func init() {
	OutputLocFlag = ggen.NewOutputFlag("<type>__linkedstack.go", false)
	TypeListFlag = ggen.NewTypeListFlag("type", true, 1, "The data type of the linked stack.")
	GenericsFlag = ggen.NewGenericsSignFlag("g", false, 1)

	TypeName = flag.String("name", "", "the name of the linked stack. Must be a valid Go identifier. If not set, "+
		"the default name of 'Linked<DataType>Stack' will be used instead.")
}

func fix_type_name(data_type string) (string, error) {
	if TypeName == nil {
		return "", errors.New("the -name flag is required")
	}

	type_name := *TypeName
	if type_name != "" {
		err := ggen.IsValidVariableName(type_name, nil, ggen.Exported)
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

func ParseFlags() (string, string, error) {
	ggen.ParseFlags()

	err := ggen.AlignGenerics(GenericsFlag, TypeListFlag)
	if err != nil {
		return "", "", err
	}

	data_type, err := TypeListFlag.Type(0)
	if err != nil {
		return "", "", err
	}

	type_name, err := fix_type_name(data_type)
	if err != nil {
		return "", "", err
	}

	return data_type, type_name, nil
}
