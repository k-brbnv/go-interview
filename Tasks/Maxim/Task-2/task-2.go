/* Task 2. Дан отсортированный массив чисел.
Нужно вернуть отсортированный массив квадратов этих чисел.

Input: nums = [0, 1, 1, 6, 12, 20]
Output: squares = [0, 1, 1, 36, 144, 400]

[-2, -1, 0, 1, 2, 3, 4] = [0, 1, 1, 4, 4, 9, 16 ]

//Input: [[1,3], [2,6], [8,10], [15,18]]
//Output: [[1,6], [8,10], [15,18]]
*/

package main

import "fmt"

// Function to return sorted squares of an array without sorting
func sortedSquares(nums []int) []int {
	n := len(nums)
	squares := make([]int, n)
	left, right := 0, n-1
	pos := n - 1

	for left <= right {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		if leftSquare > rightSquare {
			squares[pos] = leftSquare
			left++
		} else {
			squares[pos] = rightSquare
			right--
		}
		pos--
	}

	return squares
}

func main() {
	// Example 1: Sorting squares of an array
	nums1 := []int{0, 1, 1, 6, 12, 20}
	fmt.Println("Input:", nums1)
	fmt.Println("Output:", sortedSquares(nums1)) // Expected output: [0, 1, 1, 36, 144, 400]

	// Example 2: Sorting squares of an array with negative numbers
	nums2 := []int{-2, -1, 0, 1, 2, 3, 4}
	fmt.Println("Input:", nums2)
	fmt.Println("Output:", sortedSquares(nums2)) // Expected output: [0, 1, 1, 4, 4, 9, 16]
}
