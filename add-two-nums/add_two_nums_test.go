package add_two_num

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testCase(t *testing.T, f func (*ListNode, *ListNode) *ListNode) {
	assert.Equal(t, []int{7,0,8}, fmtSlice(f(fmtList([]int{2,4,3}), fmtList([]int{5,6,4}))))
}

func Test(t *testing.T) {
	testCase(t, addTwoNumbers)
}
