package brackets


func nested(l, r rune) bool {
	if l == '{' && r == '}' {
		return true
	} else if  l == '[' && r == ']' {
		return true
	} else if l == '(' && r == ')' {
		return true
	}
	return false
}


func isLeft(b rune) bool {
	return (b == '(' || b == '[' || b == '{')
}


func isRight(b rune) bool {
	return (b == ')' || b == ']' || b == '}')
}


func Bracket(brackets string) bool {
	stack := make([]rune, 0)
	for _, b := range brackets {
		if !(isLeft(b) || isRight(b)) {
			continue
		}

		length := len(stack)
		if isLeft(b) {
			stack = append(stack, b)
		} else {
			if length == 0 {
				return false
			}
			if !nested(stack[length - 1], b) {
				return false
			}
			stack = stack[:length - 1]
		}
	}
	return len(stack) == 0
}
