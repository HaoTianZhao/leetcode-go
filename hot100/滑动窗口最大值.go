package main

import (
	"fmt"
)

// https://leetcode.cn/problems/sliding-window-maximum
func main() {
	window := maxSlidingWindow([]int{7, 2, 4}, 2)
	fmt.Println(window)
}

/*
示例 1：

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
*/
func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0)

	queue := make([]int, 0)
	for i := 0; i < k; i++ {
		for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	result = append(result, nums[queue[0]])

	for i := k; i < len(nums); i++ {
		if queue[0] < i-k+1 {
			queue = queue[1:]
		}
		for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		result = append(result, nums[queue[0]])
	}

	return result
}
