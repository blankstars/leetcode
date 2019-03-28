package max_consecutive_ones_iii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testCase(t *testing.T, f func (A []int, K int) int) {
	assert.Equal(t, 6, f([]int{1,1,1,0,0,0,1,1,1,1,0}, 2))
	assert.Equal(t, 10, f([]int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}, 3))
}

func Test(t *testing.T) {
	testCase(t, longestOnes)
}
