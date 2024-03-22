package main

import "fmt"

func BinarySearchOne(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}

// 二分查找 [left, right]
func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	res1 := BinarySearchOne(nums, target)
	fmt.Printf("res1: %v\n", res1)
}
