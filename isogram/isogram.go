package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram solution
func IsIsogram(str string) bool {
	runeMap := make(map[rune]int)
	str = strings.ToLower(str)
	for _, r := range str {
		if !unicode.IsLetter(r) {
			continue
		}
		runeMap[r]++
		if runeMap[r] > 1 {
			return false
		}
	}
	return true
}
