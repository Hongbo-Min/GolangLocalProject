package main

import "fmt"

func diagonalSum(mat [][]int) int {
	var result int
	if len(mat) == 0 {
		return result
	}

	var left, right = 0, len(mat) - 1
	for i := 0; i < len(mat); i++ {
		if left < len(mat) && right >= 0 {
			result += mat[i][left]
			if left != right {
				result += mat[i][right]
			}
			left++
			right--
		}
	}

	return result
}

func main() {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("result : ", diagonalSum(mat))
}
