package main

import "fmt"

func solution(s string) (res int) {
	temp := 0
	for index, value := range s {
		if value == 'X' && index >= temp {
			fmt.Println("index: ", index, " temp: ", temp)
			res++
			temp = index + 3
		}
	}
	return
}

func main() {
	s := "OXOX"
	res := solution(s)
	fmt.Println(res)
}
