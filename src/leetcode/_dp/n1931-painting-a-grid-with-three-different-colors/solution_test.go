package n1931_painting_a_grid_with_three_different_colors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinTimeToReach_1(t *testing.T) {
	actual := colorTheGrid(1, 1)
	expected := 3
	assert.Equal(t, expected, actual)
}

func TestMinTimeToReach_2(t *testing.T) {
	actual := colorTheGrid(1, 2)
	expected := 6
	assert.Equal(t, expected, actual)
}

func TestMinTimeToReach_3(t *testing.T) {
	actual := colorTheGrid(5, 5)
	expected := 580986
	assert.Equal(t, expected, actual)
}

func TestMinTimeToReach_4(t *testing.T) {
	actual := colorTheGrid(5, 50)
	expected := 597561939
	assert.Equal(t, expected, actual)
}

func TestMinTimeToReach_5(t *testing.T) {
	actual := colorTheGrid(5, 1000)
	expected := 408208448
	assert.Equal(t, expected, actual)
}

func TestMinTimeToReach_6(t *testing.T) {
	actual := colorTheGrid(5, 10000)
	expected := 295998076
	assert.Equal(t, expected, actual)
}
