package main

import (
	"fmt"
	"sort"
)

func minimumBoxes(apple []int, capacity []int) int {
	appleCount := 0
	for _, c := range apple {
		appleCount += c
	}
	sort.Slice(capacity, func(i, j int) bool {
		return capacity[i] > capacity[j]
	})
	capacityCount := 0
	for i := 0; i < len(capacity); i++ {
		capacityCount += capacity[i]
		if capacityCount >= appleCount {
			return i + 1
		}
	}
	return 0
}

func main() {
	apple := []int{1, 3, 2}
	capacity := []int{4, 3, 1, 5, 2}
	fmt.Printf("minimumBoxes(apple, capacity): %v\n", minimumBoxes(apple, capacity))
}
