//v2 Converts the funcs from accepting a slice of int as a pass-in variable to using a custom type and associating methods with it
package main

import "fmt"

type RangeSlice []int

//MaxV takes in a slice of int and RETURNs the maximum value in the slice
func (xi RangeSlice) MaxV() int {

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
func (xi RangeSlice) MinV(max int) int {

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
func (xi RangeSlice) maxMinusMin() int {

	maxVal := xi.MaxV()
	fmt.Println("This is the maximum value: ", maxVal)

	minVal := xi.MinV(maxVal)
	fmt.Println("This is the minimum value: ", minVal)

	return maxVal - minVal
}

func main() {

	testSlice := RangeSlice{20, 10, 6, 9, 15, 44, 7, 2}

	fmt.Println(testSlice)

	fmt.Println(testSlice.maxMinusMin())

	testSlice2 := RangeSlice{199, 864, 45, 1, 143, 1000, 42, 24}
	fmt.Println(testSlice2.maxMinusMin())

}
