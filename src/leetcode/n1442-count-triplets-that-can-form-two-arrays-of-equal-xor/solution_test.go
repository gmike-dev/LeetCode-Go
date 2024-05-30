package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountTriplets1(t *testing.T) {
	assert.Equal(t, 4, countTriplets([]int{2, 3, 1, 6, 7}))
}

func TestCountTriplets2(t *testing.T) {
	assert.Equal(t, 10, countTriplets([]int{1, 1, 1, 1, 1}))
}
