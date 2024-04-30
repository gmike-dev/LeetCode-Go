package n1915_number_of_wonderful_substrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWonderfulSubstrings1(t *testing.T) {
	actual := wonderfulSubstrings("aba")
	var expected int64 = 4
	assert.Equal(t, expected, actual)
}

func TestWonderfulSubstrings2(t *testing.T) {
	actual := wonderfulSubstrings("aabb")
	var expected int64 = 9
	assert.Equal(t, expected, actual)
}

func TestWonderfulSubstrings3(t *testing.T) {
	actual := wonderfulSubstrings("he")
	var expected int64 = 2
	assert.Equal(t, expected, actual)
}
