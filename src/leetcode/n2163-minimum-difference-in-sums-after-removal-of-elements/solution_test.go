package n2163_minimum_difference_in_sums_after_removal_of_elements

import (
	"LeetCode-Go/tst"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumDifference(t *testing.T) {
	assert.Equal(t, int64(-1), minimumDifference(tst.Str2List("[3,1,2]")))
	assert.Equal(t, int64(1), minimumDifference(tst.Str2List("[7,9,5,8,1,3]")))
}
