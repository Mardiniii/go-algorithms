package main

import (
	"fmt"
	// Import a previous binary tree implementation
	// Check: https://bit.ly/2uJi7mB
	tree "github.com/Mardiniii/binarytree"
)

// DFS iterates using depth first search algorithm over the tree
func DFS(t tree.BinaryTree, f func(v interface{})) {
	// Stack using slices
	var q []*tree.Node
	// Push the root node to the stack
	q = append([]*tree.Node{t.Root()}, q...)

	for len(q) != 0 {
		// Pop the node from the stack
		n := q[0]
		q = q[1:]

		// For this case we will print the value
		// on each node
		f(n.Value())

		if n.Right() != nil {
			// Push the right node to the stack
			q = append([]*tree.Node{n.Right()}, q...)
		}
		if n.Left() != nil {
			// Push the left node to the stack
			q = append([]*tree.Node{n.Left()}, q...)
		}
	}
}

func print(v interface{}) {
	fmt.Println("Value:", v)
}

func main() {
	var t tree.BinaryTree
	t.Insert(0)
	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(4)
	t.Insert(5)
	t.Insert(6)
	fmt.Println("Printing tree using level order iteration")
	t.Print()
	fmt.Println("Printing tree using DFS iteration")
	// If we print this tree using DFS We should see
	// 0 - 1 - 3 - 4 - 2 - 5 - 6
	DFS(t, print)
}
