package util

import (
	"fmt"
	"strings"
)

// replaces whitespace chars with underscores and makes all words lower case
func ToSnakeCase(s string) string {
	snake := strings.Builder{}
	words := strings.Fields(strings.Trim(s, " "))
	for n, word := range words {
		fmt.Fprintf(&snake, "%s", strings.ToLower(word))
		if n < len(words)-1 {
			fmt.Fprint(&snake, "_")
		}
	}
	return snake.String()
}

func SplitLines(s string) []string {
	return strings.Split(s, "\n")
}
