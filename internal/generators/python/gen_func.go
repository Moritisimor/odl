package python

import (
	"fmt"
	"strings"

	"github.com/Moritisimor/odl/internal/helpers"
	"github.com/Moritisimor/odl/internal/models"
)

func GeneratePython(objDefs []models.ObjectDefinition) (string, error) {
	var builder strings.Builder
	pythonTypes := map[string]string{
		"string": "str",
		"bool": "bool",
		"int": "int",
		"float": "float",
	}

	for _, obj := range objDefs {
		fmt.Fprintf(&builder, "class %s:\n", helpers.PascalCase(obj.Name))
		fmt.Fprintf(&builder, "\tdef __init__(self, ")
		for i, f := range obj.Fields {
			pythonType, ok := pythonTypes[f.FieldType]
			if !ok {
				return builder.String(), fmt.Errorf("i could not find a python equivalent to the type '%s'", f.FieldType)
			}

			fmt.Fprintf(&builder, "%s: %s", helpers.SnakeCase(f.Name), pythonType)
			if i < len(obj.Fields) - 1 {
				fmt.Fprint(&builder, ", ")
			}
		}

		fmt.Fprintf(&builder, "):")
		for _, f := range obj.Fields {
			fmt.Fprintf(
				&builder, 
				"\n\t\tself.%s = %s", 
				helpers.SnakeCase(f.Name), 
				helpers.SnakeCase(f.Name),
			)
		}
	}

	builder.WriteString("\n")
	return builder.String(), nil
}
