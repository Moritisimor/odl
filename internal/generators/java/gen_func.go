package java

import (
	"fmt"
	"strings"

	"github.com/Moritisimor/odl/internal/helpers"
	"github.com/Moritisimor/odl/internal/models"
)

// GenerateJava takes a slice of object definitions and
// returns (map[string]string, error), where the key of
// the map is the name of the file (class) and the value
// is the actual file
func GenerateJava(objDefs []models.ObjectDefinition) (map[string]string, error) {
	files := map[string]string{}
	javaTypes := map[string]string{
		"string": "String",
		"int":    "int",
		"float":  "double",
		"bool":   "boolean",
	}

	for _, obj := range objDefs {
		var builder strings.Builder

		// Class Name
		builder.WriteString("public class ")
		builder.WriteString(helpers.PascalCase(obj.Name))
		builder.WriteString(" {\n")

		// Fields
		for _, f := range obj.Fields {
			javaType, ok := javaTypes[f.FieldType]
			if !ok {
				return files, fmt.Errorf("i could not find a java equivalent to the type '%s'", f.FieldType)
			}

			fmt.Fprintf(
				&builder,
				"\tprivate %s %s;\n",
				javaType, helpers.CamelCase(f.Name),
			)
		}

		// Getters
		for _, f := range obj.Fields {
			javaType, ok := javaTypes[f.FieldType]
			if !ok {
				return files, fmt.Errorf("i could not find a java equivalent to the type '%s'", f.FieldType)
			}

			if javaType == "boolean" && f.Name[0] == "is" {
				fmt.Fprintf(
					&builder,
					"\n\tpublic boolean %s() {\n\t\treturn this.%s;\n\t}\n",
					helpers.CamelCase(f.Name), helpers.CamelCase(f.Name),
				)
			} else {
				fmt.Fprintf(
					&builder,
					"\n\tpublic %s get%s() {\n\t\treturn this.%s;\n\t}\n",
					javaType, helpers.PascalCase(f.Name), helpers.CamelCase(f.Name),
				)
			}
		}

		// Setters
		for _, f := range obj.Fields {
			javaType, ok := javaTypes[f.FieldType]
			if !ok {
				return files, fmt.Errorf("i could not find a java equivalent to the type '%s'", f.FieldType)
			}

			fmt.Fprintf(
				&builder,
				"\n\tpublic void set%s(%s %s) {\n\t\tthis.%s = %s;\n\t}\n",
				helpers.PascalCase(f.Name),
				javaType,
				helpers.CamelCase(f.Name),
				helpers.CamelCase(f.Name),
				helpers.CamelCase(f.Name),
			)
		}

		// Constructor
		fmt.Fprintf(
			&builder,
			"\n\tpublic %s(\n",
			helpers.PascalCase(obj.Name),
		)

		for i, f := range obj.Fields {
			javaType, ok := javaTypes[f.FieldType]
			if !ok {
				return files, fmt.Errorf("i could not find a java equivalent to the type '%s'", f.FieldType)
			}

			fmt.Fprintf(
				&builder,
				"\t\t%s %s",
				javaType, helpers.CamelCase(f.Name),
			)

			if i < len(obj.Fields)-1 {
				builder.WriteString(",")
			}

			builder.WriteString("\n")
		}

		builder.WriteString("\t) {\n")
		for _, f := range obj.Fields {
			fmt.Fprintf(
				&builder,
				"\t\tthis.%s = %s;\n",
				helpers.CamelCase(f.Name),
				helpers.CamelCase(f.Name),
			)
		}

		builder.WriteString("\t}\n}\n")
		files[helpers.PascalCase(obj.Name)] = builder.String()
	}

	return files, nil
}
