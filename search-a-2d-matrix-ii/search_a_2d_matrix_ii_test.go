package search_a_2d_matrix_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testCase(t *testing.T, f func (matrix [][]int, target int) bool) {
	matrix := [][]int{
		{1, 4,  7, 11, 15},
		{2, 5,  8, 12, 19},
		{3, 6,  9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	assert.Equal(t, true, f(matrix, 5))
	assert.Equal(t, false, f(matrix, 20))

	matrix = [][]int{{}}
	assert.Equal(t, false, f(matrix, 1))
}

func Test(t *testing.T) {
	testCase(t, searchMatrix)
}
