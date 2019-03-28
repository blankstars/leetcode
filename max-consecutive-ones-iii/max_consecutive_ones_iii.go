package max_consecutive_ones_iii

/**
	https://leetcode-cn.com/problems/max-consecutive-ones-iii/

	题目:
	给定一个由若干 0 和 1 组成的数组 A，我们最多可以将 K 个值从 0 变成 1 。
	返回仅包含 1 的最长（连续）子数组的长度。

	示例 1：
	输入：A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
	输出：6
	解释：
	[1,1,1,0,0,1,1,1,1,1,1]
	粗体数字从 0 翻转到 1，最长的子数组长度为 6。

	示例 2：
	输入：A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
	输出：10
	解释：
	[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
	粗体数字从 0 翻转到 1，最长的子数组长度为 10。
 */

func longestOnes(A []int, K int) int {
	var cur int // 当前滑动窗口中1的个数
	var max int // 符合条件的最大长度
	var start, end int // 滑动窗口起止
	var count [2]int // count[0]为滑动窗口中0的个数, count[1]为滑动窗口中1的个数
	for {
		// 退出条件
		if end  >= len(A) {
			break
		}

		count[A[end]]++ // 滑动窗口中0/1个数增加1
		if A[end] == 1 { // 若新增加的元素为1，则更新滑动窗口中1的个数
			cur = count[A[end]]
		}
		if end - start + 1 - cur > K { // 若新增元素后，滑动窗口中不符合条件，则更新滑动窗口的起点和相应的count数
			count[A[start]]--
			start++
		}

		if max < cur + K {
			max = cur + K
		}

		end++
	}

	if max > len(A) {
		max = len(A)
	}

	return max
}