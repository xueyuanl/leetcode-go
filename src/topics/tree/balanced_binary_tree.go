package tree

import (
	. "basic"
)

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return isBalanced(root.Left) && isBalanced(root.Right) && abs(depthOfTree(root.Left), depthOfTree(root.Right)) <2

}

func depthOfTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(depthOfTree(root.Left), depthOfTree(root.Right)) + 1

}

func abs(a int, b int) int {
	if a > b {
		return a - b
	}else {
		return b - a
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}