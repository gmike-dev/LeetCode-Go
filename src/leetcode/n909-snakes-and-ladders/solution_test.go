package n909_snakes_and_ladders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnakesAndLadders1(t *testing.T) {
	assert.Equal(t, 4, snakesAndLadders([][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1},
	}))
}

func TestSnakesAndLadders2(t *testing.T) {
	assert.Equal(t, 1, snakesAndLadders([][]int{
		{-1, -1},
		{-1, 3},
	}))
}

func TestSnakesAndLadders3(t *testing.T) {
	assert.Equal(t, 4, snakesAndLadders([][]int{
		{2, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1},
	}))
}

func TestSnakesAndLadders4(t *testing.T) {
	assert.Equal(t, -1, snakesAndLadders([][]int{
		{1, 1, -1},
		{1, 1, 1},
		{-1, 1, 1},
	}))
}
