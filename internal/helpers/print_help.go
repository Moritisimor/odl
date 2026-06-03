package helpers

import "fmt"

func PrintHelp() {
	fmt.Println("Object Definition Language.")
	fmt.Println("A domain-specific-language for defining objects and transpiling them to other languages.")
	fmt.Println("https://github.com/Moritisimor/odl")

	fmt.Println("\nUsage: <input file> <options...>")
	fmt.Println("Options: ")
	fmt.Println("\t-t | --target <target language> (REQUIRED)")
	fmt.Println("\t-o | --output <output file> (OPTIONAL)")
	fmt.Println("\t\tNote: --output flag must be omitted when transpiling for Java.")

	fmt.Println("\nAvailable Targets:")
	fmt.Println("\t- Java\n\t- Python\n\t- C#\n\t- Go\n\t- Rust")

	fmt.Println("\nGeneral Syntax:")
	fmt.Println("\tclass <name> <options...>")
	fmt.Println("\t\t<type> <name> <options...>")
	fmt.Println("\tend")

	fmt.Println("\nExample:")
	fmt.Println("\tclass person")
	fmt.Println("\t\tstring name")
	fmt.Println("\t\tint age")
	fmt.Println("\t\tfloat monthly_salary")
	fmt.Println("\tend")
}
