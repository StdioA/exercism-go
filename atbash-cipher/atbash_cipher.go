package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var builder strings.Builder
	
	for _, r := range strings.ToLower(s) {
		if unicode.IsDigit(r) {
			builder.WriteRune(r)
		} else if unicode.IsLower(r) {
			w := 25 - r + 'a' * 2
			builder.WriteRune(w)
		}
		if builder.Len() % 6 == 5 {
			builder.WriteRune(' ')
		}
	}
	return strings.Trim(builder.String(), " ")
}
