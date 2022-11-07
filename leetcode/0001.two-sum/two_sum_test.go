package leetcode

import "testing"

func Test_twoSum(t *testing.T) {
	var tests = []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}

	for _, e := range tests {
		result := twoSum(e.nums, e.target)
		if !isSliceEqual(result, e.expected) {
			t.Fatalf("Expected %v, but got %v", e.expected, result)
		}
	}
}

func isSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, num := range a {
		if num != b[idx] {
			return false
		}
	}

	return true
}
