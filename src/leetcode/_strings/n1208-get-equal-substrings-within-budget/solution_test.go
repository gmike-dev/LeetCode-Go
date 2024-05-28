package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqualSubstring1(t *testing.T) {
	assert.Equal(t, 3, equalSubstring("abcd", "bcdf", 3))
}

func TestEqualSubstring2(t *testing.T) {
	assert.Equal(t, 1, equalSubstring("abcd", "cdef", 3))
}

func TestEqualSubstring3(t *testing.T) {
	assert.Equal(t, 1, equalSubstring("abcd", "acde", 0))
}
