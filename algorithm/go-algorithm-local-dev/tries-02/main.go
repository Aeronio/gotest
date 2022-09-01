package main

import "fmt"

type TrieNode struct {
	children   map[rune]*TrieNode
	endOfWords bool
}

func getNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode), endOfWords: false}
}

func insert(root *TrieNode, key string) {
	temp := root

	for _, ch := range key {
		if child, ok := temp.children[ch]; ok {
			temp = child
		} else {
			newChild := &TrieNode{children: make(map[rune]*TrieNode)}
			temp.children[ch] = newChild
			temp = newChild
		}
	}
	temp.endOfWords = true
}

func search(root *TrieNode, key string) bool {
	temp := root
	for _, ch := range key {
		if child, ok := temp.children[ch]; ok {
			temp = child
			continue
		}
		return false
	}
	return temp.endOfWords
}

func main() {
	words := []string{"a", "and", "go", "golang", "mango"}

	root := getNode()

	for i := 0; i < len(words); i++ {
		insert(root, words[i])
	}

	fmt.Println("contains the word and", search(root, "and"))
	fmt.Println("contains the word golang", search(root, "golang"))
	fmt.Println("contains the word man", search(root, "man"))
}
