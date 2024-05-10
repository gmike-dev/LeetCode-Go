package leetcode

import list "LeetCode-Go/collections"

type ListNode = list.ListNode

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
