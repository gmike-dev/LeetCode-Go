package n3495_minimum_operations_to_make_array_elements_zero

import (
	"LeetCode-Go/tst"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minOperations1(t *testing.T) {
	assert.Equal(t, int64(3), minOperations(tst.Str2Matrix("[[1,2],[2,4]]")))
}

func Test_minOperations2(t *testing.T) {
	assert.Equal(t, int64(4), minOperations(tst.Str2Matrix("[[2,6]]")))
}

func Test_minOperations3(t *testing.T) {
	assert.Equal(t, int64(3), minOperations(tst.Str2Matrix("[[32,33]]")))
}

func Test_minOperations4(t *testing.T) {
	assert.Equal(t, int64(7), minOperations(tst.Str2Matrix("[[1,8]]")))
}

func Test_minOperations5(t *testing.T) {
	assert.Equal(t, int64(44), minOperations(tst.Str2Matrix("[[14,22],[3,9],[2,18],[20,23]]")))
}
