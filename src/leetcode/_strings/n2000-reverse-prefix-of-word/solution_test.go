package n2000_reverse_prefix_of_word

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReversePrefix1(t *testing.T) {
	actual := reversePrefix("abcdefd", 'd')
	expect := "dcbaefd"
	assert.Equal(t, expect, actual)
}

func TestReversePrefix2(t *testing.T) {
	actual := reversePrefix("xyxzxe", 'z')
	expect := "zxyxxe"
	assert.Equal(t, expect, actual)
}

func TestReversePrefix3(t *testing.T) {
	actual := reversePrefix("abcd", 'z')
	expect := "abcd"
	assert.Equal(t, expect, actual)
}
