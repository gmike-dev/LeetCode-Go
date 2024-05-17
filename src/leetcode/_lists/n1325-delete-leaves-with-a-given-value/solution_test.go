package n1325_delete_leaves_with_a_given_value

import (
	. "LeetCode-Go/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveLeafNodes1(t *testing.T) {
	actual := removeLeafNodes(StrToTree("1,2,3,2,null,2,4"), 2)
	expect := "1,null,3,null,4"
	assert.Equal(t, expect, TreeToStr(actual))
}

func TestRemoveLeafNodes2(t *testing.T) {
	actual := removeLeafNodes(StrToTree("1,3,3,3,2"), 3)
	expect := "1,3,null,null,2"
	assert.Equal(t, expect, TreeToStr(actual))
}

func TestRemoveLeafNodes3(t *testing.T) {
	actual := removeLeafNodes(StrToTree("1,2,null,2,null,2"), 2)
	expect := "1"
	assert.Equal(t, expect, TreeToStr(actual))
}
