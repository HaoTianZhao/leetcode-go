package main

import (
	"sort"
)

// https://leetcode.cn/problems/merge-intervals/description
func main() {
	merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
}

/*
示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

[1,6],[9,10],[7,15]
*/
func merge(intervals [][]int) [][]int {
	result := make([][]int, 0)
	if len(intervals) == 0 {
		return result
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	index := 0
	result = append(result, []int{intervals[0][0], intervals[0][1]})
	for i := 1; i < len(intervals); i++ {
		left := intervals[i][0]
		right := intervals[i][1]
		if left <= result[index][1] {
			if right > result[index][1] {
				result[index][1] = right
			}
		} else {
			result = append(result, []int{left, right})
			index++
		}
	}
	return result
}
