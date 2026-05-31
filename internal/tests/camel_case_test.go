package tests

import (
	"testing"

	"github.com/Moritisimor/odl/internal/helpers"
)

func TestCamelCase(t *testing.T) {
	name := []string{"a", "really", "cool", "class"}
	camel := helpers.CamelCase(name)
	expected := "aReallyCoolClass"

	if camel != expected {
		t.Errorf(
			"Camel Case did not produce the expected string. Expected: '%s', got: '%s",
			expected, camel,
		)
	}
}
