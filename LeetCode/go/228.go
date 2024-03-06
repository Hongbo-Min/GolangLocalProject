package main

import "strconv"

func summaryRanges(nums []int) []string {
	ans := []string{}
	for i, n := 0, len(nums); i < n; {
		left := i
		for i++; i < n && nums[i-1]+1 == nums[i]; i++ {
		}
		s := strconv.Itoa(nums[left])
		if left != i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		ans = append(ans, s)
	}
	return ans
}

func main() {
	resStr := summaryRanges([]int{0, 1, 2, 4, 5, 7})
	for _, res := range resStr {
		println(res)
	}
}
