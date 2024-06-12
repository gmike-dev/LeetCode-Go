package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortColors1(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	assert.Equal(t, []int{0, 0, 1, 1, 2, 2}, nums)
}

func TestSortColors2(t *testing.T) {
	nums := []int{2, 0, 1}
	sortColors(nums)
	assert.Equal(t, []int{0, 1, 2}, nums)
}

func TestSortColors3(t *testing.T) {
	nums := []int{1, 2}
	sortColors(nums)
	assert.Equal(t, []int{1, 2}, nums)
}

func TestSortColors4(t *testing.T) {
	nums := []int{1, 2, 0}
	sortColors(nums)
	assert.Equal(t, []int{0, 1, 2}, nums)
}
