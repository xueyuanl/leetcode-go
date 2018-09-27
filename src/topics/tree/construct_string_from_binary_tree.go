package tree

import (
	. "basic"
	"strconv"
)

func tree2str (t *TreeNode) string {
	if t == nil {
		return ""
	}

	str := strconv.Itoa(t.Val)

	leftStr :=tree2str(t.Left)
	rightStr := tree2str(t.Right)

	if leftStr != "" && rightStr != "" {
		str += string('(' ) + leftStr + string(')')  + string('(' ) + rightStr + string(')')
	}

	if leftStr != "" && rightStr == "" {
		str += string('(' ) + leftStr  + string(')')
	}

	if leftStr == "" && rightStr != "" {
		str += string('(' ) +  string(')')  + string('(' ) + rightStr + string(')')
	}

	return str
}