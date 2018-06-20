package tree

import (
	"fmt"
)

// Tree struct describes binary tree. The value of the tree could be anything.
type Tree struct {
	Left  *Tree
	Value interface{}
	Right *Tree
}

// New creates tree with root equal to the passed value.
func New(value interface{}) *Tree {
	return &Tree{nil, value, nil}
}

// Print the tree
func (t *Tree) Print() {
	if t.Left != nil {
		t.Left.Print()
	}

	if t.Right != nil {
		t.Right.Print()
	}

	fmt.Println(t.Value)
}
