package main

// https://leetcode.cn/problems/container-with-most-water
func main() {
	maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
}

func maxArea(height []int) int {
	maxArea := 0
	l, r := 0, len(height)-1
	for l < r {
		area := getHeight(height, l, r) * (r - l)
		if maxArea < area {
			maxArea = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxArea
}

func getHeight(height []int, i, j int) int {
	if height[i] < height[j] {
		return height[i]
	} else {
		return height[j]
	}
}
