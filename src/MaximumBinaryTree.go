package src

import . "basic"

func constructMaximumBinaryTree(nums []int) *TreeNode {

	left, right := 0, len(nums)-1

	var head *TreeNode = new(TreeNode)
	head = construct(nums, left, right)
	return head
}

func construct(nums [] int, left int, right int) *TreeNode {

	if left > right {
		return nil
	}

	var head = new(TreeNode)
	if left == right {

		head.Val = nums[left]
		return head
	}

	index_max := findMax(nums, left, right)

	head.Val = nums[index_max]
	head.Left = construct(nums, left, index_max - 1)
	head.Right = construct(nums, index_max + 1, right)

	return head
}

func findMax(nums [] int, left int, right int) int {
	var max = left

	for ; left <= right; left ++ {
		if nums[left] > nums[max] {
			max = left
		}
	}
	return max
}
