package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRotateSteps1(t *testing.T) {
	assert.Equal(t, 3, numDistinct("rabbbit", "rabbit"))
}

func TestFindRotateSteps2(t *testing.T) {
	assert.Equal(t, 5, numDistinct("babgbag", "bag"))
}
