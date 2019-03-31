package container_with_most_water

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testCase(t *testing.T, f func ([]int)int) {
	assert.Equal(t, 49, f([]int{1,8,6,2,5,4,8,3,7}))
}

func Test(t *testing.T) {
	testCase(t, maxArea)
	testCase(t, maxAreaOptimize)
}

