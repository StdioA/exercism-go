package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	abbrs := make([]string, 0)
	var letter bool
	for _, r := range s {
		if unicode.IsLetter(r) && !letter {
			abbrs = append(abbrs, strings.ToUpper(string(r)))
		}
		letter = (unicode.IsLetter(r) || r == '\'')
	}
	return strings.Join(abbrs, "")
}
