package n1325_delete_leaves_with_a_given_value

import . "LeetCode-Go/collections"

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root.Left != nil {
		root.Left = removeLeafNodes(root.Left, target)
	}
	if root.Right != nil {
		root.Right = removeLeafNodes(root.Right, target)
	}
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}
