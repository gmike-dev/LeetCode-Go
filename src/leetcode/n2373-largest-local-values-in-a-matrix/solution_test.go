package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestLocal1(t *testing.T) {
	actual := largestLocal([][]int{
		{9, 9, 8, 1},
		{5, 6, 2, 6},
		{8, 2, 6, 4},
		{6, 2, 2, 2},
	})
	expected := [][]int{
		{9, 9},
		{8, 6},
	}
	assert.Equal(t, expected, actual)
}

func TestLargestLocal2(t *testing.T) {
	actual := largestLocal([][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 2, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	})
	expected := [][]int{
		{2, 2, 2},
		{2, 2, 2},
		{2, 2, 2},
	}
	assert.Equal(t, expected, actual)
}
