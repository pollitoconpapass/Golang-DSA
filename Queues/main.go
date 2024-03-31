package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) Dequeue() int {
	removedValue := q.items[0] // -> the first item is the first to be removed
	q.items = q.items[1:]      // -> skip the first item (index 0) and keep all the others (from index 1 to the end)

	return removedValue
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	myQueue.Enqueue(400)
	fmt.Println(myQueue) // -> output: [200 300 400]

	myQueue.Dequeue()
	fmt.Println(myQueue) // -> output: [300 400]
}
