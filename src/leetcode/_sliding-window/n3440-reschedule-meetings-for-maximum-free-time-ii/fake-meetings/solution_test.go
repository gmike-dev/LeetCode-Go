package fake_meetings

import (
	"LeetCode-Go/tst"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxFreeTime(t *testing.T) {
	assert.Equal(t, 2, maxFreeTime(5, tst.Str2List("[1,3]"), tst.Str2List("[2,5]")))
	assert.Equal(t, 7, maxFreeTime(10, tst.Str2List("[0,7,9]"), tst.Str2List("[1,8,10]")))
	assert.Equal(t, 6, maxFreeTime(10, tst.Str2List("[0,3,7,9]"), tst.Str2List("[1,4,8,10]")))
	assert.Equal(t, 0, maxFreeTime(5, tst.Str2List("[0,1,2,3,4]"), tst.Str2List("[1,2,3,4,5]")))
}
