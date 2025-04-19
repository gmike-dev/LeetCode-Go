package n2563_count_the_number_of_fair_pairs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountFairPairs1(t *testing.T) {
	actual := countFairPairs([]int{0, 1, 7, 4, 4, 5}, 3, 6)
	expected := int64(6)
	assert.Equal(t, expected, actual)
}

func TestCountFairPairs2(t *testing.T) {
	actual := countFairPairs([]int{1, 7, 9, 2, 5}, 11, 11)
	expected := int64(1)
	assert.Equal(t, expected, actual)
}
