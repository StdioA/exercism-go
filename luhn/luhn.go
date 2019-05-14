package luhn

import (
	"strconv"
	"unicode"
)

// Valid solution
func Valid(numbers string) bool {
	if len(numbers) <= 1 {
		return false
	}
	if numbers[0] == ' ' {
		return false
	}
	var (
		double bool
		sum    int
	)
	length := len(numbers)
	for i := length - 1; i >= 0; i-- {
		nstr := rune(numbers[i])
		if nstr == ' ' {
			continue
		} else if !unicode.IsDigit(nstr) {
			return false
		}

		n, _ := strconv.Atoi(string(nstr))
		if !double {
			sum += n
		} else {
			n *= 2
			if n > 9 {
				n -= 9
			}
			sum += n
		}
		double = !double
	}
	return sum%10 == 0
}