package v1

import (
	"LeetCode-Go/tst"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setZeroes1(t *testing.T) {
	mx := tst.Str2Matrix("[[1,1,1],[1,0,1],[1,1,1]]")
	setZeroes(mx)
	assert.Equal(t, "[[1,0,1],[0,0,0],[1,0,1]]", tst.Mx2Str(mx))
}

func Test_setZeroes2(t *testing.T) {
	mx := tst.Str2Matrix("[[0,1,2,0],[3,4,5,2],[1,3,1,5]]")
	setZeroes(mx)
	assert.Equal(t, "[[0,0,0,0],[0,4,5,0],[0,3,1,0]]", tst.Mx2Str(mx))
}
