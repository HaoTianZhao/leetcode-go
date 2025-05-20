package main

// https://leetcode.cn/problems/maximum-subarray
func main() {
	maxSubArray([]int{-2, -1})

}

/*
示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
*/
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
		if pre+nums[i] > nums[i] {
			pre = pre + nums[i]
		} else {
			pre = nums[i]
		}
		if pre > max {
			max = pre
		}
	}
	return max
}
