package main

import (
	"fmt"

	tree "github.com/ManhHungAF/Tutorial/go/tree-building"
)

var (
	input = []tree.Record{
		{ID: 5, Parent: 1},
		{ID: 3, Parent: 2},
		{ID: 2, Parent: 0},
		{ID: 4, Parent: 1},
		{ID: 1, Parent: 0},
		{ID: 0},
		{ID: 6, Parent: 2}}
)

func main() {
	a, _ := tree.Build(input)
	fmt.Println(*a)
}
