package main

import "fmt"

func rotate(nums []int, count int) []int {
	len := len(nums)
	count = count % len
	result := make([]int, len)
	for i := 0; i < len; i++ {
		result[(i+count)%len] = nums[i]
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("nums: %v\n", nums)
	count := 2
	nums = rotate(nums, count)
	fmt.Printf("nums: %v\n", nums)
}
