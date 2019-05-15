package strain

type Ints []int
type Lists [][]int
type Strings []string

func (ints Ints) Keep(fn func(int) bool) Ints {
	if ints == nil {
		return ints
	}
	result := make(Ints, 0)
	for _, i := range ints {
		if fn(i) {
			result = append(result, i)
		}
	}
	return result
}

func (ints Ints) Discard(fn func(int) bool) Ints {
	if ints == nil {
		return ints
	}
	result := make(Ints, 0)
	for _, i := range ints {
		if !fn(i) {
			result = append(result, i)
		}
	}
	return result
}

func (lists Lists) Keep(fn func([]int) bool) Lists {
	if lists == nil {
		return lists
	}
	result := make(Lists, 0)
	for _, i := range lists {
		if fn(i) {
			result = append(result, i)
		}
	}
	return result
}

func (strings Strings) Keep(fn func(string) bool) Strings {
	if strings == nil {
		return strings
	}
	result := make(Strings, 0)
	for _, i := range strings {
		if fn(i) {
			result = append(result, i)
		}
	}
	return result
}
