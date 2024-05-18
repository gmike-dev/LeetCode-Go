package n979_distribute_coins_in_binary_tree

import . "LeetCode-Go/collections"

func distributeCoins(root *TreeNode) int {
	moves := 0
	getCoins(root, &moves)
	return moves
}

func getCoins(node *TreeNode, moves *int) int {
	if node == nil {
		return 0
	}
	coins := node.Val + getCoins(node.Left, moves) + getCoins(node.Right, moves)
	extra := coins - 1
	*moves += abs(extra)
	return extra
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
