package src

import . "basic"

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	if t1 == nil && t2 == nil {
		return nil
	}

	if t2 == nil {
		return t1
	}

	if t1 == nil {
		return t2
	}

	head := new(TreeNode)
	head.Val = (*t1).Val + (*t2).Val

	head.Left = mergeTrees((*t1).Left, (*t2).Left)
	head.Right = mergeTrees((*t1).Right, (*t2).Right)


	return head
}