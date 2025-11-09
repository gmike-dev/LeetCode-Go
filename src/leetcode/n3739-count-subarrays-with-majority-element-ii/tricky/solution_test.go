package tricky

import (
	"LeetCode-Go/tst"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countMajoritySubarrays(t *testing.T) {
	assert.Equal(t, int64(5), countMajoritySubarrays(tst.Str2List("[1,2,2,3]"), 2))
	assert.Equal(t, int64(10), countMajoritySubarrays(tst.Str2List("[1,1,1,1]"), 1))
	assert.Equal(t, int64(0), countMajoritySubarrays(tst.Str2List("[1,2,3]"), 4))
}
