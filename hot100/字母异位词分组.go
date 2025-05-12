package main

import (
	"sort"
	"strings"
)

// https://leetcode.cn/problems/group-anagrams
func main() {
	/*	给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

			字母异位词 是由重新排列源单词的所有字母得到的一个新单词。



			示例 1:

		输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
			输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
			示例 2:

			输入: strs = [""]
			输出: [[""]]
			示例 3:

			输入: strs = ["a"]
			输出: [["a"]]


			提示：

			1 <= strs.length <= 104
			0 <= strs[i].length <= 100
			strs[i] 仅包含小写字母*/

}

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	group := make(map[string][]string)
	for _, str := range strs {
		split := strings.Split(str, "")
		sort.Slice(split, func(i, j int) bool {
			return split[i] < split[j]
		})
		key := strings.Join(split, "")
		if _, ok := group[key]; ok {
			group[key] = append(group[key], str)
		} else {
			group[key] = []string{str}
		}
	}
	for _, v := range group {
		temp := make([]string, 0, len(v))
		temp = append(temp, v...)
		res = append(res, temp)
	}
	return res
}
