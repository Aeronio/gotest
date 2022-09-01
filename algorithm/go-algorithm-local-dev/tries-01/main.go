package main

import (
	"fmt"
)

//AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26

//Node represents each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

//Trie
type Trie struct {
	root *Node
}

//InitTrie will create a new Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

//Insert takes in a word and adds it to the trie
func (t *Trie) Insert(word string) {
	wordLength := len(word)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

//Search takes in a word and RETURNs true if in the trie
func (t *Trie) Search(word string) bool {
	wordLength := len(word)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

func main() {

	testTrie := InitTrie()
	fmt.Println(testTrie)
	fmt.Println(testTrie.root)

	testTrie.Insert("grass")
	fmt.Println("Grass in the trie: ", testTrie.Search("grass"))
	fmt.Println("Gras in the trie: ", testTrie.Search("gras"))

	toAdd := []string{
		"crass",
		"brass",
		"loss",
		"lost",
		"loser",
		"longer",
		"long",
		"loner",
		"loot",
		"lotto",
	}

	for _, v := range toAdd {
		testTrie.Insert(v)
	}

	searchWord := "looter"

	fmt.Printf("%s in the trie: %v", searchWord, testTrie.Search(searchWord))
}
