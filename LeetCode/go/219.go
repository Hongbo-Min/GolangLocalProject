package main

import "sort"

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}
	numberAndIndexs := make(map[int][]int)
	for index, num := range nums {
		if _, ok := numberAndIndexs[num]; ok {
			numberAndIndexs[num] = append(numberAndIndexs[num], index)
		} else {
			indexs := []int{index}
			numberAndIndexs[num] = indexs
		}
	}
	for _, indexs := range numberAndIndexs {
		sort.Ints(indexs)
		for i := 1; i < len(indexs); i++ {
			if indexs[i]-indexs[i-1] <= k {
				return true
			}
		}
	}
	return false
}

func main() {
	print(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
}
