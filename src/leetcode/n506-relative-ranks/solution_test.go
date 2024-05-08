package n506_relative_ranks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindRelativeRanks1(t *testing.T) {
	actual := findRelativeRanks([]int{5, 4, 3, 2, 1})
	expect := []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}
	assert.Equal(t, expect, actual)
}

func TestFindRelativeRanks2(t *testing.T) {
	actual := findRelativeRanks([]int{10, 3, 8, 9, 4})
	expect := []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}
	assert.Equal(t, expect, actual)
}
