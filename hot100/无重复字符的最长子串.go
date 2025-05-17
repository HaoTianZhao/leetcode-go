package main

import (
	"fmt"
)

// https://leetcode.cn/problems/longest-substring-without-repeating-characters
func main() {
	fmt.Println(lengthOfLongestSubstring("bpfbhmipx"))

}

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	maxLength := 0
	strMap := make(map[byte]bool)

	right := 0
	for left := 0; left < len(s); left++ {
		if left > 0 {
			delete(strMap, s[left-1])
		}

		for right < len(s) && !strMap[s[right]] {
			strMap[s[right]] = true
			right++
		}

		maxLength = max(maxLength, right-left)
	}
	return maxLength
}
