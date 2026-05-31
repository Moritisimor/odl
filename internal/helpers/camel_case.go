package helpers

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CamelCase(name []string) string {
	var builder strings.Builder
	caser := cases.Title(language.English)

	for i, word := range name {
		if i == 0 {
			builder.WriteString(strings.ToLower(word))
		} else {
			builder.WriteString(caser.String(word))
		}
	}

	return builder.String()
}
