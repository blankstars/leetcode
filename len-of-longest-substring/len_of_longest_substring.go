package len_of_longest_substring

/**
	https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

	给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

	示例 1:
	输入: "abcabcbb"
	输出: 3
	解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

	示例 2:
	输入: "bbbbb"
	输出: 1
	解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

	示例 3:
	输入: "pwwkew"
	输出: 3
	解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
	请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 */

 // 滑动窗口法
 // 时间复杂度O(N)
func lengthOfLongestSubstring(s string) int {
	var (
		b, e int // 子串开始/结束索引
		l int // 子串长度
		c int32
	)
	set := make(map[int32]int)

	for e, c = range s {
		if v, ok := set[c]; ok && v >= b {
			if l < e-b {
				l = e-b
			}
			b = v+1
		}
		set[c] = e
	}

	if l < e-b+1  {
		l = e-b+1
	}
	if len(s) < 2 {
		l = len(s)
	}

	return l
}