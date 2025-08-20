package n2088_count_fertile_pyramids_in_a_land

import (
	"LeetCode-Go/tst"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countPyramids1(t *testing.T) {
	assert.Equal(t, 2, countPyramids(tst.Str2Matrix("[[0,1,1,0],[1,1,1,1]]")))
}
func Test_countPyramids2(t *testing.T) {
	assert.Equal(t, 2, countPyramids(tst.Str2Matrix("[[1,1,1],[1,1,1]]")))
}
func Test_countPyramids3(t *testing.T) {
	assert.Equal(t, 13, countPyramids(tst.Str2Matrix("[[1,1,1,1,0],[1,1,1,1,1],[1,1,1,1,1],[0,1,0,0,1]]")))
}
func Test_countPyramids4(t *testing.T) {
	assert.Equal(t, 10, countPyramids(tst.Str2Matrix("[[0,0,1,1,0,0],[0,1,1,1,1,0],[1,1,1,1,1,1]]")))
}
