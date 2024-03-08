package main

import "fmt"

type Node struct {
	value any
	next  *Node
}

type LinkList struct {
	head *Node
}

func main() {

	list := &LinkList{}

	list.append(1)
	list.append(2)
	list.append(3)
	list.printLIst()
	list.preappend(6)
	list.preappend(7)
	list.printLIst()

	list.insert(6, 7)
	list.insert(2, 3)

	list.printLIst()

	// delet by index
	list.deleteitem(1)

	list.printLIst()

}

func (li *LinkList) printLIst() {
	currNode := li.head
	for currNode != nil {
		fmt.Print(currNode.value, "--> ")
		currNode = currNode.next
	}
	fmt.Println()
}

func (li *LinkList) append(value any) {

	newNode := &Node{value: value, next: nil}

	if li.head == nil {
		li.head = newNode
		return
	}

	startNode := li.head
	for startNode.next != nil {
		startNode = startNode.next
	}
	startNode.next = newNode
	return

}

func (li *LinkList) preappend(data any) {
	head := li.head

	newNode := &Node{
		value: data,
		next:  nil,
	}

	if head == nil {
		head = newNode
		return
	}
	li.head = newNode
	newNode.next = head

}

func (li *LinkList) insert(data any, index int) {

	newNode := &Node{value: data, next: nil}
	if index < 0 {
		fmt.Println("inavlid index")
		return
	}

	if index == 0 {
		li.preappend(data)
		return
	}
	prevNode := li.head

	for i := 0; i < index-1 && prevNode != nil; i++ {

		prevNode = prevNode.next
	}
	if prevNode == nil {
		li.append(data)
		return
	}

	newNode.next = prevNode.next
	prevNode.next = newNode

}

func (li *LinkList) deleteitem(index int) {

	if index == 0 && li.head != nil {
		newHead := li.head.next
		li.head = newHead
		return
	}
	if li.head == nil {
		fmt.Println("list is already empty")
		return
	}

	prevNode := li.head

	for i := 0; i < index-1 && prevNode != nil; i++ {

		prevNode = prevNode.next
	}
	if prevNode == nil {
		fmt.Println("index out of range")
		return
	}

	deleteNode := prevNode.next

	prevNode.next = deleteNode.next

}
