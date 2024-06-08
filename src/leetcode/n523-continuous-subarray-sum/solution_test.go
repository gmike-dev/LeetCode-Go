package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckSubarraySum1(t *testing.T) {
	assert.True(t, checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))
}

func TestCheckSubarraySum2(t *testing.T) {
	assert.True(t, checkSubarraySum([]int{23, 2, 6, 4, 7}, 6))
}

func TestCheckSubarraySum3(t *testing.T) {
	assert.False(t, checkSubarraySum([]int{23, 2, 6, 4, 7}, 13))
}

func TestCheckSubarraySum4(t *testing.T) {
	assert.True(t, checkSubarraySum([]int{23, 2, 4, 6, 6}, 7))
}

func TestCheckSubarraySum5(t *testing.T) {
	assert.True(t, checkSubarraySum([]int{1, 2, 3}, 5))
}

func TestCheckSubarraySum6(t *testing.T) {
	assert.False(t, checkSubarraySum([]int{1, 0}, 2))
}
