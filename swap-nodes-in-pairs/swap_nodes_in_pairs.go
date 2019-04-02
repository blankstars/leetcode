package swap_nodes_in_pairs

/**
	https://leetcode-cn.com/problems/swap-nodes-in-pairs/

	题目:
	给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
	你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

	示例:
	给定 1->2->3->4, 你应该返回 2->1->4->3.
 */

/**
* Definition for singly-linked list.
*/
type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var left, cur, right *ListNode
	cur = head
	right = cur.Next
	head = right
	for {
		cur.Next = right.Next
		right.Next = cur
		if left != nil {
			left.Next = right
		}
		left = cur
		cur = cur.Next
		if cur == nil {
			break
		}
		right = cur.Next
		if right == nil {
			break
		}
	}
	return head
}