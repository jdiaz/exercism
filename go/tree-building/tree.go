package tree

import (
	"fmt"
	"sort"
)

// Record represents an entry in the database
type Record struct {
	ID     int
	Parent int
}

// Node represents a item to display
type Node struct {
	ID       int
	Children []*Node
}

// Build constructs a tree structure out of a slice of records
func Build(records []Record) (*Node, error) {
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	nodes := make(map[int]*Node)
	for i := range records {
		rec := records[i]
		if rec.ID != i {
			return nil, fmt.Errorf("unexpected record id: %d, expected: %d", rec.ID, i)
		}
		if i == 0 && rec.Parent != 0 {
			return nil, fmt.Errorf("root node should not have a parent (%d)", rec.Parent)
		}
		if i != 0 && rec.Parent >= i {
			return nil, fmt.Errorf("parent id (%d) should be lower than its own id (%d)", rec.Parent, i)
		}
		nodes[i] = &Node{ID: i}
		if i != 0 {
			parent := nodes[rec.Parent]  
			parent.Children = append(parent.Children, nodes[i])
		}
	}
	return nodes[0], nil
}
