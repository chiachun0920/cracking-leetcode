package leetcode

import (
	"fmt"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	tests := []struct {
		inputA   []int
		inputB   []int
		expected bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 1}, []int{1, 1, 2}, false},
	}

	for _, e := range tests {
		rootA := buildTree(e.inputA, 0, len(e.inputA))
		rootB := buildTree(e.inputB, 0, len(e.inputB))
		if isSameTree(rootA, rootB) != e.expected {
			t.Errorf("expected: %t", e.expected)
		}
	}
}

func Test_buildTree(t *testing.T) {
	node := buildTree([]int{1, 2, 3}, 0, 3)
	fmt.Println(node)
}

func buildTree(nums []int, i, n int) *TreeNode {
	var root TreeNode
	if i < n {
		root = TreeNode{Val: nums[i]}
		root.Left = buildTree(nums, 2*i+1, n)
		root.Right = buildTree(nums, 2*i+2, n)
	}
	return &root
}
