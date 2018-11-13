package easy

import "fmt"
/* Given an array of integers, return indices of the two numbers such that they add up to a specific target.
	You may assume that each input would have exactly one solution, and you may not use the same element twice.
	Example:
		Given nums = [2, 7, 11, 15], target = 9,
		Because nums[0] + nums[1] = 2 + 7 = 9,
		return [0, 1].
*/
func twoSum(nums []int, target int) []int {
	for i := range nums{
		if key:= intInSlice(target - nums[i], nums[i+1:]); key != len(nums[i+1:]) {
			return []int{i, key+1+i}
		}
	}
	return []int{}
}

func intInSlice(item int, scope []int) int {
	for key := range scope {
		if item == scope[key] {
			return key
		}
	}
	return len(scope)
}

func main() {
	t := []int{3,2,4}
	fmt.Println(twoSum(t, 6))
}
