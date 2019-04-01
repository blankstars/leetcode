package friend_circles

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testCase(t *testing.T, f func ([][]int)int) {
	assert.Equal(t, 2, f([][]int{{1,1,0},{1,1,0},{0,0,1}}))
	assert.Equal(t, 1, f([][]int{{1,0,0,1},{0,1,1,0},{0,1,1,1},{1,0,1,1}}))
}

func Test(t *testing.T) {
	testCase(t, findCircleNum)
}

