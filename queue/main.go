package main

import "fmt"

type Queue struct {
	items []any
}

func main() {

	queue := &Queue{}
	queue.isEmpty()
	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)

	queue.frontElement()

	queue.deQueue()
	queue.frontElement()
	queue.isEmpty()

}

func (e *Queue) enqueue(data any) {

	e.items = append(e.items, data)
	fmt.Println("added to queue", data)

}

func (e *Queue) deQueue() {

	if len(e.items) == 0 {
		fmt.Println("queue is empty")
		return
	}
	item := e.items[0]
	e.items = e.items[1:]

	fmt.Println("dequeued", item)
}

func (e *Queue) frontElement() {
	if len(e.items) == 0 {
		fmt.Println("ite is empty")
		return
	}
	item := e.items[0]
	fmt.Println("front element is", item)
}

func (q *Queue) isEmpty() {
	if len(q.items) == 0 {
		fmt.Println("queue is empty")
		return
	}

	fmt.Println("qeue is not empty")
	return
}
