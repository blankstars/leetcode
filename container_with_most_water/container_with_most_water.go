package container_with_most_water

/**
	https://leetcode-cn.com/problems/container-with-most-water/
 */

func maxArea(height []int) int {
	var maxArea int
	for i, v := range height {
		for ii, vv := range height[i+1:] {
			h := vv
			if v < vv {
				h = v
			}
			w := ii + 1
			area := h * w
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func maxAreaOptimize(height []int) int {
	maxArea := 0
	leftIdx := 0
	rightIdx := len(height)-1
	for {
		if leftIdx == rightIdx {
			break
		}
		var h int
		leftHeight := height[leftIdx]
		rightHeight := height[rightIdx]
		w := rightIdx-leftIdx
		if leftHeight < rightHeight {
			h = leftHeight
			leftIdx++
		} else {
			h = rightHeight
			rightIdx--
		}
		area := w * h
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
