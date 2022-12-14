package leetcode

import (
	"crackingleetcode/ds"
)

type TreeNode = ds.TreeNode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)
}
