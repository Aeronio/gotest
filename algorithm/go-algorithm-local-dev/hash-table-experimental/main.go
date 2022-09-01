package main

import (
	"fmt"
)

const ArraySize int = 7

//HashTable structure

type HashTable struct {
	array [ArraySize]*bucket
}

//bucket structure - it's a linked list
type bucket struct {
	head *bucketNode
}

//bucketNode structure
type bucketNode struct {
	key  string //These will be the names like Eric, etc.
	next *bucketNode
}

//Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

//Search takes in a key and returns true if that key is stored in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

//Delete takes in a key and deletes it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

//insert will take in a key, create a node with the key and insert the node in the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}
}

//search will take in a key and return true if the bucket has that key
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

//delete will take a key and delete the node from the bucket
func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		//This is the where the search comes into play
		if previousNode.next.key == k {
			//delete
			if previousNode.next.next != nil {
				previousNode.next = previousNode.next.next
			} else {
				previousNode.next = nil
				return
			}
		}
		//This is how it updates to the next iteration in the for loop
		previousNode = previousNode.next
	}
}

//hash function
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

//Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {

	hashTable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"REDO",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	hashTable.Delete("STAN")
	fmt.Println("STAN:", hashTable.Search("STAN"))
	fmt.Println("KENNY:", hashTable.Search("KENNY"))
	fmt.Println("ERIC:", hashTable.Search("ERIC"))
	fmt.Println("REDO:", hashTable.Search("REDO"))

	hashTable.Delete("ERIC")
	fmt.Println("ERIC:", hashTable.Search("ERIC"))

	hashTable.Insert("ERIC")
	fmt.Println("ERIC:", hashTable.Search("ERIC"))

	hashTable.Delete("KENNY")
	fmt.Println("KENNY:", hashTable.Search("KENNY"))

	hashTable.Insert("KENNY")
	fmt.Println("KENNY:", hashTable.Search("KENNY"))

	hashTable.Delete("REDO")
	fmt.Println("REDO:", hashTable.Search("REDO"))

	hashTable.Insert("REDO")
	fmt.Println("REDO:", hashTable.Search("REDO"))

	// fmt.Println("Something")
	// testHash := Init()
	// fmt.Println(testHash)
	// fmt.Println(hash("RANDY"))

	// testBucket := &bucket{}
	// testBucket.insert("RANDY")
	// testBucket.insert("RANDY")
	// testBucket.delete("RANDY")
	// fmt.Println(testBucket.search("RANDY"))
	//fmt.Println(testBucket.search("JARED"))
}
