package main

import "fmt"

/*
1 2 3 0 0 0
2 5 6
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	var index, index1, index2 int
	index = m + n - 1
	index1 = m - 1
	index2 = n - 1

	for ; index >= 0 && index1 >= 0 && index2 >= 0; index-- {
		if nums1[index1] >= nums2[index2] {
			nums1[index] = nums1[index1]
			index1--
		} else {
			nums1[index] = nums2[index2]
			index2--
		}
	}

	for ; index1 >= 0; index1-- {
		nums1[index] = nums1[index1]
		index--
	}

	for ; index2 >= 0; index2-- {
		nums1[index] = nums2[index2]
		index--
	}

	for _, num := range nums1 {
		fmt.Printf("num: %v\n", num)
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
}
