package n289_game_of_life

import (
	"LeetCode-Go/tst"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_gameOfLife1(t *testing.T) {
	board := tst.Str2Matrix("[[0,1,0],[0,0,1],[1,1,1],[0,0,0]]")
	gameOfLife(board)
	assert.Equal(t, "[[0,0,0],[1,0,1],[0,1,1],[0,1,0]]", tst.Mx2Str(board))
}

func Test_gameOfLife2(t *testing.T) {
	board := tst.Str2Matrix("[[1,1],[1,0]]")
	gameOfLife(board)
	assert.Equal(t, "[[1,1],[1,1]]", tst.Mx2Str(board))
}
