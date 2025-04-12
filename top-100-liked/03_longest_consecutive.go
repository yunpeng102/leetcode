package top_100_liked

import "sort"

func longestConsecutive(nums []int) int {
	sort.Ints(removeDuplication(nums))

	maxLength := 0
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j]-nums[i] == j-i {
			continue
		} else {
			if j-i > maxLength {
				maxLength = j - i
			}
			i = j
		}
	}

	return maxLength
}

func removeDuplication(nums []int) []int {
	m := map[int]bool{}

	results := []int{}
	for _, v := range nums {
		if !m[v] {
			m[v] = true
			results = append(results, v)
		}
	}
	return results
}
