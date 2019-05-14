package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	segs := make([]string, 0, 3)
	if number%3 == 0 {
		segs = append(segs, "Pling")
	}
	if number%5 == 0 {
		segs = append(segs, "Plang")
	}
	if number%7 == 0 {
		segs = append(segs, "Plong")
	}
	result := strings.Join(segs, "")
	if result == "" {
		result = strconv.Itoa(number)
	}
	return result
}
