package etl

import "strings"

func Transform(m map[int][]string) map[string]int {
	result := make(map[string]int)
	for num, chars := range m {
		for _, char := range chars {
			char = strings.ToLower(char)
			result[char] = num
		}
	}
	return result
}
