//Divide and Conquer is when you divide a problem down into smaller problems recursively, then return back upwards into a single answer after reaching a base case.
// 3 steps: Divide, Conquer, Combine
//We divide a problem so small that it cannot recurse anymore, the smallest unit is called a "Base case"
//The recursion "Bottoms out" and then starts returning the results
//max subarray problem (Return the maximum "profit" you could get from the difference in values in an array),
//Bottoms out on the individual values in an array, returns those values to the next problem up (where some problem solving can be done),
//Continue on until we get to the top and have a final total solution
//In a weird way, it's like PEMDAS haha with a lot of nested Parentheses
//The subarray that catches the greatest sum will fall into 3 cases: entirely in the left, entirely in the right, or crossing the middle

package main

import "fmt"

//maxSubArray is here so the user only needs to input the array, then the func does the rest of the work by calling findMaxSubArray and recursing
func maxSubArray(nums []int) int {
	right := len(nums) - 1
	return findMaxSubArray(nums, 0, right)
}

//findMaxSubArray is function used to divide and conquer, it returns the maximum "profit" int of the array passed in
func findMaxSubArray(nums []int, left int, right int) int {

	//This "if" checks our base case, we are at a single number it will start returning upward
	if left == right {
		return nums[left]
	}
	mid := (left + right) / 2
	leftMax := findMaxSubArray(nums, left, mid)
	rightMax := findMaxSubArray(nums, mid+1, right)
	crossMax := maxCrossing(nums, left, right, mid)

	return max(leftMax, rightMax, crossMax)
}

//This is a helper function specifically for the maximum sub array problem, not part of the divide and conquer method
func maxCrossing(nums []int, left int, right int, mid int) int {
	//This is the lowest number in a int32 bit architecture, we do this because the leftSum always needs to lose the first comparison
	//leftSum := -2147483648
	//Ideally, in a real world problem, you want to set this number to the lowest maximum allowed by your program
	//It can also be expressed like this
	leftSum := -1 << 31
	rightSum := -1 << 31

	sum := 0

	for i := mid; i >= left; i-- {
		sum += nums[i]
		leftSum = max(leftSum, sum)
	}

	sum = 0

	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		rightSum = max(rightSum, sum)
	}

	return leftSum + rightSum
}

//Review: Variadic function! "..." means you can plug in as many inputs of int as you want, and it assigns it to a []int named values
//max returns a single integer that will be largest of the passed in inputs
func max(values ...int) int {
	maxVal := values[0]
	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func main() {
	problem := []int{-2, 1, -3, 4, -1, 2, 1}
	answer := maxSubArray(problem)
	fmt.Println(answer)
}
