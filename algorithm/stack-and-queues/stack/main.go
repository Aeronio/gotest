package main

import "fmt"

//Stack represents stack that hold a slice
type Stack struct {
	items []int
}

//Push will add a value at the end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

//Pop will remove a value at the end and RETURNs the removed value
func (s *Stack) Pop() int {
	if s.Check() {
		err := fmt.Errorf("there are no more items in this Stack")
		fmt.Println(err.Error())
		return 0
	}
	l := len(s.items) - 1
	popped := s.items[l]
	s.items = s.items[:l]
	return popped
}

//Checks to see if the Stack is empty
func (s *Stack) Check() bool {
	return len(s.items) == 0
}

func main() {
	testStack := Stack{}

	fmt.Println(testStack)

	for i := 0; i < 500; i += 100 {
		testStack.Push(i)
	}

	fmt.Println(testStack)
	fmt.Println(testStack.Pop())
	fmt.Println(testStack)
	fmt.Println(testStack.Pop())
	fmt.Println(testStack)
	fmt.Println(testStack.Pop())
	fmt.Println(testStack)
	fmt.Println(testStack.Pop())
	fmt.Println(testStack)
	fmt.Println(testStack.Pop())
	fmt.Println(testStack)
	fmt.Println(testStack.Check())

	//Run out of the stack
	fmt.Println(testStack.Pop())
	fmt.Println(testStack)
	fmt.Println(testStack.Pop())
	fmt.Println(testStack)

	testStack.Push(101)
	fmt.Println(testStack)
	fmt.Println(testStack.Check())
	fmt.Println(testStack.Pop())
}
