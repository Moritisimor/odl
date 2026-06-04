package parsing

import (
	"fmt"
	"strings"

	"github.com/Moritisimor/odl/internal/helpers"
	"github.com/Moritisimor/odl/internal/models"
)

func ParseObjects(file []string) ([]models.ObjectDefinition, error) {
	objects := []models.ObjectDefinition{}
	
	for i, line := range file {
		if strings.TrimSpace(line) == "" {
			continue
		}

		lineNumber := i + 1
		tokens := strings.Fields(line)

		if tokens[0] == "class" {
			if len(tokens) < 2 {
				return objects, fmt.Errorf("line %d: expected identifier for class", lineNumber)
			}

			name := strings.Split(tokens[1], "_")
			for _, obj := range objects {
				if helpers.PascalCase(obj.Name) == helpers.PascalCase(name) {
					return objects, fmt.Errorf("tried declaring class '%s' twice", helpers.SnakeCase(name))
				}
			}

			wasEnded := false
			class := models.ObjectDefinition{
				Name:   	strings.Split(tokens[1], "_"),
				Fields: 	[]models.FieldDefinition{},
				Options: 	tokens[2:],
			}

			for j, field := range file[i+1:] {
				if strings.TrimSpace(field) == "end" {
					wasEnded = true
					break
				}

				f, err := ParseField(strings.TrimSpace(field))
				if err != nil {
					return objects, fmt.Errorf("line %d: error while parsing class: %s", lineNumber + j + 1, err.Error())
				}

				class.Fields = append(class.Fields, f)
			}

			if !wasEnded {
				return objects, fmt.Errorf("error while parsing class %s: this class was never ended", class.Name)
			}

			objects = append(objects, class)
		}
	}

	return objects, nil
}
