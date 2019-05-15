package binarysearchtree

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

func Bst(num int) *SearchTreeData {
	return &SearchTreeData{data: num}
}

func (tree *SearchTreeData) Insert(num int) {
	node := tree
	for {
		if num <= node.data {
			if node.left == nil {
				node.left = Bst(num)
				break
			} else {
				node = node.left
			}
		} else {
			if node.right == nil {
				node.right = Bst(num)
				break
			} else {
				node = node.right
			}
		}
	}
}

func (tree *SearchTreeData) MapString(fn func(int) string) []string {
	result := make([]string, 0)

	if tree.left != nil {
		result = append(result, tree.left.MapString(fn)...)
	}
	result = append(result, fn(tree.data))
	if tree.right != nil {
		result = append(result, tree.right.MapString(fn)...)
	}
	return result
}

func (tree *SearchTreeData) MapInt(fn func(int) int) []int {
	result := make([]int, 0)

	if tree.left != nil {
		result = append(result, tree.left.MapInt(fn)...)
	}
	result = append(result, fn(tree.data))
	if tree.right != nil {
		result = append(result, tree.right.MapInt(fn)...)
	}
	return result
}
