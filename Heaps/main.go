package main

import (
	"fmt"
)

type MaxHeap struct {
	array []int
}

// === ADDS VALUES TO THE HEAP ===
func (h *MaxHeap) Insert(value int) {
	h.array = append(h.array, value)
	h.maxHeapifyUp(len(h.array) - 1)
}

// === WILL EXTRACT THE LARGEST NUMBER IN THE HEAP (ROOT) ===
func (h *MaxHeap) Extract() int {
	extract := h.array[0]    // -> first element in an array is the root
	left := len(h.array) - 1 // -> last element in an array is the last index

	if len(h.array) == 0 { // in case it's empty
		fmt.Print("Cannot extract from an empty heap")
		return -1
	}

	// takes the last index and places it in the root
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:left]

	h.maxHeapifyDown(0)
	return extract
}

// === WILL HEAPIFY FROM BOTTOM TO TOP ===
// Heapify -> process in which the binary tree is reshaped
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] { // -> in case the parent is less than the child
		h.swap(parent(index), index)
		index = parent(index) // -> to continue the loop
	}
}

// === WILL HEAPIFY FROM TOP TO BOTTOM ===
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)

	childToCompare := 0

	// loop while index has one child
	for l <= lastIndex {
		if l == lastIndex { // -> when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // -> when left child is greater than right child
			childToCompare = l
		} else { // -> when right child is greater than left child
			childToCompare = r
		}

		// COMPARE THE CHILD WITH THE PARENT ANS SWAP (IF NECESSARY)
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// === GET THE PARENT ===
func parent(i int) int {
	return (i - 1) / 2
}

// === LEFT CHILD ===
func left(i int) int {
	return 2*i + 1
}

// === RIGHT CHILD ===
func right(i int) int {
	return 2*i + 2
}

// === SWAP THE NODES ===
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1] // -> Go method to swap the values
}

// === MAIN PROCESS ===
func main() {
	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 5, 7, 3, 1, 9, 2}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(m.Extract())
		fmt.Println(m)
	}
}
