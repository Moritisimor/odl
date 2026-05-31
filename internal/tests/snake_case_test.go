package tests

import (
	"testing"

	"github.com/Moritisimor/odl/internal/helpers"
)

func TestSnakeCase(t *testing.T) {
	name := []string{"a", "really", "cool", "class"}
	snake := helpers.SnakeCase(name)
	expected := "a_really_cool_class"

	if snake != expected {
		t.Errorf(
			"Snake Case did not produce the expected string. Expected: '%s', got: '%s",
			expected, snake,
		)
	}
}
