package main

import "fmt"

type Stack struct {
	items []int
}

// === PUSH (ADD VALUE) ===
func (s *Stack) Push(value int) {
	s.items = append(s.items, value) // -> append(list, value to append)
}

// === POP (REMOVE VALUE) ===
func (s *Stack) Pop() int { // -> don't need the value to remove, because the value removed is gonna be the last entered
	length := len(s.items) - 1      // -> get the length of the list - 1 (because the last entered value is the first to be removed)
	removedValue := s.items[length] // -> it's the index of the last entered value
	s.items = s.items[:length]      // -> keep the values of the list until reaching the last entered value

	return removedValue
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)

	myStack.Push(1)
	myStack.Push(100)
	myStack.Push(2)
	myStack.Push(200)
	myStack.Push(3)
	myStack.Push(300)

	fmt.Println(myStack) // -> output: [1 100 2 200 3 300]

	myStack.Pop()
	fmt.Println(myStack) // -> output: [1 100 2 200 3]  // -> the last entered value is the first to be removed
}
