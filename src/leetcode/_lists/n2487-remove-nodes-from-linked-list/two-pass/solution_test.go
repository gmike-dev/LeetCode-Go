package two_pass

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

func TestRemoveNodes1(t *testing.T) {
	actual := removeNodes(ArrayToList([]int{5, 2, 13, 3, 8}))
	expect := []int{13, 8}
	assert.Equal(t, expect, ListToArray(actual))
}

func TestRemoveNodes2(t *testing.T) {
	actual := removeNodes(ArrayToList([]int{1, 1, 1, 1}))
	expect := []int{1, 1, 1, 1}
	assert.Equal(t, expect, ListToArray(actual))
}

func TestRemoveNodes3(t *testing.T) {
	actual := removeNodes(ArrayToList([]int{1}))
	expect := []int{1}
	assert.Equal(t, expect, ListToArray(actual))
}

func TestRemoveNodes4(t *testing.T) {
	actual := removeNodes(ArrayToList([]int{1, 2}))
	expect := []int{2}
	assert.Equal(t, expect, ListToArray(actual))
}
