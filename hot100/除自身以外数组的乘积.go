package main

// https://leetcode.cn/problems/product-of-array-except-self
func main() {

}

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	result[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i]
	}
	right := 1
	for i := len(nums) - 1; i > 0; i-- {
		result[i] = result[i-1] * right
		right = right * nums[i]
	}
	result[0] = right

	return result
}
