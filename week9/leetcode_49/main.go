package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)
	fmt.Printf("result: %v\n", result)
}

func groupAnagrams(strs []string) [][]string {
	hashMap := map[string][]string{}

	for _, str := range strs {
		b := []byte(str)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		sorted := string(b)
		hashMap[sorted] = append(hashMap[sorted], str)

	}

	ans := make([][]string, 0, len(hashMap))

	for _, array := range hashMap {
		ans = append(ans, array)
	}

	return ans
}
