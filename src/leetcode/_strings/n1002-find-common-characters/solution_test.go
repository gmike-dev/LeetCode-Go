package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommonChars1(t *testing.T) {
	assert.Equal(t, []string{"e", "l", "l"}, commonChars([]string{"bella", "label", "roller"}))
}

func TestCommonChars2(t *testing.T) {
	assert.Equal(t, []string{"c", "o"}, commonChars([]string{"cool", "lock", "cook"}))
}
