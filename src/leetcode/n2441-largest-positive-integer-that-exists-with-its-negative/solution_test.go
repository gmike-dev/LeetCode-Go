package n2441_largest_positive_integer_that_exists_with_its_negative

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxK1(t *testing.T) {
	actual := findMaxK([]int{-1, 2, -3, 3})
	expected := 3
	assert.Equal(t, expected, actual)
}

func TestFindMaxK2(t *testing.T) {
	actual := findMaxK([]int{-1, 10, 6, 7, -7, 1})
	expected := 7
	assert.Equal(t, expected, actual)
}

func TestFindMaxK3(t *testing.T) {
	actual := findMaxK([]int{-10, 8, 6, 7, -2, -3})
	expected := -1
	assert.Equal(t, expected, actual)
}
