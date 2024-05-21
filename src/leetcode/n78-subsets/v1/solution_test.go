package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsets1(t *testing.T) {
	assert.Equal(t, [][]int{{}, {3}, {2}, {2, 3}, {1}, {1, 3}, {1, 2}, {1, 2, 3}}, subsets([]int{1, 2, 3}))
}

func TestSubsets2(t *testing.T) {
	assert.Equal(t, [][]int{{}, {0}}, subsets([]int{0}))
}

func TestSubsets3(t *testing.T) {
	assert.Equal(t, [][]int{{}, {7}, {5}, {5, 7}, {3}, {3, 7}, {3, 5}, {3, 5, 7},
		{0}, {0, 7}, {0, 5}, {0, 5, 7}, {0, 3}, {0, 3, 7}, {0, 3, 5}, {0, 3, 5, 7},
		{9}, {9, 7}, {9, 5}, {9, 5, 7}, {9, 3}, {9, 3, 7}, {9, 3, 5}, {9, 3, 5, 7},
		{9, 0}, {9, 0, 7}, {9, 0, 5}, {9, 0, 5, 7}, {9, 0, 3}, {9, 0, 3, 7},
		{9, 0, 3, 5}, {9, 0, 3, 5, 7}}, subsets([]int{9, 0, 3, 5, 7}))
}
