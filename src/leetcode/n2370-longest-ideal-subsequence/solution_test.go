package n2370_longest_ideal_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestIdealString1(t *testing.T) {
	actual := longestIdealString("acfgbd", 2)
	expected := 4
	assert.Equal(t, expected, actual)
}

func TestLongestIdealString2(t *testing.T) {
	actual := longestIdealString("abcd", 3)
	expected := 4
	assert.Equal(t, expected, actual)
}

func TestLongestIdealString3(t *testing.T) {
	actual := longestIdealString("lkpkxcigcs", 6)
	expected := 7
	assert.Equal(t, expected, actual)
}
