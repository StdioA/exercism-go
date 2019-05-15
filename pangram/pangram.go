package pangram

import (
	"strings"
)

const vocabulary = "abcdefghijklmnopqrstuvwxyz"

func IsPangram(str string) bool {
	occurance := make(map[rune]int)
	for _, r := range strings.ToLower(str) {
		occurance[r]++
	}
	for _, char := range vocabulary {
		if occurance[char] == 0 {
			return false
		}
	}
	return true
}
