package dp

import (
	"LeetCode-Go/tst"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxValue1(t *testing.T) {
	actual := maxValue(tst.Str2Matrix("[[1,2,4],[3,4,3],[2,3,1]]"), 2)
	assert.Equal(t, 7, actual)
}

func Test_maxValue2(t *testing.T) {
	actual := maxValue(tst.Str2Matrix("[[1,2,4],[3,4,3],[2,3,10]]"), 2)
	assert.Equal(t, 10, actual)
}

func Test_maxValue3(t *testing.T) {
	actual := maxValue(tst.Str2Matrix("[[1,1,1],[2,2,2],[3,3,3],[4,4,4]]"), 3)
	assert.Equal(t, 9, actual)
}

func Test_maxValue4(t *testing.T) {
	actual := maxValue(tst.Str2Matrix("[[1,3,4],[2,4,1],[1,1,4],[3,5,1],[2,5,5]]"), 3)
	assert.Equal(t, 9, actual)
}
