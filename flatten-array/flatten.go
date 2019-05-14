package flatten


func Flatten(lists interface{}) []interface{} {
	list := make([]interface{}, 0)
	intLists, _ := lists.([]interface{})
	for _, element := range intLists {
		switch element.(type) {
		case int:
			list = append(list, element)
		default:
			elist, _ := element.([]interface{})
			list = append(list, Flatten(elist)...)
		}
	}
	return list
}
