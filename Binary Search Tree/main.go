package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func main() {
	fmt.Println("Binary search tree")
	bst := &BST{}
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(1)
	bst.Insert(9)

	isFound := bst.Search(9)
	fmt.Println("found : ", isFound)

}

func (b *BST) Insert(value int) {
	newRoot := &Node{
		data:  value,
		left:  nil,
		right: nil,
	}

	if b.root == nil {
		b.root = newRoot
		return
	} else {
		b.root.insert(value)
	}

}

func (b *Node) insert(value int) {
	newNode := &Node{
		data:  value,
		left:  nil,
		right: nil,
	}
	if value <= b.data {
		if b.left == nil {
			b.left = newNode
			return
		} else {
			b.left.insert(value)
		}
	} else {

		if b.right == nil {
			b.right = newNode
			return
		} else {
			b.right.insert(value)
			return
		}

	}
}

func (b *BST) Search(value int) bool {
	return b.root.search(value)
}

func (n *Node) search(value int) bool {
	if n == nil {
		return false
	}
	if value == n.data {
		return true
	}
	if value < n.data {
		return n.left.search(value)
	} else {
		return n.right.search(value)
	}

	return false
}



