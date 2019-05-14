package tree


import (
	"fmt"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID int
	Children []*Node
}

type NodeSlice []*Node


// Implement sort.Interface
func (n NodeSlice) Len() int {
	return len(n)
}

func (n NodeSlice) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n NodeSlice) Less(i, j int) bool {
	return n[i].ID < n[j].ID
}


func Build(records []Record) (*Node, error) {
	length := len(records)
	if length == 0 {
		return nil, nil
	}

	nodes := make([]Node, length)
	parents := make([]*Node, len(records))
	seen := make([]bool, len(records))
	// Error check
	for _, record := range records {
		if record.ID > 0 && record.ID <= record.Parent {
			return nil, fmt.Errorf("ID smaller than parent")
		}
		if record.ID >= length {
			return nil, fmt.Errorf("ID too large")
		}
		if seen[record.ID] {
			return nil, fmt.Errorf("Record with id %d occurs multiple times", record.ID)
		}
		seen[record.ID] = true
		if record.ID != 0 {
			// Bucket sort for parent setting
			parents[record.ID] = &nodes[record.Parent]
		} else if record.Parent != 0 {
			return nil, fmt.Errorf("Root node has non-0 parent %d", record.Parent)
		}
	}
	
	for i := 1; i < len(nodes); i++ {
		parents[i].Children = append(parents[i].Children, &nodes[i])
	}

	for i, node := range nodes {
		nodes[i].ID = i
		sort.Sort(NodeSlice(node.Children))
	}
	return &nodes[0], nil
}
