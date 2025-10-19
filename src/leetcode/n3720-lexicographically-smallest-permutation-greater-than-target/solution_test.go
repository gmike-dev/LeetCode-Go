package n3720_lexicographically_smallest_permutation_greater_than_target

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lexGreaterPermutation1(t *testing.T) {
	assert.Equal(t, "bca", lexGreaterPermutation("abc", "bba"))
}

func Test_lexGreaterPermutation2(t *testing.T) {
	assert.Equal(t, "eelt", lexGreaterPermutation("leet", "code"))
}

func Test_lexGreaterPermutation3(t *testing.T) {
	assert.Equal(t, "", lexGreaterPermutation("baba", "bbaa"))
}

func Test_lexGreaterPermutation4(t *testing.T) {
	assert.Equal(t, "abdc", lexGreaterPermutation("abcd", "abcd"))
}

func Test_lexGreaterPermutation5(t *testing.T) {
	assert.Equal(t, "bcad", lexGreaterPermutation("abcd", "bbcd"))
}

func Test_lexGreaterPermutation6(t *testing.T) {
	assert.Equal(t, "", lexGreaterPermutation("kb", "ui"))
}

func Test_lexGreaterPermutation7(t *testing.T) {
	assert.Equal(t, "", lexGreaterPermutation("joo", "zju"))
}

func Test_lexGreaterPermutation8(t *testing.T) {
	assert.Equal(t, "baa", lexGreaterPermutation("aab", "abb"))
}

func Test_lexGreaterPermutation9(t *testing.T) {
	assert.Equal(t, "ba", lexGreaterPermutation("ba", "ab"))
}
