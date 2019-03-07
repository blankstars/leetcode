package two_sum

/**
	https://leetcode-cn.com/problems/two-sum/

	题目:
	给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
	你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

	示例:
	给定 nums = [2, 7, 11, 15], target = 9
	因为 nums[0] + nums[1] = 2 + 7 = 9
	所以返回 [0, 1]
 */

 // 暴力法: 时间复杂度O(n^2), 空间复杂度O(1)
func force(nums []int, target int) []int{
    for i, n := range nums {
        for j, m := range nums[i+1:] {
            if n+m == target {
                return []int{i, i+1+j}
            }
        }
    }
    return nil
}

// 两遍hash表: 时间复杂度O(n), 空间复杂度O(n)
func twoHash(nums []int, target int) []int {
	// 为了解决{3,3} 6这种case
	halfTargetIndex := -1

	// 将nums映射为map组
	m := make(map[int]int)
	for i, n := range nums {
	    m[n] = i
	    if 2*n == target {
	    	if halfTargetIndex < 0 {
				halfTargetIndex = i
			} else {
				return []int{halfTargetIndex, i}
			}
		}
    }

	// 再遍历map组查看target-n是否也存在
	for n, i := range m {
		if j, ok := m[target-n]; ok {
			if i == j {
				continue
			}
			return []int{i, j}
		}
	}

    return nil
}

// 一遍hash表: 时间复杂度O(n), 空间复杂度O(n)
func oneHash(nums []int, target int) []int {
	// 将nums映射为map组, 边映射边查看是否满足条件
	m := make(map[int]int)
	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return []int{j, i}
		}
		m[n] = i
	}

    return nil
}


