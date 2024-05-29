package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumSteps1(t *testing.T) {
	assert.Equal(t, 6, numSteps("1101"))
}

func TestNumSteps2(t *testing.T) {
	assert.Equal(t, 6, numSteps("1010"))
}

func TestNumSteps3(t *testing.T) {
	assert.Equal(t, 3, numSteps("1000"))
}

func TestNumSteps4(t *testing.T) {
	assert.Equal(t, 7, numSteps("1001"))
}

func TestNumSteps5(t *testing.T) {
	assert.Equal(t, 1, numSteps("10"))
}

func TestNumSteps6(t *testing.T) {
	assert.Equal(t, 0, numSteps("1"))
}

func TestNumSteps7(t *testing.T) {
	assert.Equal(t, 5, numSteps("1100"))
}

func TestNumSteps8(t *testing.T) {
	assert.Equal(t, 13, numSteps("110101101"))
}
