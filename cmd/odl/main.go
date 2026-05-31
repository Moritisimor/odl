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
	if len(args) < 3 {
		fmt.Printf("Usage: odl <file_name> <transpiler_target>\n")
		os.Exit(1)
	}

	fileName := args[1]
	transpilerTarget := args[2]

	data, err := os.ReadFile(fileName)
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

	var files map[string]string
	switch transpilerTarget {
	case "java":
		files, err = java.GenerateJava(objs)

	default:
		fmt.Printf("Unknown transpiler target '%s'\n", transpilerTarget)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Error while generating %s source code: %s", transpilerTarget, err.Error())
		os.Exit(1)
	}

	for className, fileContent := range files {
		var fileName string
		switch transpilerTarget {
		case "java":
			fileName = className + ".java"

		default:
			fileName = className
		}

		err := os.WriteFile(fileName, []byte(fileContent), 0755)
		if err != nil {
			fmt.Printf("Could not write generator output to '%s': %s\n", fileName, err.Error())
			continue
		}
	}
}
