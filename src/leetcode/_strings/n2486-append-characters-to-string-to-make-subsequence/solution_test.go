package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppendCharacters1(t *testing.T) {
	assert.Equal(t, 4, appendCharacters("coaching", "coding"))
}

func TestAppendCharacters2(t *testing.T) {
	assert.Equal(t, 0, appendCharacters("abcde", "a"))
}

func TestAppendCharacters3(t *testing.T) {
	assert.Equal(t, 5, appendCharacters("z", "abcde"))
}
