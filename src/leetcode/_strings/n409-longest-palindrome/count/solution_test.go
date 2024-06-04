package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindrome1(t *testing.T) {
	assert.Equal(t, 7, longestPalindrome("abccccdd"))
}

func TestLongestPalindrome2(t *testing.T) {
	assert.Equal(t, 1, longestPalindrome("a"))
}
