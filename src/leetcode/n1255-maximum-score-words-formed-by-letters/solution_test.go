package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxScoreWords1(t *testing.T) {
	assert.Equal(t, 23, maxScoreWords(
		[]string{"dog", "cat", "dad", "good"},
		[]byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'},
		[]int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}

func TestMaxScoreWords2(t *testing.T) {
	assert.Equal(t, 27, maxScoreWords(
		[]string{"xxxz", "ax", "bx", "cx"},
		[]byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'},
		[]int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10}))
}

func TestMaxScoreWords3(t *testing.T) {
	assert.Equal(t, 0, maxScoreWords(
		[]string{"leetcode"},
		[]byte{'l', 'e', 't', 'c', 'o', 'd'},
		[]int{0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}))
}
