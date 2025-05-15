package main

// https://leetcode.cn/problems/move-zeroes
func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{1, 0, 1})
}

func moveZeroes(nums []int) {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			if nums[left] != nums[right] {
				nums[left], nums[right] = nums[right], nums[left]
			}
			left++
		}
		right++
	}
}
