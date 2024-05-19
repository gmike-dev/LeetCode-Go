package n3068_find_the_maximum_sum_of_node_values

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximumValueSum1(t *testing.T) {
	assert.Equal(t, int64(6), maximumValueSum([]int{1, 2, 1}, 3, [][]int{{0, 1}, {0, 2}}))
}

func TestMaximumValueSum2(t *testing.T) {
	assert.Equal(t, int64(9), maximumValueSum([]int{2, 3}, 7, [][]int{{0, 1}}))
}

func TestMaximumValueSum3(t *testing.T) {
	assert.Equal(t, int64(42), maximumValueSum([]int{7, 7, 7, 7, 7, 7}, 3, [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}}))
}
