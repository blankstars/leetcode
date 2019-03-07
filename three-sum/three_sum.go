package three_sum

import (
	"reflect"
	"sort"
)

/**
	https://leetcode-cn.com/problems/3sum/

	题目:
	给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

	注意:
	答案中不可以包含重复的三元组。

	示例:
	例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
	满足要求的三元组集合为：
	[
  		[-1, 0, 1],
  		[-1, -1, 2]
	]
 */

 // 问题拆分法，由三数和为0拆分为两数和为某target
 // 时间复杂度为O(N^2): 两数求和为O(N), 嵌套再遍历一遍则为O(N^2)
func divide(nums []int) [][]int {
	twoSum := func(ns []int, target int)[][2]int{
		var r [][2]int
		m := make(map[int]int)
		for i, n := range ns {
			if _, ok := m[target-n]; ok {
				r = append(r, [2]int{target-n, n})
			}
			m[n] = i
		}
		return r
	}

	var r [][]int
	sort.Slice(nums, func(i, j int)bool{return nums[i]<nums[j]})
	for i, n := range nums {
		if m := twoSum(nums[i+1:], -n); m != nil {
			re:
			for _, v := range m {
				vv := []int{n, v[0], v[1]}
				for _, rr := range r {
					if reflect.DeepEqual(rr, vv) {
						continue re
					}
				}
				r = append(r, vv)
			}
		}
	}

	return r
}

// 1.将数组排序
// 2.定义三个指针，i，j，k。
// 3.遍历i，那么这个问题就可以转化为在i之后的数组中寻找nums[j]+nums[k]=-nums[i]这个问题，
// 4.也就将三数之和问题转变为二数之和---（可以使用双指针)
func threePointer(nums []int) [][]int {
	sort.Slice(nums, func(i, j int)bool{return nums[i]<nums[j]})

	var res [][]int
	for i, n := range nums {
		if i == 0 || n > nums[i-1] {
			l := i+1
			r := len(nums)-1
			for l < r {
				s := n + nums[l] + nums[r]
				if s == 0 {
					res = append(res, []int{n, nums[l], nums[r]})
					l += 1
					r -= 1
					for l < r && nums[l] == nums[l-1] {
						l += 1
					}
					for l < r && nums[r] == nums[r+1] {
						r -= 1
					}
				} else if s > 0 {
					r -= 1
				} else {
					l += 1
				}
			}
		}
	}

	return res
}
