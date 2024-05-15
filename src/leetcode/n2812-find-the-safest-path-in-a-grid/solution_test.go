package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximumSafenessFactor1(t *testing.T) {
	actual := maximumSafenessFactor([][]int{
		{1, 0, 0},
		{0, 0, 0},
		{0, 0, 1},
	})
	expected := 0
	assert.Equal(t, expected, actual)
}

func TestMaximumSafenessFactor2(t *testing.T) {
	actual := maximumSafenessFactor([][]int{
		{0, 0, 1},
		{0, 0, 0},
		{0, 0, 0},
	})
	expected := 2
	assert.Equal(t, expected, actual)
}

func TestMaximumSafenessFactor3(t *testing.T) {
	actual := maximumSafenessFactor([][]int{
		{0, 0, 0, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 0, 0, 0},
	})
	expected := 2
	assert.Equal(t, expected, actual)
}

func TestMaximumSafenessFactor4(t *testing.T) {
	actual := maximumSafenessFactor([][]int{
		{1},
	})
	expected := 0
	assert.Equal(t, expected, actual)
}
