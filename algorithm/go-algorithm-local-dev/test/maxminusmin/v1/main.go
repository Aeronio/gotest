//This finds the difference between the highest value and lowest value within a given slice of int
package main

import "fmt"

//MaxV takes in a slice of int and RETURNs the maximum value in the slice
func MaxV(xi []int) int {

	maxValue := 0
	//sliceLen := len(xi) - 1
	for _, v := range xi {
		if maxValue < v {
			maxValue = v
		}
	}
	return maxValue
}

//MinV takes in a slice of int and current maximum, then RETURNs the minimum value in the slice
func MinV(xi []int, max int) int {

	minValue := max
	//sliceLen := len(xi) - 1
	for _, v := range xi {
		if minValue > v {
			minValue = v
		}
	}
	return minValue
}

//Calculates the difference in range in a slice of int and RETURNs it as an int
func maxMinusMin(xi []int) int {

	maxVal := MaxV(xi)
	fmt.Println("This is the maximum value: ", maxVal)

	minVal := MinV(xi, maxVal)
	fmt.Println("This is the minimum value: ", minVal)

	return maxVal - minVal
}

func main() {

	testSlice := []int{20, 10, 6, 9, 15, 44, 7, 2}

	fmt.Println(testSlice)

	fmt.Println(maxMinusMin(testSlice))

	testSlice2 := []int{199, 864, 45, 1, 143, 1000, 42, 24}
	fmt.Println(maxMinusMin(testSlice2))

}
