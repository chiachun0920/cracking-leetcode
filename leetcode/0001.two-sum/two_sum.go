package leetcode

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for idx, num := range nums {
		numMap[num] = idx
	}

	for idx, num := range nums {
		value, ok := numMap[target-num]
		if !ok || value == idx {
			continue
		}
		return []int{idx, value}
	}
	return []int{}
}
