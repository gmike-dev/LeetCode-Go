package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_smallestEquivalentString1(t *testing.T) {
	assert.Equal(t, "makkek", smallestEquivalentString("parker", "morris", "parser"))
}

func Test_smallestEquivalentString2(t *testing.T) {
	assert.Equal(t, "hdld", smallestEquivalentString("hello", "world", "hold"))
}

func Test_smallestEquivalentString3(t *testing.T) {
	assert.Equal(t, "aauaaaaada", smallestEquivalentString("leetcode", "programs", "sourcecode"))
}
