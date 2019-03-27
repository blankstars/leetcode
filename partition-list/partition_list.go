package partition_list

/**
	https://leetcode-cn.com/problems/partition-list/

	题目:
	给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
	你应当保留两个分区中每个节点的初始相对位置。

	示例:
	输入: head = 1->4->3->2->5->2, x = 3
	输出: 1->2->2->4->3->5
 */

/**
* Definition for singly-linked list.
*/
type ListNode struct {
	Val int
	Next *ListNode
}

func partition(head *ListNode, x int) (*ListNode) {
	var lHead, lTail, rHead, rTail *ListNode

	if head == nil {
		return nil
	}

	for {
		if head == nil {
			break
		}
		if head.Val < x {
			if lTail == nil {
				lHead = head
			} else {
				lTail.Next = head
			}
			lTail = head
		} else {
			if rTail == nil {
				rHead = head
			} else {
				rTail.Next = head
			}
			rTail = head
		}
		head = head.Next
	}

	if lHead == nil {
		return rHead
	}
	if rHead == nil {
		return lHead
	}

	lTail.Next = rHead
	rTail.Next = nil

	return lHead
}