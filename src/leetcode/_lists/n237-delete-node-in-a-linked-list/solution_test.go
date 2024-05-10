package leetcode

import (
	list "LeetCode-Go/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteNode1(t *testing.T) {
	node := list.ArrayToList([]int{5, 1, 9})
	deleteNode(node)
	assert.Equal(t, []int{1, 9}, list.ListToArray(node))
}

func TestDeleteNode2(t *testing.T) {
	node := list.ArrayToList([]int{1, 9})
	deleteNode(node)
	assert.Equal(t, []int{9}, list.ListToArray(node))
}
