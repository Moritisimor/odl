package parsing

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Moritisimor/odl/internal/models"
)

func ParseField(line string, allowedTypes []string) (models.FieldDefinition, error) {
	fieldDef := models.FieldDefinition{}
	tokens := strings.Fields(line)
	if len(tokens) < 2 {
		return fieldDef, fmt.Errorf("malformed line (expected at least 2 tokens)")
	}

	if !slices.Contains(allowedTypes, tokens[0]) {
		return fieldDef, fmt.Errorf("%s is not a valid type", tokens[0])
	}

	fieldDef.FieldType = tokens[0]
	fieldDef.Name = strings.Split(tokens[1], "_")
	fieldDef.Options = tokens[2:]
	return fieldDef, nil
} 
