package n2816_double_a_number_represented_as_a_linked_list

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

func TestDoubleIt1(t *testing.T) {
	actual := doubleIt(ArrayToList([]int{1, 8, 9}))
	expect := []int{3, 7, 8}
	assert.Equal(t, expect, ListToArray(actual))
}

func TestDoubleIt2(t *testing.T) {
	actual := doubleIt(ArrayToList([]int{9, 9, 9}))
	expect := []int{1, 9, 9, 8}
	assert.Equal(t, expect, ListToArray(actual))
}

func TestDoubleIt3(t *testing.T) {
	actual := doubleIt(ArrayToList([]int{1}))
	expect := []int{2}
	assert.Equal(t, expect, ListToArray(actual))
}

func TestDoubleIt4(t *testing.T) {
	actual := doubleIt(ArrayToList([]int{5}))
	expect := []int{1, 0}
	assert.Equal(t, expect, ListToArray(actual))
}
