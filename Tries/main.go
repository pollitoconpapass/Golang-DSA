package main

import "fmt"

const AlphabetSize int = 26 // -> number of possible characters in the trie

// === NODE DECLARATION ===
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// === TRIE DECLARATION ===
type Trie struct {
	root *Node // -> the root itself is a Node as well
}

// === INIT TRIE FUNCTION (WILL CREATE A NEW TRIE) ===
func InitTrie() *Trie {
	result := &Trie{root: &Node{}} // -> will create a Node which will be declared as the root of the whole trie
	return result
}

// === INSERT METHOD ===
func (t *Trie) Insert(word string) {
	currentNode := t.root
	for i := 0; i < len(word); i++ {
		charIndex := word[i] - 'a' // -> convert the character to an index
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{} // -> created a new Node if it's an empty position
		}

		currentNode = currentNode.children[charIndex] // -> move to the next Node
	}
	currentNode.isEnd = true // -> mark the last Node as the end of the word
}

// === SEARCH METHOD ===
func (t *Trie) Search(word string) bool {
	currentNode := t.root
	for i := 0; i < len(word); i++ {
		charIndex := word[i] - 'a' // -> convert the character to an index
		if currentNode.children[charIndex] == nil {
			return false // -> if the Node is empty, return false ... is not the same as the others, because this time we implemented a isEnd variable
		}
		currentNode = currentNode.children[charIndex] // -> move to the next Node
	}
	if currentNode.isEnd == true {
		return true // -> if the Node is the end of the word, return true
	}

	return false // -> if the Node is not the end of the word, return false
}

func main() {
	myTrie := InitTrie()

	myTrie.Insert("hello")
	myTrie.Insert("hell")
	result := myTrie.Search("hell")
	fmt.Println(result)

}
