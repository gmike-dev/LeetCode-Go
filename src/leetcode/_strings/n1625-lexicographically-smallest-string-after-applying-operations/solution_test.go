package n1625_lexicographically_smallest_string_after_applying_operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findLexSmallestString1(t *testing.T) {
	assert.Equal(t, "2050", findLexSmallestString("5525", 9, 2))
}

func Test_findLexSmallestString2(t *testing.T) {
	assert.Equal(t, "24", findLexSmallestString("74", 5, 1))
}

func Test_findLexSmallestString3(t *testing.T) {
	assert.Equal(t, "0011", findLexSmallestString("0011", 4, 2))
}

func Test_findLexSmallestString4(t *testing.T) {
	assert.Equal(t, "00553311", findLexSmallestString("43987654", 7, 3))
}
