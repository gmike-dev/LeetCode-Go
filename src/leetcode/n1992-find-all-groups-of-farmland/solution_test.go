package n1992_find_all_groups_of_farmland

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindFarmland1(t *testing.T) {
	actual := findFarmland([][]int{
		{1, 0, 0},
		{0, 1, 1},
		{0, 1, 1},
	})
	expected := [][]int{
		{0, 0, 0, 0},
		{1, 1, 2, 2},
	}
	assert.Equal(t, expected, actual)
}
func TestFindFarmland2(t *testing.T) {
	actual := findFarmland([][]int{
		{1, 1},
		{1, 1},
	})
	expected := [][]int{
		{0, 0, 1, 1},
	}
	assert.Equal(t, expected, actual)
}
func TestFindFarmland3(t *testing.T) {
	actual := findFarmland([][]int{
		{0},
	})
	expected := [][]int{}
	assert.Equal(t, expected, actual)
}
