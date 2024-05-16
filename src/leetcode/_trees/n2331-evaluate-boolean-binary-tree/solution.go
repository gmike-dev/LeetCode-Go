package n2331_evaluate_boolean_binary_tree

import . "LeetCode-Go/collections"

func evaluateTree(root *TreeNode) bool {
	switch root.Val {
	case 2:
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	case 3:
		return evaluateTree(root.Left) && evaluateTree(root.Right)
	default:
		return root.Val == 1
	}
}
