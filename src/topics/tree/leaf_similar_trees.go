package tree

import . "basic"

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var res bool
	leaves1 := make([]int, 10)
	leaves2 := make([]int, 10)

	appendleaves(&leaves1, root1)
	appendleaves(&leaves2, root2)

	res = intSliceEqual(leaves1, leaves2)
	return res
}

func appendleaves(leaves *[]int, root *TreeNode) {
	if root.Left == nil && root.Right == nil {
		*leaves = append(*leaves, root.Val)
		return;
	}
	if root.Left != nil {
		appendleaves(leaves, root.Left)
	}
	if root.Right != nil {
		appendleaves(leaves, root.Right)
	}

}

func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}