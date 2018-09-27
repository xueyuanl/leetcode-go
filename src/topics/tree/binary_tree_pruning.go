package tree

import . "basic"

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	pruneTree(root.Left)
	pruneTree(root.Right)

	if is0sub(root.Left) {
		root.Left = nil
	}
	if is0sub(root.Right) {
		root.Right = nil
	}
	return root
}

func is0sub(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Val == 1 {
		return false
	}
	if root.Val == 0 && isLeaf(root) {
		return true
	}

	return is0sub(root.Left) && is0sub(root.Right)

}

func isLeaf(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	return false
}
