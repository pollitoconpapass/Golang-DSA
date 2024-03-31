package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// === INSERT METHOD ===
func (n *Node) Insert(value int) {
	if n.Key < value {
		// GO TO THE RIGHT
		if n.Right == nil { // -> if the right node is empty...
			n.Right = &Node{Key: value} // -> create it and insert the value
		} else {
			n.Right.Insert(value) //-> try again with recursion
		}
	} else if n.Key > value {
		// GO TO THE LEFT
		if n.Left == nil { // -> if the left node is empty...
			n.Left = &Node{Key: value} // -> create it and insert the value
		} else {
			n.Left.Insert(value) // -> try again with recursion
		}
	}
}

// === SEARCH METHOD ===
func (n *Node) Search(value int) bool {
	if n == nil { // CASE 1: EMPTY NODE
		return false
	}

	// CASE 2: FOUND THE NODE
	if n.Key < value { // -> if the node (root) is smaller than the value go to the right
		// ( because it means that the value we are looking for is greater than the root)
		return n.Right.Search(value)
	} else if n.Key > value {
		return n.Left.Search(value) // -> goes back at it with recursion
	}

	// As it's not empty, not grater than the root and not smaller than the root, it means that the root is the value we are looking for
	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(50)
	tree.Insert(200)
	tree.Insert(150)
	fmt.Println(tree)

	founded := tree.Search(150)
	fmt.Println(founded)
}
