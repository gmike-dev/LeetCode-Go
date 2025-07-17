package n3202_find_the_maximum_length_of_valid_subsequence_ii

import (
	"LeetCode-Go/tst"
	"math/rand/v2"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_maximumLength(t *testing.T) {
	assert.Equal(t, 5, maximumLength(tst.Str2List("[1,2,3,4,5]"), 2))
	assert.Equal(t, 4, maximumLength(tst.Str2List("[1,4,2,3,1,4]"), 3))
	assert.Equal(t, 50, maximumLength(tst.Str2List("[4,5,10,2,4,9,3,10,1,1,1,4,8,6,5,8,2,4,9,8,2,2,1,3,4,6,6,6,7,5,9,4,2,5,7,2,8,3,9,2,6,9,4,10,1,4,1,6,7,6]"), 1))
}

func Test_maximumLength_Bench(t *testing.T) {
	a := make([]int, 1000)
	for i := range a {
		a[i] = rand.N(1000000)
	}

	done := make(chan int)
	go func() {
		defer close(done)
		maximumLength(a, 1000)
	}()
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Errorf("timeout")
	}
}
