package strand


import "strings"

var RNAMap = map[rune]rune{
	'C': 'G',
	'G': 'C',
	'T': 'A',
	'A': 'U',
}


func ToRNA(dna string) string {
	var builder strings.Builder
	for _, d := range dna {
		builder.WriteRune(RNAMap[d])
	}
	return builder.String()
}
