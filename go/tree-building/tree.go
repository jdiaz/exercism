package tree

import (
	"errors"
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
	n := len(records)
	if n == 0 {
		return nil, nil
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	nodeMap := make(map[int]*Node)
	rootSeen := false
	for i, rec := range records {
		childID := rec.ID
		parentID := rec.Parent
		if childID == 0 && parentID > 0 {
			return nil, errors.New("Invalid structure, root can not have parent")
		}
		if parentID > childID {
			return nil, errors.New("Record cannot have higher or equal id parent of lower id")
		}
		if childID != 0 && parentID != 0 && childID == parentID {
			return nil, errors.New("Records cycle directly")
		}
		if _, prevExists := nodeMap[childID-1]; !prevExists && i != 0 {
			return nil, errors.New("Records are non-continuous")
		}

		// Child logic
		var childNodePtr *Node
		_, childExists := nodeMap[childID]
		if childExists && rootSeen {
			//childNodePtr = child
			//fmt.Printf("who am  i? => %+v", nodeMap)
			return nil, errors.New("Duplicate record")
		} else if !childExists {
			childNodePtr = &Node{ID: childID}
			nodeMap[childID] = childNodePtr
		}
		// Parent logic
		if childID != 0 {
			var parentNodePtr *Node
			parent, parentExists := nodeMap[parentID]
			if parentExists {
				parentNodePtr = parent
			} else {
				parentNodePtr = &Node{ID: parentID, Children: []*Node{}}
				nodeMap[parentID] = parentNodePtr
			}
			parentNodePtr.Children = append(parentNodePtr.Children, childNodePtr)
		}

		if childID == 0 {
			rootSeen = true
		}
	}

	if !rootSeen {
		return nil, errors.New("Root not found")
	}
	root, rootExists := nodeMap[0]
	if !rootExists {
		return nil, errors.New("Root node does not exists")
	}
	return root, nil
}
