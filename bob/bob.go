package bob

import (
	"regexp"
	"strings"
)

var rgx = regexp.MustCompile("^[^a-z]*[A-Z][^a-z]*$")

// Hey hey hey
func Hey(remark string) string {
	type toneType struct {
		exclamation, question bool
	}
	var (
		tone  toneType
		reply string
	)
	remark = strings.Trim(remark, " \t\r\n")
	if remark == "" {
		return "Fine. Be that way!"
	}

	if rgx.MatchString(remark) {
		tone.exclamation = true
	}
	for i := len(remark) - 1; i >= 0; i-- {
		r := remark[i]
		if r == '!' {
		} else if r == '?' {
			tone.question = true
		} else {
			break
		}
	}

	switch tone {
	case toneType{false, false}:
		reply = "Whatever."
	case toneType{true, false}:
		reply = "Whoa, chill out!"
	case toneType{false, true}:
		reply = "Sure."
	case toneType{true, true}:
		reply = "Calm down, I know what I'm doing!"
	}

	return reply
}
