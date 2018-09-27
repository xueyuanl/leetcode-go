package tree

import . "basic"

func increasingBST(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	leftHead := increasingBST(root.Left)

	rightHead := increasingBST(root.Right)

	root.Right = rightHead
	root.Left = nil // a critical step

	if leftHead != nil {
		leftEnd := getEnd(leftHead)
		leftEnd.Right = root

		return leftHead
	}else {
		return root
	}

}

func getEnd(root *TreeNode) *TreeNode {

	if root.Right == nil {
		return root
	}

	return getEnd(root.Right)
}