package main

import "fmt"

func reverseWords(s string) string {
	wordAndIndex := make(map[int]string)
	maxIndex := 0
	temp := ""
	for _, c := range s {
		if string(c) == " " {
			if len(temp) > 0 {
				wordAndIndex[maxIndex] = temp
				maxIndex = maxIndex + 1
				temp = ""
			}
		} else {
			temp += string(c)
		}
	}
	if len(temp) != 0 {
		maxIndex = maxIndex + 1
		wordAndIndex[maxIndex] = temp
		temp = ""
	}
	result := ""
	for i := maxIndex; i >= 0; i-- {
		if v, ok := wordAndIndex[i]; ok {
			result += v + " "
		}
	}
	if len(result) > 0 {
		return result[:len(result)-1]
	} else {
		return result
	}
}

func main() {
	s := "the sky is blue"
	res := reverseWords(s)
	fmt.Printf("res: %v\n", res)
}
