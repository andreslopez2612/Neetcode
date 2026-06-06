package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	s := []string{"act", "pots", "tops", "cat", "stop", "hat"}

	f := groupAnagrams(s)

	fmt.Println(f)
}

func groupAnagrams(strs []string) [][]string {
	group := make(map[string][]string)

	for _, v := range strs {
		stringHead := order(v)
		group[stringHead] = append(group[stringHead], v)
	}

	result := [][]string{}

	for _, v := range group {
		result = append(result, v)
	}

	return result
}

func order(s string) string {
	res := strings.Split(s, "")
	slices.Sort(res)
	return strings.Join(res, "")
}
