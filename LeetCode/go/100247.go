package main

import (
	"fmt"
	"sort"
)

func maximumHappinessSum(happiness []int, k int) int64 {
	maxHappinessSum := 0
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})
	for i := 0; i < k; i++ {
		if (happiness[i] - i) > 0 {
			maxHappinessSum += happiness[i] - i
		} else {
			return int64(maxHappinessSum)
		}
	}
	return int64(maxHappinessSum)
}

func main() {
	happiness := []int{1, 2, 3}
	k := 2
	fmt.Printf("maximumHappinessSum(happiness, k): %v\n", maximumHappinessSum(happiness, k))
}
