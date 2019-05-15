package isbn

import (
	"regexp"
	"strings"
)

var reg = regexp.MustCompile("^[0-9]{9}[0-9X]$")

func IsValidISBN(isbn string) bool {
	isbn = strings.Replace(isbn, "-", "", -1)
	if !reg.MatchString(isbn) {
		return false
	}

	sum := 0
	for i := 0; i < 10; i++ {
		var n int
		if isbn[i] == 'X' {
			n = 10
		} else {
			n = int(isbn[i] - '0')
		}
		sum += n * (10 - i)
	}
	return sum%11 == 0
}
