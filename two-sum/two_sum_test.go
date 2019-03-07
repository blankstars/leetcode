package two_sum

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func testCase(t *testing.T, f func ([]int, int)[]int) {
	assert.Equal(t, []int{0,1}, f([]int{2,7,11,15}, 9))
	assert.Equal(t, []int{1,2}, f([]int{3,2,4}, 6))
	assert.Equal(t, []int{0,1}, f([]int{3,3}, 6))
}

func TestForce(t *testing.T) {
	testCase(t, force)
}

func TestTwoHash(t *testing.T) {
	testCase(t, twoHash)
}

func TestOneHash(t *testing.T) {
	testCase(t, oneHash)
}
