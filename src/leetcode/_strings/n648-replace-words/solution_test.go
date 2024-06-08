package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceWords1(t *testing.T) {
	assert.Equal(t, "the cat was rat by the bat",
		replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
}

func TestReplaceWords2(t *testing.T) {
	assert.Equal(t, "a a b c",
		replaceWords([]string{"a", "b", "c"}, "aadsfasf absbs bbab cadsfafs"))
}
