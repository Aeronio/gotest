package main

import "fmt"

//Node represents the components of a binary search trees (each Node will possibly point to a left and right node)
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//Insert adds a node to the tree, the added key should not be already in the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		//Move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//Move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

//Search will take in a key value and RETURN true if there is a node with that value
func (n *Node) Search(k int) bool {

	if n == nil {
		return false
	}

	if n.Key < k {
		//move right
		return n.Right.Search(k)
	} else if n.Key > k {
		//move left
		return n.Left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(50)
	fmt.Println(tree)

	tree.Insert(150)
	fmt.Println(tree)

	tree.Insert(68)
	tree.Insert(567)
	tree.Insert(345)
	tree.Insert(310)
	tree.Insert(52)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(42)

	fmt.Println("Search Results for 7: ", tree.Search(7))
	fmt.Println("Search Results for 567: ", tree.Search(567))
	fmt.Println("Search Results for 1000: ", tree.Search(1000))
	fmt.Println("Search Results for 3: ", tree.Search(3))

	tree.Insert(3)
	fmt.Println("Search Results for 3: ", tree.Search(3))
	tree.Insert(1000)
	fmt.Println("Search Results for 1000: ", tree.Search(1000))
}
