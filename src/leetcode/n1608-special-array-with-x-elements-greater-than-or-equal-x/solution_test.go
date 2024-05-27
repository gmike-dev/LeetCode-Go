package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpecialArray1(t *testing.T) {
	assert.Equal(t, 2, specialArray([]int{3, 5}))
}

func TestSpecialArray2(t *testing.T) {
	assert.Equal(t, -1, specialArray([]int{0, 0}))
}

func TestSpecialArray3(t *testing.T) {
	assert.Equal(t, 3, specialArray([]int{0, 4, 3, 0, 4}))
}
