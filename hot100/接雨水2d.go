package main

// https://leetcode.cn/problems/trapping-rain-water
func main() {

}

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	result := 0
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]

	for left < right {
		if leftMax < rightMax {
			result += leftMax - height[left]
			left++
			leftMax = max(height[left], leftMax)
		} else {
			result += rightMax - height[right]
			right--
			rightMax = max(height[right], rightMax)
		}
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
