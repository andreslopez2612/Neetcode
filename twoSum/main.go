package main

import "fmt"

func main() {
	s, t := []int{3, 4, 5, 7}, 10

	f := twoSum(s, t)

	fmt.Println(f)
}

func twoSum(nums []int, target int) []int {

	result := []int{}

	check := make(map[int]int)
	for i, v := range nums {
		diference := target - v
		if prev, ok := check[diference]; ok {
			result = append(result, prev, i)
			return result
		}
		check[v] = i
	}

	return nil
}
