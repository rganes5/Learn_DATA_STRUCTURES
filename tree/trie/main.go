package main

import (
	"fmt"
	"strings"
)

// main function
func main() {
	var word string
	var searchWord string
	trie := newTrie()
	for i := 0; i < 3; i++ {
		fmt.Println("Enter your word")
		fmt.Scan(&word)
		trie.insert(word)
	}
	fmt.Println("Enter the word for search")
	fmt.Scan(&searchWord)
	found := trie.Search(searchWord)
	if found == 1 {
		fmt.Printf("\"%s\" Word found in trie\n", searchWord)
	} else {
		fmt.Printf(" \"%s\" Word not found in trie\n", searchWord)
	}

}

type Node struct {
	char  string
	child [26]*Node
	end   bool
}

type Trie struct {
	rootNode *Node
}

func newNode(char string) *Node {
	node := &Node{char: char}
	for i := 0; i < 26; i++ {
		node.child[i] = nil
	}
	return node
}

func newTrie() *Trie {
	root := newNode(" ")
	return &Trie{rootNode: root}
}

// insert function
func (t *Trie) insert(word string) {
	current := t.rootNode
	stripWord := strings.ToLower(word)
	for _, v := range stripWord {
		index := v - 'a'
		if current.child[index] == nil {
			current.child[index] = newNode(string(v))
		}
		current = current.child[index]
	}
	current.end = true
	return
}

// search function
func (t *Trie) Search(word string) int {
	current := t.rootNode
	for _, v := range word {
		index := v - 'a'
		if current.child[index] == nil {
			return 0
		}
		current = current.child[index]
	}
	if current.end {
		return 1
	}
	return 0
}
