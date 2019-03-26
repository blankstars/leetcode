package longest_palindrome

/**
	https://leetcode-cn.com/problems/longest-palindromic-substring/

	题目:
	给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

	示例 1：
	输入: "babad"
	输出: "bab"
	注意: "aba" 也是一个有效答案。

	示例 2：
	输入: "cbbd"
	输出: "bb"
 */

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	lps := make([][]bool, len(s), len(s))
	for i := range lps {
		lps[i] = make([]bool, len(s), len(s))
	}
	maxLen := 1
	start := 0

	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if j - i + 1 <= 2 {
				lps[i][j] = s[i] == s[j]
			} else {
				lps[i][j] = lps[i+1][j-1] && s[i] == s[j]
			}
			if lps[i][j] && j - i + 1 > maxLen {
				maxLen = j - i + 1
				start = i
			}
		}
	}

	return s[start:start+maxLen]
}