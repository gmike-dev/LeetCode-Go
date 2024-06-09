package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubarraysDivByK1(t *testing.T) {
	assert.Equal(t, 7, subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
}

func TestSubarraysDivByK2(t *testing.T) {
	assert.Equal(t, 0, subarraysDivByK([]int{5}, 9))
}

func TestSubarraysDivByK3(t *testing.T) {
	assert.Equal(t, 2, subarraysDivByK([]int{-1, 2, 9}, 2))
}

func TestSubarraysDivByK4(t *testing.T) {
	assert.Equal(t, 2, subarraysDivByK([]int{2, -2, 2, -4}, 6))
}

func TestSubarraysDivByK5(t *testing.T) {
	assert.Equal(t, 1, subarraysDivByK([]int{7, 4, -10}, 5))
}
