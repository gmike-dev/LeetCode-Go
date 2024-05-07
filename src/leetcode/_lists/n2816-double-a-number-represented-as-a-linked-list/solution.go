package n2816_double_a_number_represented_as_a_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func doubleIt(head *ListNode) *ListNode {
	head = ReverseList(head)
	node := head
	val := 0
	for node != nil {
		val += node.Val * 2
		node.Val = val % 10
		val /= 10
		node = node.Next
	}
	head = ReverseList(head)
	if val == 0 {
		return head
	}
	return &ListNode{Val: val, Next: head}
}

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		prev, head, head.Next = head, head.Next, prev
	}
	return prev
}
