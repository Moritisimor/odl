package tests

import (
	"fmt"
	"testing"

	"github.com/Moritisimor/odl/internal/parsing"
)

func TestParseObject(t *testing.T) {
	legalTypes := []string{
		"string",
		"int",
		"float",
		"bool",
	}

	goodLines := []string{
		"class person",
		"    string name foo bar",
		"    int age foo",
		"    float salary bar",
		"    bool unemployed baz",
		"end",
	}

	neverEndedLines := []string{
		"class hello",
		"    string foo",
		"    float bar",
	}

	objs, err := parsing.ParseObjects(goodLines, legalTypes)
	if err != nil {
		t.Errorf("Error while parsing: %s", err.Error())
	}

	_, err = parsing.ParseObjects(neverEndedLines, legalTypes)
	if err == nil {
		t.Errorf("Error expected while parsing lines that were never ended, got nil")
	}

	for _, obj := range objs {
		for i, field := range obj.Fields {
			fmt.Printf("Field %d\n\t-> Name: %s\n\t-> Type: %s\n\t-> Options:\n", i+1, field.Name, field.FieldType)
			for j, opt := range field.Options {
				fmt.Printf("\t\t-> %d: %s\n", j+1, opt)
			}
		}
	}
}
