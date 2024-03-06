package main

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	numberAndIndex := make(map[int]int)
	for i, num := range nums {
		find := target - num
		if index, ok := numberAndIndex[find]; ok {
			if index != i {
				return []int{i, index}
			}
		} else {
			numberAndIndex[num] = i
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	for _, num := range twoSum(nums, target) {
		println(num)
	}
}
