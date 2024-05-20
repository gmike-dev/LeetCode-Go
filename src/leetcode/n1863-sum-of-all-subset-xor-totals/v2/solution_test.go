package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsetXORSum1(t *testing.T) {
	assert.Equal(t, 6, subsetXORSum([]int{1, 3}))
}

func TestSubsetXORSum2(t *testing.T) {
	assert.Equal(t, 28, subsetXORSum([]int{5, 1, 6}))
}

func TestSubsetXORSum3(t *testing.T) {
	assert.Equal(t, 480, subsetXORSum([]int{3, 4, 5, 6, 7, 8}))
}
