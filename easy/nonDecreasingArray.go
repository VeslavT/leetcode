/*Given an array with n integers, your task is to check if it could become non-decreasing by modifying at most 1 element.

We define an array is non-decreasing if array[i] <= array[i + 1] holds for every i (1 <= i < n).

Example 1:
Input: [4,2,3]
Output: True
Explanation: You could modify the first 4 to 1 to get a non-decreasing array.
Example 2:
Input: [4,2,1]
Output: False
Explanation: You can't get a non-decreasing array by modify at most one element.
Note: The n belongs to [1, 10,000].
*/
package main

import "fmt"

func checkPossibility(nums []int) bool {
	n := len(nums) - 1
	once := false
	var max *int
	for i := n; i > 0; i-- {
		if max != nil && (*max < nums[i-1]) {
			return false
		}
		if nums[i-1] > nums[i] {
			if once {
				return false
			}
			once = true
			max = &nums[i]
		}
	}
	return true
}

func main(){
	a := []int{3,4,2,3}
	fmt.Println(checkPossibility(a))
}
