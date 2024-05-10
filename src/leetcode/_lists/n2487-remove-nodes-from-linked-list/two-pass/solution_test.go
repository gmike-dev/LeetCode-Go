package leetcode

import (
	list "LeetCode-Go/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveNodes1(t *testing.T) {
	actual := removeNodes(list.ArrayToList([]int{5, 2, 13, 3, 8}))
	expect := []int{13, 8}
	assert.Equal(t, expect, list.ListToArray(actual))
}

func TestRemoveNodes2(t *testing.T) {
	actual := removeNodes(list.ArrayToList([]int{1, 1, 1, 1}))
	expect := []int{1, 1, 1, 1}
	assert.Equal(t, expect, list.ListToArray(actual))
}

func TestRemoveNodes3(t *testing.T) {
	actual := removeNodes(list.ArrayToList([]int{1}))
	expect := []int{1}
	assert.Equal(t, expect, list.ListToArray(actual))
}

func TestRemoveNodes4(t *testing.T) {
	actual := removeNodes(list.ArrayToList([]int{1, 2}))
	expect := []int{2}
	assert.Equal(t, expect, list.ListToArray(actual))
}
