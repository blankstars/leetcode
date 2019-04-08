package merge_sort

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	i := len(arr)/2
	left := mergeSort(arr[0:i])
	right := mergeSort(arr[i:])
	return merge(left, right)
}

func merge(left, right []int) []int{
	result := make([]int, 0)
	l, r := 0, 0
	lLen, rLen := len(left), len(right)
	for l < lLen && r < rLen {
		if left[l] > right[r] {
			result = append(result, right[r])
			r++
			continue
		}
		result = append(result, left[l])
		l++
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}