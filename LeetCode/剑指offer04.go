/* 在一个 n * m 的二维数组中，每一行都按照从左到右 非递减 的顺序排序，每一列都按照从上到下 非递减 的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。 */

package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	res := false
	row_index := len(matrix) - 1
	col_index := 0
	for row_index >= 0 {
		for col_index < len(matrix[row_index]) {
			fmt.Printf("matrix[%v][%v]: %v\n", row_index, col_index, matrix[row_index][col_index])
			if matrix[row_index][col_index] > target {
				row_index--
				break
			} else if matrix[row_index][col_index] < target {
				col_index++
				continue
			} else {
				res = true
				return res
			}
		}
		if col_index == len(matrix[0]) {
			break
		}
	}
	return res
}

func main() {
	/* matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	} */

	matrix := [][]int{
		{-5},
	}

	res := findNumberIn2DArray(matrix, -10)
	fmt.Printf("res: %v\n", res)
	return
}
