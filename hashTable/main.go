package main

import (
	"fmt"
)

const arraySize = 20

type Node struct {
	key   string
	value any
	next  *Node
}

type HashTable struct {
	array [arraySize]*Node
}

func main() {
	ht := &HashTable{}
	ht.insert("banana", 56)
	ht.insert("apple", 3)
	ht.insert("apple", 5)
	fmt.Println("insert successfully")
	key, isFound := ht.search("apple")
	if isFound {
		fmt.Println("found with key", key)
	}
	if !isFound {
		fmt.Println("not found")
	}
	ht.delete("banana")
	k, isFound := ht.search("banana")
	if isFound {
		fmt.Println("found with key", k)
	}
	if !isFound {
		fmt.Println("not found")
	}

}

func (ht *HashTable) insert(key string, value any) {

	index := HashKey(key)

	newNode := &Node{
		value: value,
		key:   key,
		next:  nil,
	}
	if ht.array[index] == nil {
		ht.array[index] = newNode
	} else {

		current := ht.array[index]

		for current.next != nil {
			current = current.next

		}
		current.next = newNode
	}
	return

}

func HashKey(key string) int {

	hash := 0
	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}

	return hash % arraySize

}

func (ht *HashTable) search(key string) (string, bool) {

	index := HashKey(key)
	curr := ht.array[index]
	for curr != nil {
		if curr.key == key {
			return curr.key, true
		}
		curr = curr.next
	}

	return key, false
}

func (ht *HashTable) delete(key string) {

	index := HashKey(key)
	current := ht.array[index]
	var prev *Node
	for current != nil {
		if current.key == key {
			if prev == nil {
				ht.array[index] = current.next
			}
			if prev != nil {
				prev.next = current.next
			}
			return
		}
		prev = current
		current = current.next
	}

}
