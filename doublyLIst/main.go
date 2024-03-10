package main

import (
	"fmt"
)

type Node struct {
	value any
	next  *Node
	prev  *Node
}

type List struct {
	head *Node
	tail *Node
}

func main() {
	li := &List{}
	li.append(5)
	li.append(7)
	li.append(8)
	li.printList()
	li.preappend(9)
	li.preappend(2)
	li.printList()
	li.insert(3, 8)
	li.insert(7, 66)
	li.printList()
	li.delete(0)
	li.delete(3)
	li.delete(3)

	li.printList()

}
func (li *List) printList() {
	currentNode := li.head
	for currentNode != nil {
		fmt.Print(currentNode.value, "----->")
		currentNode = currentNode.next
	}
	fmt.Println()
}
func (li *List) append(data any) {
	newNode := &Node{
		value: data,
		next:  nil,
		prev:  nil,
	}

	if li.head == nil {
		li.head = newNode
		li.tail = newNode
		return
	}
	li.tail.next = newNode
	newNode.prev = li.tail
	li.tail = newNode

}

func (li *List) preappend(data any) {
	newNode := &Node{
		value: data,
		next:  nil,
		prev:  nil,
	}

	if li.head == nil {
		li.head = newNode
		li.tail = newNode
		return
	}

	li.head.prev = newNode
	newNode.next = li.head
	li.head = newNode
}

func (li *List) insert(index int, data any) {

	newNode := &Node{
		value: data,
		next:  nil,
		prev:  nil,
	}
	if li.head == nil || index == 0 {
		li.head = newNode
		li.tail = newNode
		return
	}
	currentNode := li.head

	for i := 0; i < index-1 && currentNode != nil; i++ {
		currentNode = currentNode.next
	}

	if currentNode == nil {
		li.append(data)
		return

	}

	newNode.next = currentNode.next
	newNode.prev = currentNode
	if currentNode.next != nil {

		currentNode.next.prev = newNode
	}

}

func (li *List) delete(index int) {
	if index < 0 {
		fmt.Println("invalid index..........")
		return
	}

	if index == 0 && li.head != nil {
		if li.head.next != nil {
			li.head = li.head.next
			li.head.prev = nil
		} else {
			li.head = nil
		}
		return
	}
	currentNode := li.head
	for i := 0; i < index && currentNode != nil; i++ {
		currentNode = currentNode.next
	}

	if currentNode == nil {
		fmt.Println("index out of range")
		return
	}
	if currentNode.next == nil {
		currentNode.prev.next = nil

	}

	if currentNode.next != nil {
		currentNode.next.prev = currentNode.prev
		currentNode.prev.next = currentNode.next

	}

}
