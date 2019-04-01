package remove_nth_node_from_end_of_list

/**
	https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

	题目:
	给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

	示例：
	给定一个链表: 1->2->3->4->5, 和 n = 2.
	当删除了倒数第二个节点后，链表变为 1->2->3->5.

	说明：
	给定的 n 保证是有效的。

	进阶：
	你能尝试使用一趟扫描实现吗？
 */

/**
* Definition for singly-linked list.
*/

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	list := head
	buff := make([]*ListNode, n+1)
	beginBuff := 0
	for i := range buff {
		if list == nil {
			return head.Next
		}
		buff[i] = list
		list = list.Next
	}

	for {
		if list == nil {
			break
		}

		buff[beginBuff] = list
		beginBuff++
		if beginBuff == n+1 {
			beginBuff = 0
		}
		list = list.Next
	}

	buff[beginBuff].Next = buff[beginBuff].Next.Next

	return head
}