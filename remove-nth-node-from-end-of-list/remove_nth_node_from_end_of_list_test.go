package remove_nth_node_from_end_of_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func unconv(list *ListNode) ([]int) {
	s := make([]int, 0)
	for {
		if list == nil {
			break
		}
		s = append(s, list.Val)
		list = list.Next
	}
	return s
}

func conv(s []int) (*ListNode) {
	var head, tail *ListNode
	for _, v := range s {
		node := &ListNode{Val: v}
		if tail == nil {
			head = node
		} else {
			tail.Next = node
		}
		tail = node
	}
	return head
}

func testCase(t *testing.T, f func (*ListNode, int) *ListNode ) {
	assert.Equal(t, []int{1,2,3,5}, unconv(f(conv([]int{1,2,3,4,5}), 2)))
	assert.Equal(t, []int{}, unconv(f(conv([]int{1}), 1)))
}

func Test(t *testing.T) {
	testCase(t, removeNthFromEnd)
}

