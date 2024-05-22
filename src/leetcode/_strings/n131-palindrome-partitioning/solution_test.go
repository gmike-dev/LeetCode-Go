package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartition1(t *testing.T) {
	assert.Equal(t, [][]string{{"a", "a", "b"}, {"aa", "b"}}, partition("aab"))
}

func TestPartition2(t *testing.T) {
	assert.Equal(t, [][]string{{"a"}}, partition("a"))
}
