package main

// https://leetcode.cn/problems/longest-consecutive-sequence
func main() {
	longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})

}

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool, len(nums))
	for _, num := range nums {
		numSet[num] = true
	}

	maxLength := 0
	for num := range numSet {
		if numSet[num-1] {
			continue
		}
		currentLength := 1
		for numSet[num+1] {
			num++
			currentLength++
		}
		if maxLength < currentLength {
			maxLength = currentLength
		}
	}
	return maxLength
}
