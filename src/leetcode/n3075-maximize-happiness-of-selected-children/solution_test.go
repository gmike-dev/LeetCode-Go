package n3075_maximize_happiness_of_selected_children

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximumHappinessSum1(t *testing.T) {
	actual := maximumHappinessSum([]int{1, 2, 3}, 2)
	expect := int64(4)
	assert.Equal(t, expect, actual)
}

func TestMaximumHappinessSum2(t *testing.T) {
	actual := maximumHappinessSum([]int{1, 1, 1, 1}, 2)
	expect := int64(1)
	assert.Equal(t, expect, actual)
}

func TestMaximumHappinessSum3(t *testing.T) {
	actual := maximumHappinessSum([]int{2, 3, 4, 5}, 1)
	expect := int64(5)
	assert.Equal(t, expect, actual)
}
