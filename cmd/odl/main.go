package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Moritisimor/odl/internal/flags"
	"github.com/Moritisimor/odl/internal/generators/java"
	"github.com/Moritisimor/odl/internal/generators/python"
	"github.com/Moritisimor/odl/internal/generators/rust"
	"github.com/Moritisimor/odl/internal/parsing"
)

func main() {
	flags, err := flags.ParseFlags(os.Args[1:])
	if err != nil {
		fmt.Printf("Error while parsing flags: %s\n", err.Error())
		os.Exit(1)
	}

	if flags.Target == "java" && flags.Output != "" {
		fmt.Println("Transpiler target may not be 'java' when output flag is set.")
		fmt.Println("Why? Because a public Java Class expects its name to be the name of the file.")
		fmt.Println("Omit the output flag. I will generate the appropriate files for you.")
		os.Exit(1)
	}

	data, err := os.ReadFile(flags.Input)
	if err != nil {
		fmt.Printf("Error while reading '%s': %s\n", flags.Input, err.Error())
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")
	objs, err := parsing.ParseObjects(lines)
	if err != nil {
		fmt.Printf("error while parsing: %s\n", err.Error())
		os.Exit(1)
	}

	// Java gets special treatment due to how public classes work
	if flags.Target == "java" {
		files, err := java.GenerateJava(objs)
		if err != nil {
			fmt.Printf("Error while generating java source code: %s\n", err.Error())
			os.Exit(1)
		}

		for name, content := range files {
			if err := os.WriteFile(name + ".java", []byte(content), 0755); err != nil {
				fmt.Printf("Error while writing to file '%s': %s\n", name, err.Error())
				continue
			}
		}

		return
	}

	var name, content string
	name = flags.Output
	if name == "" {
		name = "out"
	}

	switch flags.Target {
	case "python", "py":
		name = strings.TrimSuffix(name, ".py") + ".py"
		content, err = python.GeneratePython(objs)

	case "rust":
		name = strings.TrimSuffix(name, ".rs") + ".rs"
		content, err = rust.GenerateRust(objs)

	default:
		fmt.Printf("Unknown target '%s'\n", flags.Target)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Error while generating %s code: %s\n", flags.Target, err.Error())
		os.Exit(1)
	}


	if err := os.WriteFile(name, []byte(content), 0755); err != nil {
		fmt.Printf("Error while writing to file '%s': %s\n", name, err.Error())
		os.Exit(1)
	}
}
