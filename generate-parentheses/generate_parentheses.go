package generate_parentheses

/**
	https://leetcode-cn.com/problems/generate-parentheses/

	题目:
	给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

	例如，给出 n = 3，生成结果为：

	[
  		"((()))",
  		"(()())",
  		"(())()",
  		"()(())",
  		"()()()"
	]
 */

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var ans string
	var left, right int
	res = generate(res, ans, n, left, right)
	return res
}

// res: 当前满足条件的括号串
// ans: 目前括号串
// n: 括号对数量
// left: 左括号数据
// right: 右括号数量
func generate(res []string, ans string, n int, left int, right int) []string {
	if right == n {
		res = append(res, ans)
	}

	if left < n {
		res = generate(res, ans+"(", n, left+1, right)
	}
	if right < left {
		res = generate(res, ans+")", n, left, right+1)
	}
	return res
}