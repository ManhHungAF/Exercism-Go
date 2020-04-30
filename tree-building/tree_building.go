package tree

import (
	"errors"
	"sort"
)

// Record holds an individual record with an ID and ID of its parent
type Record struct {
	ID     int
	Parent int
}

// Node holds a node and a slice of pointers to its child node
type Node struct {
	ID       int
	Children []*Node
}

// Build builds a tree from a slice of Record. It returns an error if the
// input records is invalid
func Build(records []Record) (head *Node, err error) {

	if len(records) == 0 {
		return nil, nil
	}
	// sort input records base on the order of ID
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make([]Node, len(records))
	for i, r := range records {
		if r.ID != i || r.ID < r.Parent || (r.ID != 0 && r.ID == r.Parent) {
			return nil, errors.New("non-consecutive record or invalid parent")
		}
		
		if i > 0 {
			nodes[r.ID].ID = r.ID
			nodes[r.Parent].Children = append(nodes[r.Parent].Children, &nodes[r.ID])
		}
	}
	return &nodes[0], nil
}
