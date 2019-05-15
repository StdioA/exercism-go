package sieve

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}
	array := make([]int, limit+1)
	for i := 2; i < len(array); i++ {
		array[i] = i
	}

	for i := 2; i < len(array); i++ {
		if array[i] == 0 {
			continue
		}
		k := i * 2
		for k < len(array) {
			array[k] = 0
			k += i
		}
	}

	result := make([]int, 0)
	for _, num := range array {
		if num > 0 {
			result = append(result, num)
		}
	}
	return result
}
