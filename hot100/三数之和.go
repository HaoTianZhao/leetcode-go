package main

import (
	"sort"
)

// https://leetcode.cn/problems/3sum/description
func main() {

}

// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	result := make([][]int, 0)

	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			target := 0 - nums[i] - nums[j]
			for target < nums[k] && j < k {
				k--
			}
			if j == k {
				break
			}
			if nums[k] == target {
				result = append(result, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return result
}
