package flags

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/Moritisimor/odl/internal/helpers"
)

func ParseFlags(args []string) (Flags, error) {
	flags := Flags{}
	ignoreable := []int{}

	if slices.Contains(args, "--help") {
		helpers.PrintHelp()
		os.Exit(0)
	}

	for idx, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			if slices.Contains(ignoreable, idx) {
				continue
			}

			if flags.Input != "" {
				return flags, fmt.Errorf("Tried to set input file twice")
			}

			flags.Input = arg
		}

		switch arg {
		case "-t", "--target":
			if len(args) < idx + 2 {
				return flags, fmt.Errorf("Expected value after '%s'", arg)
			}

			flags.Target = args[idx + 1]
			ignoreable = append(ignoreable, idx + 1)

		case "-o", "--output":
			if len(args) < idx + 2 {
				return flags, fmt.Errorf("Expected value after '%s'", arg)
			}

			flags.Output = args[idx + 1]
			ignoreable = append(ignoreable, idx + 1)

		default:
			return flags, fmt.Errorf("Unknown flag: '%s'", arg)
		}
	}

	if flags.Input == "" {
		return flags, fmt.Errorf("Input file not provided")
	}

	if flags.Target == "" {
		return flags, fmt.Errorf("Required Flag '--target' is not set.")
	}

	return flags, nil
}
