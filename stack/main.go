package main

import "fmt"

type Stack struct {
	stack []any
}

func main() {

	fmt.Println("stack in go........")
	st := &Stack{}
	st.push(1)
	st.push(2)
	st.push(3)
	st.peek()
	st.pop()
	st.peek()
}

func (st *Stack) push(value any) {

	st.stack = append(st.stack, value)
}

func (st *Stack) pop() {

	if len(st.stack) == 0 {
		fmt.Println("stack is empty.")
		return
	}

	item := st.stack[len(st.stack)-1]

	st.stack = st.stack[:len(st.stack)-1]

	fmt.Println("item is removed", item)

}

func (st *Stack) peek() {

	if len(st.stack) == 0 {
		fmt.Println("stack is empty")
		return
	}

	item := st.stack[len(st.stack)-1]
	fmt.Println("peek element is ", item)
}
