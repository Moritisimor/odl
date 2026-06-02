package rust

import (
	"fmt"
	"strings"

	"github.com/Moritisimor/odl/internal/helpers"
	"github.com/Moritisimor/odl/internal/models"
)

func GenerateRust(objDefs []models.ObjectDefinition) (string, error) {
	var builder strings.Builder
	rustTypes := map[string]string{
		"string": "String",
		"int": "i64",
		"float": "f64",
		"bool": "bool",
	}

	for _, obj := range objDefs {
		fmt.Fprintf(&builder, "pub struct %s {\n", helpers.PascalCase(obj.Name))
		for _, f := range obj.Fields {
			rustType, ok := rustTypes[f.FieldType]
			if !ok {
				return builder.String(), fmt.Errorf("i could not find a rust equvialent to the type '%s'", f.FieldType)
			}

			fmt.Fprintf(&builder, "\tpub %s: %s,\n", helpers.SnakeCase(f.Name), rustType)
		}

		fmt.Fprintf(&builder, "}\n\n")
	}

	return strings.TrimSuffix(builder.String(), "\n"), nil
}
