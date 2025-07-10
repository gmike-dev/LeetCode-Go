package n3439_reschedule_meetings_for_maximum_free_time_i

import (
	"LeetCode-Go/tst"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxFreeTime(t *testing.T) {
	assert.Equal(t, 2, maxFreeTime(5, 1, tst.Str2List("[1,3]"), tst.Str2List("[2,5]")))
	assert.Equal(t, 6, maxFreeTime(10, 1, tst.Str2List("[0,2,9]"), tst.Str2List("[1,4,10]")))
	assert.Equal(t, 0, maxFreeTime(5, 2, tst.Str2List("[0,1,2,3,4]"), tst.Str2List("[1,2,3,4,5]")))
	assert.Equal(t, 7, maxFreeTime(21, 1, tst.Str2List("[7,10,16]"), tst.Str2List("[10,14,18]")))
}
