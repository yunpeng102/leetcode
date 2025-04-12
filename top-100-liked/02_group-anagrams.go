package top_100_liked

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}

	for _, s := range strs {
		ss := sortString(s)
		m[ss] = append(m[ss], s)
	}

	result := [][]string{}
	for _, ss := range m {
		result = append(result, ss)
	}
	return result
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
