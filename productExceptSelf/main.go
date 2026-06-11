package main

import "fmt"

func main() {

	val := []int{1, 2, 4, 6}
	//Output: [48,24,12,8]
	f := productExceptSelf(val)

	fmt.Println(f)

}

func productExceptSelf(nums []int) []int {
	// Pre-asignamos el slice con la longitud exacta y valores inicializados en 0
	length := len(nums)
	result := make([]int, length)

	prefix := 1
	for i := 0; i < length; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := length - 1; i >= 0; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}

	return result
}
