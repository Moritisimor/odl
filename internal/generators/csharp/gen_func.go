package csharp

import (
	"fmt"
	"strings"

	"github.com/Moritisimor/odl/internal/helpers"
	"github.com/Moritisimor/odl/internal/models"
)

func GenerateCSharp(objDefs []models.ObjectDefinition) (string, error) {
	var builder strings.Builder
	CSharpTypes := map[string]string{
		"string": "string",
		"float": "double",
		"int": "int",
		"bool": "bool",
	}

	for _, obj := range objDefs {
		fmt.Fprintf(&builder, "public class %s\n{\n", helpers.PascalCase(obj.Name))
		for _, f := range obj.Fields { 
			cSharpType, ok := CSharpTypes[f.FieldType]
			if !ok {
				return builder.String(), fmt.Errorf("i could not find a C# equivalent to the type '%s'", f.FieldType)
			}

			fmt.Fprintf(
				&builder, 
				"\tpublic required %s %s { get; set; }\n", 
				cSharpType, 
				helpers.PascalCase(f.Name),
			)
		}

		fmt.Fprintf(&builder, "}\n\n")
	}

	return strings.TrimSuffix(builder.String(), "\n"), nil
}
