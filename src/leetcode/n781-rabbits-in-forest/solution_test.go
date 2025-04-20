package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsNStraightHand1(t *testing.T) {
	assert.Equal(t, 5, numRabbits([]int{1, 1, 2}))
}

func TestIsNStraightHand2(t *testing.T) {
	assert.Equal(t, 11, numRabbits([]int{10, 10, 10}))
}
