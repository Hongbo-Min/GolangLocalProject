package main

import "fmt"

func circularGameLosers(n int, k int) []int {
	result := make([]int, 0)
	mark := make([]bool, n+1)

	var index = 1
	var count = 1
	for {
		if mark[index] {
			break
		} else {
			mark[index] = true // 1 4 10(0)
			index += (count * k)
			if index > n {
				if index%n == 0 {
					index = n
				} else {
					index = index % n
				}
			}
			count = count + 1
		}
	}

	for index, ok := range mark {
		if index == 0 {
			continue
		}
		if !ok {
			result = append(result, index)
		}
	}
	return result
}

func main() {
	var n int
	var k int
	n = 5
	k = 3
	for _, result := range circularGameLosers(n, k) {
		fmt.Println("result : ", result)
	}
}
