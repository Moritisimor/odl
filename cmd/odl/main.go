package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Moritisimor/odl/internal/generators/java"
	"github.com/Moritisimor/odl/internal/parsing"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: odl <file_name>\n")
		os.Exit(1)
	}

	data, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Printf("Error while reading '%s': %s\n", args[1], err.Error())
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")
	legalTypes := []string{
		"string",
		"int",
		"float",
		"bool",
	}

	objs, err := parsing.ParseObjects(lines, legalTypes)
	if err != nil {
		fmt.Printf("error while parsing: %s\n", err.Error())
		os.Exit(1)
	}

	javaFiles, err := java.GenerateJava(objs)
	for fileName, fileContent := range javaFiles {
		fmt.Printf("==( %s )==\n%s", fileName, fileContent)
	}
}
