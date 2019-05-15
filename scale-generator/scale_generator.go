package scale

import (
	"strings"
	"unicode"
)


type scale []int

func (s *scale) ToString(tonic string) []string {
	var notation int
	pitch := parsePitch(tonic)

	// Deal with minor scale
	if unicode.IsLower(rune(tonic[0])) {
		pitch += 3
	}
	// Deal with notation
	if len(tonic) == 2 {
		if tonic[1] == '#' {
			notation = 0
		} else if tonic[1] == 'b' {
			notation = 1
		}
	} else {
		notation = notationMap[pitch]
	}

	length := len(*s)
	stringScale := make([]string, length, length)
	for i, t := range *s {
		var choice string
		choices := stringMap[t]
		choice = choices[0]
		if len(choices) == 2 && notation == 1 {
			choice = choices[1]
		}
		stringScale[i] = choice
	}
	return stringScale
}

func parsePitch(pitchStr string) int {
	pitch := pitchMap[strings.ToUpper(pitchStr[:1])]

	if len(pitchStr) == 2 {
		if pitchStr[1] == '#' {
			pitch++
		} else if pitchStr[1] == 'b' {
			pitch--
		}
	}
	return pitch % 12
}

// Scale generator ISN'T FUN AT ALL!
func Scale(tonic, interval string) []string {
	pitch := parsePitch(tonic)
	if interval == "" {
		interval = strings.Repeat("m", 12)
	}

	scalePitch := make(scale, len(interval))
	scalePitch[0] = pitch
	for index, inter := range interval[:len(interval)-1] {
		switch inter {
		case 'm':
			// half tone
			pitch++
		case 'M':
			// whole tone
			pitch += 2
		case 'A':
			// argumented first interval: 3/2 tone
			pitch += 3
		}
		scalePitch[index+1] = pitch % 12
	}
	return scalePitch.ToString(tonic)
}
