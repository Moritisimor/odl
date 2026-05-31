package tests

import (
	"testing"

	"github.com/Moritisimor/odl/internal/helpers"
)

func TestPascalCase(t *testing.T) {
	name := []string{"a", "really", "cool", "class"}
	pascal := helpers.PascalCase(name)
	expected := "AReallyCoolClass"

	if pascal != expected {
		t.Errorf(
			"Pascal Case did not produce the expected string. Expected: '%s', got: '%s",
			expected, pascal,
		)
	}
}
