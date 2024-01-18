package main

import "fmt"

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) (res []int) {
	get := func(nums []int) (s [101]int) {
		for _, value := range nums {
			s[value] = 1
		}
		return
	}

	s1, s2, s3 := get(nums1), get(nums2), get(nums3)
	for i := 1; i <= 100; i++ {
		if s1[i]+s2[i]+s3[i] > 1 {
			res = append(res, i)
		}
	}
	return
}

func main() {
	nums1 := []int{1, 1, 3, 2}
	nums2 := []int{2, 3}
	nums3 := []int{3}

	res := twoOutOfThree(nums1, nums2, nums3)
	for _, value := range res {
		fmt.Println("value: ", value)
	}
	return
}
