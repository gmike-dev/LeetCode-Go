package n861_score_after_flipping_matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatrixScore1(t *testing.T) {
	actual := matrixScore([][]int{
		{0, 0, 1, 1},
		{1, 0, 1, 0},
		{1, 1, 0, 0},
	})
	expected := 39
	assert.Equal(t, expected, actual)
}

func TestMatrixScore2(t *testing.T) {
	actual := matrixScore([][]int{
		{0},
	})
	expected := 1
	assert.Equal(t, expected, actual)
}
