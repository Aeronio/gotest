package main

import "fmt"

//Queue represents a queue that holds a slice
type Queue struct {
	items []int
}

//Enqueue adds the items to the queue
func (q *Queue) Enqueue(i int) {
	//Adds the items to the queue
	q.items = append(q.items, i)
}

//Dequeue returns the first item and deletes it from the slice of items
func (q *Queue) Dequeue() int {
	if q.Check() {
		err := fmt.Errorf("there are no more items in this Queue")
		fmt.Println(err.Error())
		return 0
	}
	dequeued := q.items[0]
	q.items = q.items[1:]
	return dequeued
}

//Checks to see if the Queue is empty
func (q *Queue) Check() bool {
	return len(q.items) == 0
}

func main() {

	testQueue := Queue{}

	fmt.Println(testQueue)
	fmt.Println(testQueue.Check())

	for i := 0; i < 500; i += 100 {
		testQueue.Enqueue(i)
	}
	fmt.Println(testQueue.Check())

	fmt.Println(testQueue)
	fmt.Println(testQueue.Dequeue())
	fmt.Println(testQueue)
	fmt.Println(testQueue.Dequeue())
	fmt.Println(testQueue)
	fmt.Println(testQueue.Dequeue())
	fmt.Println(testQueue)
	fmt.Println(testQueue.Dequeue())
	fmt.Println(testQueue)
	fmt.Println(testQueue.Check())
	fmt.Println(testQueue.Dequeue())
	fmt.Println(testQueue)
	fmt.Println(testQueue.Check())
	fmt.Printf("Value of test queue: %v \n", testQueue)

	fmt.Println(testQueue.Dequeue())
	fmt.Println(testQueue)

	testQueue.Enqueue(200)
	fmt.Println(testQueue)
	fmt.Println(testQueue.Check())

	// testQueue.Enqueue(200)
	// testQueue.Enqueue(300)

	// fmt.Println(testQueue.Dequeue())
	// fmt.Println(testQueue)
	// fmt.Println(testQueue.Dequeue())
	// fmt.Println(testQueue)
	// fmt.Println(testQueue.Dequeue())
	// fmt.Println(testQueue)

}
