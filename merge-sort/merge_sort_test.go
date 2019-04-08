package merge_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testCase(t *testing.T, f func ([]int)[]int) {
	assert.Equal(t, []int{1,2,3,4,5}, f([]int{2,1,4,5,3}))
	assert.Equal(t, []int{0,1,2,3,4,5}, f([]int{2,1,4,5,3,0}))
	assert.Equal(t, []int{1}, f([]int{1}))
	assert.Equal(t, []int{}, f([]int{}))
}

func Test(t *testing.T) {
	testCase(t, mergeSort)
}

