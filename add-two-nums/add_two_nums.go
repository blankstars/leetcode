package add_two_num

/**
	题目:
	给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
	如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
	您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

	示例：
	输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
	输出：7 -> 0 -> 8
	原因：342 + 465 = 807
 */

/**
* Definition for singly-linked list.
*/
type ListNode struct {
	Val int
	Next *ListNode
}

func fmtSlice(l *ListNode) []int {
	var s []int

	if l == nil {
		return nil
	}

	s = append(s, l.Val)
	for l.Next != nil {
		s = append(s, l.Next.Val)
		l = l.Next
	}

	return s
}

func fmtList(src []int) *ListNode {
	var h, l *ListNode
	for _, s := range src {
		if l == nil {
			l = &ListNode{s, nil}
			h = l
		} else {
			l.Next = &ListNode{s, nil}
			l = l.Next
		}
	}
	return h
}

// 初等数学法: 使用变量来跟踪进位，并从包含最低有效位的表头开始模拟逐位相加的过程
// 时间复杂度O(max(M,N)), 其中M、N分别为l1与l2的长度
// 空间复杂度O(max(M,N)), 其中M、N分别为l1与l2的长度
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	l3 := &ListNode{}
	h := l3
	for (l1 == nil && l2 == nil) == false {
		if l1 == nil {
			val := l2.Val + carry
			if val >= 10 {
				carry = 1
				val = val%10
			} else {
				carry = 0
			}
			l3.Next = &ListNode{
				val,
				nil,
			}
			l2 = l2.Next
			l3 = l3.Next
		} else if l2 == nil {
			val := l1.Val + carry
			if val >= 10 {
				carry = 1
				val = val%10
			} else {
				carry = 0
			}
			l3.Next = &ListNode{
				val,
				nil,
			}
			l1 = l1.Next
			l3 = l3.Next
		} else {
			val := l1.Val + l2.Val + carry
			if val >= 10 {
				carry = 1
				val = val%10
			} else {
				carry = 0
			}
			l3.Next = &ListNode{
				val,
				nil,
			}
			l1 = l1.Next
			l2 = l2.Next
			l3 = l3.Next
		}
	}

	if carry != 0 {
		l3.Next = &ListNode{
			carry,
			nil,
		}
	}

	return h.Next
}