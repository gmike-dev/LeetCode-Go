package two_pass

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	left := head
	right := head.Next
	left.Next = nil
	for right != nil {
		for left != nil && left.Val < right.Val {
			left = left.Next
		}
		tmp := right
		right = right.Next
		tmp.Next = left
		left = tmp
	}
	return ReverseList(left)
}

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next = prev
		prev = tmp
	}
	return prev
}
