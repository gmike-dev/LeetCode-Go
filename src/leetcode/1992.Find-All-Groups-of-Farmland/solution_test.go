package leetcode_1992

import (
	"reflect"
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
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}
}
func TestFindFarmland2(t *testing.T) {
	actual := findFarmland([][]int{
		{1, 1},
		{1, 1},
	})
	expected := [][]int{
		{0, 0, 1, 1},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}
}
func TestFindFarmland3(t *testing.T) {
	actual := findFarmland([][]int{
		{0},
	})
	expected := [][]int{}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}
}
