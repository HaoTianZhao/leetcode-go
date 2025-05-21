package main

// https://leetcode.cn/problems/first-missing-positive
func main() {
	firstMissingPositive([]int{1})

}

/*
输入：nums = [3,4,-1,1]
输出：2
解释：1 在数组中，但 2 没有。
*/
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		index := abs(nums[i]) - 1
		if index < n {
			nums[index] = -abs(nums[index])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
