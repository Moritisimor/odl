package helpers

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func PascalCase(name []string) string {
	var builder strings.Builder
	caser := cases.Title(language.English)

	for _, word := range name {
		builder.WriteString(caser.String(word))
	}

	return builder.String()
}
