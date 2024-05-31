package leetcode

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestSingleNumber1(t *testing.T) {
	actual := singleNumber([]int{1, 2, 1, 3, 2, 5})
	expect := []int{3, 5}
	slices.Sort(actual)
	assert.Equal(t, expect, actual)
}

func TestSingleNumber2(t *testing.T) {
	actual := singleNumber([]int{-1, 0})
	expect := []int{-1, 0}
	slices.Sort(actual)
	assert.Equal(t, expect, actual)
}

func TestSingleNumber3(t *testing.T) {
	actual := singleNumber([]int{0, 1})
	expect := []int{0, 1}
	slices.Sort(actual)
	assert.Equal(t, expect, actual)
}
