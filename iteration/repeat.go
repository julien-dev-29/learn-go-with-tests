package iteration

import "strings"

func Repeat(character string, repeatedCount int) string {
	var repeated strings.Builder
	for range repeatedCount {
		repeated.WriteString(character)
	}
	return repeated.String()
}
