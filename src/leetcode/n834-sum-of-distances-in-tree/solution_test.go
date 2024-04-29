package n834_sum_of_distances_in_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumOfDistancesInTree1(t *testing.T) {
	actual := sumOfDistancesInTree(6, [][]int{
		{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5},
	})
	expected := []int{8, 12, 6, 10, 10, 10}
	assert.Equal(t, expected, actual)
}

func TestSumOfDistancesInTree2(t *testing.T) {
	actual := sumOfDistancesInTree(1, [][]int{})
	expected := []int{0}
	assert.Equal(t, expected, actual)
}
