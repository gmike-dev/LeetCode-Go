package n2997_minimum_number_of_operations_to_make_array_xor_equal_to_k

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinOperations1(t *testing.T) {
	actual := minOperations([]int{2, 1, 3, 4}, 1)
	expected := 2
	assert.Equal(t, expected, actual)
}

func TestMinOperations2(t *testing.T) {
	actual := minOperations([]int{2, 0, 2, 0}, 0)
	expected := 0
	assert.Equal(t, expected, actual)
}
