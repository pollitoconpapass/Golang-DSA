package main

import "fmt"

const ArraySize int = 7

// === HASH TABLE STRUCTURE ===
type HashTable struct {
	array [ArraySize]*Bucket
}

// Array is a Bucket type
// Bucket is a Bucket Node type and contains methods to insert, search and delete
// BucketNode type contains a key and it's next value

// === INSERT ===
func (h *HashTable) Insert(key string) {
	index := Hash(key)
	h.array[index].InsertBucket(key) // -> like saying: "In this array, at this index, INSERT this key"
}

// === SEARCH ===
func (h *HashTable) Search(key string) bool {
	index := Hash(key)
	return h.array[index].SearchBucket(key) // -> like saying: "In this array, at this index, SEARCH for this key (return true or false)"
}

// === DELETE ===
func (h *HashTable) Delete(key string) {
	index := Hash(key)
	h.array[index].DeleteBucket(key) // -> like saying: "In this array, at this index, DELETE this key"
}

// === BUCKET STRUCTURE ===
type Bucket struct {
	head *BucketNode // -> implement the head of the bucket (linked list) which will have the key and it's next value
	// ... thanks to the BUCKET NODE IMPLEMENTATION
}

// === BUCKET NODE IMPLEMENTATION ===
type BucketNode struct { // -> linked list that holds the key (in case of collisions)
	Key  string
	Next *BucketNode
}

// === INSERT (BUCKET) ===
func (b *Bucket) InsertBucket(key string) {
	if !b.SearchBucket(key) { // -> if the key is not in the bucket (when first searching for it) then...
		newNode := &BucketNode{Key: key} // -> creates a new BucketNode
		newNode.Next = b.head            // -> makes the head of the bucket, the next node now...
		b.head = newNode                 // -> finally makes the new node the head of the bucket
		// Like adding a new value to a linked list...
	} else {
		fmt.Println(key, "already exists")
	}

}

// === SEARCH (BUCKET)===
func (b *Bucket) SearchBucket(key string) bool {
	currentNode := b.head    // -> start at the head of the bucket
	for currentNode != nil { // -> while it's not the end of the array...
		if currentNode.Key == key { // -> if the key is equal to the key we are looking for...
			return true // -> we've found it!!
		}
		currentNode = currentNode.Next // -> continues to the next node (iterating towards the array)
	}
	return false
}

// === DELETE (BUCKET) ===
func (b *Bucket) DeleteBucket(key string) {
	if b.head.Key == key {
		b.head = b.head.Next
		return
	}

	previousNode := b.head
	for previousNode.Next != nil {
		if previousNode.Next.Key == key { // -> if the key is equal to the key we are looking for...
			previousNode.Next = previousNode.Next.Next // -> we will skip that node
			return
		}
		previousNode = previousNode.Next // -> continues to the next node
	}
	fmt.Println(key, "does not exist")
}

// === HASH FUNCTION ===
func Hash(key string) int { // -> the actual hashing function to get the index of the array where the value is gonna be stored
	var sum int = 0
	for _, v := range key { // -> loop through the key (word)...
		sum += int(v) // -> ... adding the ASCII value of each character
	}

	return sum % ArraySize // -> the result of this is the index of the array
}

// === INIT FUNCTION ===
func Init() *HashTable {
	result := &HashTable{}

	for i := 0; i < ArraySize; i++ {
		result.array[i] = &Bucket{} // -> declaring addresses in each index of the array
	}
	return result
}

func main() {
	myHashTable := Init()
	myHashTable.Insert("Michael")
	myHashTable.Insert("Jennifer")
	myHashTable.Insert("John")
	myHashTable.Insert("Rose")
	myHashTable.Insert("Giovanna")

	myHashTable.Delete("Michael")
	fmt.Println("Michael", myHashTable.Search("Michael")) // -> trying to search for a value that doesn't exist anymore

	myHashTable.Insert("Rose") // -> trying to insert a value that already exists
}
