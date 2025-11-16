package n3748_count_stable_subarrays

import (
	"LeetCode-Go/tst"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countStableSubarrays(t *testing.T) {
	assert.Equal(t, "[2,3,4]", tst.List2Str(countStableSubarrays(tst.Str2List("[3,1,2]"), tst.Str2Matrix("[[0,1],[1,2],[0,2]]"))))
	assert.Equal(t, "[3,1]", tst.List2Str(countStableSubarrays(tst.Str2List("[2,2]"), tst.Str2Matrix("[[0,1],[0,0]]"))))
	assert.Equal(t, "[15,10,6,3,1]", tst.List2Str(countStableSubarrays(tst.Str2List("[1,2,3,4,5]"), tst.Str2Matrix("[[0,4],[1,4],[2,4],[3,4],[4,4]]"))))
	assert.Equal(t, "[1]", tst.List2Str(countStableSubarrays(tst.Str2List("[8,12]"), tst.Str2Matrix("[[1,1]]"))))
}
