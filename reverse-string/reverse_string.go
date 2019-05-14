package reverse


import "strings"


func String(str string) string {
	length := len(str)
	if length == 0 {
		return str
	}

	frags := make([]string, length)
	for i, r := range str {
		frags[length - i - 1] = string(r)
	}
	return strings.Join(frags, "")
}
