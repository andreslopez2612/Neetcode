package main

import "fmt"

func main() {

	val := []int{1, 2, 3}
	f := hasDuplicate(val)

	fmt.Println(f)
}

func hasDuplicate(nums []int) bool {
	check := make(map[int]bool)

	for _, v := range nums {
		if check[v] {
			return true
		}

		check[v] = true
	}

	return false
}
