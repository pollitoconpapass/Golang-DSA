package main

import "fmt"

// === NODE DECLARATION ===
type node struct {
	data int
	next *node
}

// === ACTUAL LINKED LIST ===
type linkedList struct {
	head *node // head will contain the values of the node struct...
	// tail *node  // only on double linked lists
	length int
}

// === ADD A NEW NODE ===
func (l *linkedList) prepend(n *node) {
	second := l.head     // second is the head of the list
	l.head = n           // make the new node the head of the list
	l.head.next = second // make the second node the next node of the new node
	l.length++           // increase the length of the list by one
}

// === PRINT OUT THE LIST ===
func (l linkedList) printListData() {
	toPrint := l.head // -> declare the first node to print
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data) // -> print out the data
		toPrint = toPrint.next          // -> move to the next node
		l.length--                      // -> decrease the length of the list by one (to be able to exit the loop)
	}
	fmt.Printf("\n")
}

// === DELETE A NODE WITH A GIVEN VALUE ===
func (l *linkedList) deleteWithValue(value int) {
	// CASE 1: EMPTY LIST
	if l.length == 0 {
		return
	}

	// CASE 2: HEAD NODE
	if l.head.data == value { // -> if the head of the list is the node we want to delete
		l.head = l.head.next // -> make the next node the head of the list
		l.length--
		return
	}

	previousToDelete := l.head                // -> declare the first node to delete
	for previousToDelete.next.data != value { // -> loop until the value of the next node is the value we want to delete

		// CASE 3: NODE NOT IN LIST
		if previousToDelete.next.next == nil { // -> if the next node is nil, the node we want to delete is not in the list
			return
		}
		previousToDelete = previousToDelete.next // -> move to the next node
	}
	// CASE 4: NODE IN LIST
	previousToDelete.next = previousToDelete.next.next // -> make the next node of the previous node the next node of the next node
	// ... like skipping the node we want to delete
	l.length--
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 36}
	node3 := &node{data: 24}
	node4 := &node{data: 12}
	node5 := &node{data: 0}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)

	myList.printListData() // -> output: 0 12 24 36 48

	myList.deleteWithValue(100)

	myList.deleteWithValue(24)
	myList.printListData() // -> output: 0 12 36 48

	myList.deleteWithValue(0)
	myList.printListData() // -> output: 12 36 48
}
