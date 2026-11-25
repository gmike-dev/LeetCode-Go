package manually_impl_binary_search

import (
	"LeetCode-Go/tst"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_smallestUniqueSubarray(t *testing.T) {
	assert.Equal(t, 3, smallestUniqueSubarray(tst.Str2List("[3,3,3]")))
	assert.Equal(t, 1, smallestUniqueSubarray(tst.Str2List("[2,1,2,3,3]")))
	assert.Equal(t, 2, smallestUniqueSubarray(tst.Str2List("[1,1,2,2,1]")))
}
