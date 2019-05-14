package accumulate

func Accumulate(strings []string, converter func(string) string) []string {
	for i, s := range strings {
		strings[i] = converter(s)
	}
	return strings
}
