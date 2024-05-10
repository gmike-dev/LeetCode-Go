package leetcode

import list "LeetCode-Go/collections"

type ListNode = list.ListNode

func removeNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := removeNodes(head.Next)
	if head.Val < next.Val {
		return next
	}
	head.Next = next
	return head
}
