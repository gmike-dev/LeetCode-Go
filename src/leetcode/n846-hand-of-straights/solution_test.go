package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsNStraightHand1(t *testing.T) {
	assert.Equal(t, true, isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
}

func TestIsNStraightHand2(t *testing.T) {
	assert.Equal(t, false, isNStraightHand([]int{1, 2, 3, 4, 5}, 4))
}
