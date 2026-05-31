package tests

import (
	"testing"

	"github.com/Moritisimor/odl/internal/parsing"
)

func TestParseField(t *testing.T) {
	field, err := parsing.ParseField("string name option1 option2")
	if err != nil {
		t.Errorf("Error while parsing field 1: %s", err.Error())
	}

	ExpectedName := []string{"name"}
	ExpectedType := "string"
	if field.FieldType != ExpectedType {
		t.Errorf(
			"field does not have the expected type. Expected '%s', got '%s'",
			ExpectedType, field.FieldType,
		)
	}

	if field.Name[0] != ExpectedName[0] {
		t.Errorf(
			"field does not have the expected name. Expected '%s', got '%s'",
			ExpectedName, field.Name,
		)
	}

	if len(field.Options) != 2 {
		t.Errorf("Expected 2 options, got %d", len(field.Options))
	}

	if field.Options[0] != "option1" {
		t.Errorf("First option is not as expected. Expected 'option1', got '%s'", field.Options[0])
	}

	if field.Options[1] != "option2" {
		t.Errorf("First option is not as expected. Expected 'option2', got '%s'", field.Options[1])
	}
}
