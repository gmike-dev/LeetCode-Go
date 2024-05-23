package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBeautifulSubsets1(t *testing.T) {
	assert.Equal(t, 4, beautifulSubsets([]int{2, 4, 6}, 2))
}

func TestBeautifulSubsets2(t *testing.T) {
	assert.Equal(t, 1, beautifulSubsets([]int{1}, 1))
}

func TestBeautifulSubsets3(t *testing.T) {
	assert.Equal(t, 23, beautifulSubsets([]int{10, 4, 5, 7, 2, 1}, 3))
}

func TestBeautifulSubsets4(t *testing.T) {
	assert.Equal(t, 8039, beautifulSubsets([]int{10, 4, 5, 7, 2, 1, 10, 4, 5, 7, 2, 1, 10, 4, 5, 7, 2, 1, 1, 1}, 3))
}

func TestBeautifulSubsets5(t *testing.T) {
	assert.Equal(t, 37439, beautifulSubsets([]int{10, 4, 5, 7, 2, 1, 10, 4, 5, 7, 2, 1, 10, 4, 5, 7, 2, 1, 1, 1}, 1))
}
