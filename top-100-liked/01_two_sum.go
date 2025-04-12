package top_100_liked

import "sort"

func twoSum(nums []int, target int) []int {
	x := 0
	y := len(nums) - 1

	copySlice := make([]int, len(nums))
	copy(copySlice, nums)

	sort.Ints(copySlice)

	for {
		if copySlice[x]+copySlice[y] == target {
			break
		} else if copySlice[x]+copySlice[y] >= target {
			y--
		} else {
			x--
		}
	}

	result := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] == copySlice[x] || nums[i] == copySlice[y] {
			result = append(result, i)
		}
		if len(result) == 2 {
			break
		}
	}

	return result
}
