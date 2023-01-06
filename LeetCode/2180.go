package main

import "fmt"

func countEven(num int) int {
	y, x := num/10, num%10
	res := y * 5
	ySum := 0
	for ; y > 0; y /= 10 {
		ySum += y % 10
	}
	if ySum%2 == 0 {
		res += x / 2
	} else {
		res += (x+1)/2 - 1
	}
	return res
}

func main() {
	fmt.Printf("countEven(30): %v\n", countEven(30))
	fmt.Printf("countEven(50): %v\n", countEven(50))
	fmt.Printf("countEven(51): %v\n", countEven(51))
}
