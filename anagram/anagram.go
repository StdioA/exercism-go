package anagram

import "strings"

type occrCount = map[rune]int

func count(str string) occrCount {
	c := make(occrCount)
	for _, r := range str {
		c[r]++
	}
	return c
}

func countEqual(a, b occrCount) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

func Detect(subject string, candidates []string) []string {
	available := make([]string, 0)
	subject = strings.ToLower(subject)
	subjCount := count(subject)
	for _, cand := range candidates {
		lowerCand := strings.ToLower(cand)
		if countEqual(subjCount, count(lowerCand)) && lowerCand != subject {
			available = append(available, cand)
		}
	}
	return available
}
