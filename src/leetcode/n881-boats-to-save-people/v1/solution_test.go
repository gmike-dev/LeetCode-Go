package n881_boats_to_save_people_v1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumRescueBoats1(t *testing.T) {
	actual := numRescueBoats([]int{1, 2}, 3)
	expect := 1
	assert.Equal(t, expect, actual)
}

func TestNumRescueBoats2(t *testing.T) {
	actual := numRescueBoats([]int{3, 2, 2, 1}, 3)
	expect := 3
	assert.Equal(t, expect, actual)
}

func TestNumRescueBoats3(t *testing.T) {
	actual := numRescueBoats([]int{3, 5, 3, 4}, 5)
	expect := 4
	assert.Equal(t, expect, actual)
}

func TestNumRescueBoats4(t *testing.T) {
	actual := numRescueBoats([]int{1}, 5)
	expect := 1
	assert.Equal(t, expect, actual)
}

func TestNumRescueBoats5(t *testing.T) {
	actual := numRescueBoats([]int{1, 2, 3}, 4)
	expect := 2
	assert.Equal(t, expect, actual)
}
