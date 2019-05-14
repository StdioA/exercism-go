package listops


type IntList []int
type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(n int) int


func (list IntList) Foldl(fn binFunc, initial int) int {
	result := initial
	for _, num := range list {
		result = fn(result, num)
	}
	return result
}

func (list IntList) Foldr(fn binFunc, initial int) int {
	result := initial
	length := list.Length()
	for i:=length-1; i>=0; i-- {
		result = fn(list[i], result)
	}
	return result
}


func (list IntList) Length() int {
	return len(list)
}


func (list IntList) Filter(fn predFunc) IntList {
	result := make(IntList, 0)
	for _, num := range list {
		if fn(num) {
			result = append(result, num)
		}
	}
	return result
}


func (list IntList) Map(fn unaryFunc) IntList {
	result := make(IntList, list.Length())
	for i, num := range list {
		result[i] = fn(num)
	}
	return result
}


func (list IntList) Reverse() IntList {
	length := list.Length()
	result := make(IntList, length)
	for i := 0; i < length; i++ {
		result[length-i-1] = list[i]
	}
	return result
}


func (list IntList) Append(another IntList) IntList {
	return append(list, another...)
}


func (list IntList) Concat(lists []IntList) IntList {
	for _, l := range lists {
		list = append(list, l...)
	}
	return list
}
