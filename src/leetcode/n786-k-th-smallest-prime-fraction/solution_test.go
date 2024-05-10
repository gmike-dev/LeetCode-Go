package n786_k_th_smallest_prime_fraction

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKthSmallestPrimeFractionV1_1(t *testing.T) {
	actual := kthSmallestPrimeFractionV1([]int{1, 2, 3, 5}, 3)
	expect := []int{2, 5}
	assert.Equal(t, expect, actual)
}

func TestKthSmallestPrimeFractionV1_2(t *testing.T) {
	actual := kthSmallestPrimeFractionV1([]int{1, 7}, 1)
	expect := []int{1, 7}
	assert.Equal(t, expect, actual)
}

func TestKthSmallestPrimeFractionV2_1(t *testing.T) {
	actual := kthSmallestPrimeFractionV2([]int{1, 2, 3, 5}, 3)
	expect := []int{2, 5}
	assert.Equal(t, expect, actual)
}

func TestKthSmallestPrimeFractionV2_2(t *testing.T) {
	actual := kthSmallestPrimeFractionV2([]int{1, 7}, 1)
	expect := []int{1, 7}
	assert.Equal(t, expect, actual)
}
