package main

// https://leetcode.cn/problems/minimum-window-substring
func main() {
	minWindow("ADOBECODEBANC", "ABC")
}

func minWindow(s string, t string) string {
	left, right := 0, 0
	rightM := make(map[byte]int)
	for i := range t {
		rightM[t[i]]++
	}
	leftM := make(map[byte]int)
	less := len(rightM)

	result := ""
	for right < len(s) {
		if less > 0 {
			leftM[s[right]]++
			if _, ok := rightM[s[right]]; ok && leftM[s[right]] == rightM[s[right]] {
				less--
			}
			right++
		}
		if less == 0 {
			for less == 0 {
				if num, ok := leftM[s[left]]; !ok || num > rightM[s[left]] {
					leftM[s[left]]--
					left++
				} else {
					break
				}
			}

			str := s[left:right]
			if result == "" || len(str) < len(result) {
				result = str
			}
			leftM[s[left]]--
			left++
			less++
		}

	}
	return result
}
