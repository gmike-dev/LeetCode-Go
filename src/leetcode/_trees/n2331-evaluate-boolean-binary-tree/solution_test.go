package n2331_evaluate_boolean_binary_tree

import (
	. "LeetCode-Go/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvaluateTree1(t *testing.T) {
	assert.True(t, evaluateTree(StrToTree("2,1,3,null,null,0,1")))
}

func TestEvaluateTree2(t *testing.T) {
	assert.False(t, evaluateTree(StrToTree("0")))
}

func TestEvaluateTree3(t *testing.T) {
	assert.True(t, evaluateTree(StrToTree("2,0,1")))
}

func TestEvaluateTree4(t *testing.T) {
	assert.False(t, evaluateTree(StrToTree("3,0,1")))
}

func TestEvaluateTree5(t *testing.T) {
	assert.True(t, evaluateTree(StrToTree("3,1,1")))
}
