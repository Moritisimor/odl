package golang

import (
	"fmt"
	"strings"

	"github.com/Moritisimor/odl/internal/helpers"
	"github.com/Moritisimor/odl/internal/models"
)

func GenerateGo(objDefs []models.ObjectDefinition) (string, error) {
	var builder strings.Builder
	goTypes := map[string]string{
		"string": "string",
		"int": "int",
		"float": "float64",
		"bool": "bool",
	}

	fmt.Fprintf(&builder, "package main\n\n")
	for _, obj := range objDefs {
		fmt.Fprintf(&builder, "type %s struct{\n", helpers.PascalCase(obj.Name))
		for _, f := range obj.Fields {
			goType, ok := goTypes[f.FieldType]
			if !ok {
				return builder.String(), fmt.Errorf("i could not find a go equivalent to the type '%s'", f.FieldType)
			}

			fmt.Fprintf(&builder, "\t%s %s\n", helpers.PascalCase(f.Name), goType)
		}

		fmt.Fprintf(&builder, "}\n\n")
	}

	return strings.TrimSuffix(builder.String(), "\n"), nil
}
