package bitmasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateParenthesis(t *testing.T) {
	assert.ElementsMatch(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, generateParenthesis(3))
	assert.ElementsMatch(t, []string{"()"}, generateParenthesis(1))
}
