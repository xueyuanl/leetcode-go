package tree

import (
	. "basic"
)

func trimBST(root *TreeNode, L int, R int) *TreeNode {

	if L > R {
		return nil
	}
	if root == nil {
		return nil
	}

	if root.Val > R {
		return trimBST(root.Left, L, R)
	}

	if root.Val < L {
		return trimBST(root.Right, L, R)
	}

	if root.Val >= L && root.Val <= R {
		root.Left = trimBST(root.Left, L, root.Val - 1)
		root.Right = trimBST(root.Right, root.Val + 1, R)

	}
	return root
}