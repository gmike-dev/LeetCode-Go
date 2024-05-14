package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMaximumGold1(t *testing.T) {
	actual := getMaximumGold([][]int{
		{0, 6, 0},
		{5, 8, 7},
		{0, 9, 0},
	})
	expected := 24
	assert.Equal(t, expected, actual)
}

func TestGetMaximumGold2(t *testing.T) {
	actual := getMaximumGold([][]int{
		{1, 0, 7},
		{2, 0, 6},
		{3, 4, 5},
		{0, 3, 0},
		{9, 0, 20},
	})
	expected := 28
	assert.Equal(t, expected, actual)
}
