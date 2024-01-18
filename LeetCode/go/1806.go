package main

import (
	"fmt"
)

func reinitializePermutation(n int) (res int) {
	base := make([]int, n)
	perm := make([]int, n)
	for i := 0; i < n; i++ {
		perm[i] = i
		base[i] = i
	}

	for {
		res++
		arr := make([]int, n)
		for i := range arr {
			if i%2 == 0 {
				arr[i] = perm[i/2]
			} else {
				arr[i] = perm[n/2+i/2]
			}
		}
		perm = arr
		if Compare(perm, base) {
			return
		}
	}
}

func Compare(a, b []int) bool {
	for index, value := range a {
		if value != b[index] {
			return false
		}
	}
	return true
}

func main() {
	res := reinitializePermutation(6)
	fmt.Printf("res: %v\n", res)
}
