package n237_delete_node_in_a_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func ArrayToList(a []int) *ListNode {
	head := &ListNode{}
	node := head
	for _, x := range a {
		node.Next = &ListNode{Val: x}
		node = node.Next
	}
	return head.Next
}

func ListToArray(head *ListNode) []int {
	var result []int
	node := head
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}

func TestDeleteNode1(t *testing.T) {
	node := ArrayToList([]int{5, 1, 9})
	deleteNode(node)
	assert.Equal(t, []int{1, 9}, ListToArray(node))
}

func TestDeleteNode2(t *testing.T) {
	node := ArrayToList([]int{1, 9})
	deleteNode(node)
	assert.Equal(t, []int{9}, ListToArray(node))
}
