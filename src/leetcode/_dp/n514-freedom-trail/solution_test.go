package n514_freedom_trail

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindRotateSteps1(t *testing.T) {
	actual := findRotateSteps("godding", "gd")
	expect := 4
	assert.Equal(t, expect, actual)
}

func TestFindRotateSteps2(t *testing.T) {
	actual := findRotateSteps("godding", "godding")
	expect := 13
	assert.Equal(t, expect, actual)
}
