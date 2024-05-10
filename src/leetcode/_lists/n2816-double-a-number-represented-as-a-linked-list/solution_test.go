package leetcode

import (
	list "LeetCode-Go/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoubleIt1(t *testing.T) {
	actual := doubleIt(list.ArrayToList([]int{1, 8, 9}))
	expect := []int{3, 7, 8}
	assert.Equal(t, expect, list.ListToArray(actual))
}

func TestDoubleIt2(t *testing.T) {
	actual := doubleIt(list.ArrayToList([]int{9, 9, 9}))
	expect := []int{1, 9, 9, 8}
	assert.Equal(t, expect, list.ListToArray(actual))
}

func TestDoubleIt3(t *testing.T) {
	actual := doubleIt(list.ArrayToList([]int{1}))
	expect := []int{2}
	assert.Equal(t, expect, list.ListToArray(actual))
}

func TestDoubleIt4(t *testing.T) {
	actual := doubleIt(list.ArrayToList([]int{5}))
	expect := []int{1, 0}
	assert.Equal(t, expect, list.ListToArray(actual))
}
