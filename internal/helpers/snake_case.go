package helpers

import "strings"

func SnakeCase(name []string) string {
	var buf strings.Builder
	for _, word := range name {
		buf.WriteString(strings.ToLower(word))
		buf.WriteString("_")
	}

	return strings.TrimSuffix(buf.String(), "_")
}
