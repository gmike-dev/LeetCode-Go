package n310_minimum_height_trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMinHeightTrees1(t *testing.T) {
	actual := findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}})
	expect := []int{1}
	assert.Equal(t, expect, actual)
}
func TestFindMinHeightTrees2(t *testing.T) {
	actual := findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}})
	expect := []int{3, 4}
	assert.Equal(t, expect, actual)
}
func TestFindMinHeightTrees3(t *testing.T) {
	actual := findMinHeightTrees(5, [][]int{{1, 4}, {4, 2}, {4, 3}, {1, 0}})
	expect := []int{1, 4}
	assert.Equal(t, expect, actual)
}
func TestFindMinHeightTrees4(t *testing.T) {
	actual := findMinHeightTrees(1, [][]int{})
	expect := []int{0}
	assert.Equal(t, expect, actual)
}
func TestFindMinHeightTrees5(t *testing.T) {
	actual := findMinHeightTrees(3, [][]int{{0, 1}, {0, 2}})
	expect := []int{0}
	assert.Equal(t, expect, actual)
}
func TestFindMinHeightTrees6(t *testing.T) {
	actual := findMinHeightTrees(7, [][]int{{0, 1}, {1, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 6}})
	expect := []int{1, 2}
	assert.Equal(t, expect, actual)
}
