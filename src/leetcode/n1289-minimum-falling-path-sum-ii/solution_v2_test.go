package n1289_minimum_falling_path_sum_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinFallingPathSumV2_1(t *testing.T) {
	actual := minFallingPathSumV2([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	expected := 13
	assert.Equal(t, expected, actual)
}

func TestMinFallingPathSumV2_2(t *testing.T) {
	actual := minFallingPathSumV2([][]int{
		{7},
	})
	expected := 7
	assert.Equal(t, expected, actual)
}
