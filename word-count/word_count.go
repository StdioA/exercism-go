package wordcount


import (
	"strings"
	"regexp"
)


type Frequency map[string]int


var wordRegex = regexp.MustCompile(`[\w\d]([\w\d']*[\w\d])?`)

func WordCount(text string) Frequency {
	freq := make(Frequency)
	text = strings.ToLower(text)
	for _, word := range wordRegex.FindAllString(text, -1) {
		freq[word]++
	}
	return freq
}
