package tree

import . "basic"

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sumOfLeft := sumOfLeft(root.Left)
	sumOfRight := sumOfRight(root.Right)

	return sumOfLeft + sumOfRight

}

func sumOfLeft(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if isLeaf(root) {
		return root.Val
	}

	sumOfLeft := sumOfLeft(root.Left)
	sumOfRight := sumOfRight(root.Right)

	return sumOfLeft + sumOfRight

}

func sumOfRight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sumOfLeft := sumOfLeft(root.Left)
	sumOfRight := sumOfRight(root.Right)

	return sumOfLeft + sumOfRight

}

func isLeaf(root *TreeNode) bool {
	if root.Left ==nil && root.Right == nil {
		return true
	}
	return false
}
