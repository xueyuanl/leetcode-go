package tree

import . "basic"

func searchBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return nil
	}

	if val == root.Val {
		return root
	}

	if val < root.Val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}

}
