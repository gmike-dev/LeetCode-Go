package tst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStr2List(t *testing.T) {
	assert.Equal(t, []int{}, Str2List("[]"))
	assert.Equal(t, []int{1, 2, 3}, Str2List("[1,2,3]"))
	assert.Equal(t, []int{-10}, Str2List("[-10]"))
}

func TestStr2Matrix(t *testing.T) {
	assert.Equal(t, [][]int{}, Str2Matrix("[]"))
	assert.Equal(t, [][]int{{}}, Str2Matrix("[[]]"))
	assert.Equal(t, [][]int{{}, {}}, Str2Matrix("[[],[]]"))
	assert.Equal(t, [][]int{{1, 2, 3}}, Str2Matrix("[[1,2,3]]"))
	assert.Equal(t, [][]int{{1, 2, 3}, {4}, {5, 6}}, Str2Matrix("[[1,2,3],[4],[5,6]]"))
	assert.Equal(t, [][]int{{1, 2, 3}, {4}, {}}, Str2Matrix("[[1,2,3],[4],[]]"))
}

func TestList2String(t *testing.T) {
	assert.Equal(t, "[]", List2Str([]int{}))
	assert.Equal(t, "[-10]", List2Str([]int{-10}))
	assert.Equal(t, "[1,2,3]", List2Str([]int{1, 2, 3}))
}

func TestMx2String(t *testing.T) {
	assert.Equal(t, "[]", Mx2Str([][]int{}))
	assert.Equal(t, "[[]]", Mx2Str([][]int{{}}))
	assert.Equal(t, "[[],[]]", Mx2Str([][]int{{}, {}}))
	assert.Equal(t, "[[1,2,3]]", Mx2Str([][]int{{1, 2, 3}}))
	assert.Equal(t, "[[1,2,3],[4],[5,6]]", Mx2Str([][]int{{1, 2, 3}, {4}, {5, 6}}))
	assert.Equal(t, "[[1,2,3],[4],[]]", Mx2Str([][]int{{1, 2, 3}, {4}, {}}))
}
