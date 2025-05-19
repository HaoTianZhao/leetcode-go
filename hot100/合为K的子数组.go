package main

// https://leetcode.cn/problems/subarray-sum-equals-k
func main() {

}

/*
示例 1：

输入：nums = [1,-1,0], k = 0
输出：3
示例 2：

输入：nums = [1, 2, 3], k = 3
输出：2
*/
func subarraySum(nums []int, k int) int {

	result := 0

	pre := 0
	m := map[int]int{0: 1}
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if count, ok := m[pre-k]; ok {
			result += count
		}
		m[pre]++
	}

	return result
}
