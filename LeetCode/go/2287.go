package main

import (
	"fmt"
	"math"
)

func rearrangeCharacters(s string, target string) (res int) {
	targetMap := make(map[string]int)
	sMap := make(map[string]int)
	for _, value := range target {
		targetMap[string(value)]++
	}

	for _, value := range s {
		sMap[string(value)]++
	}

	res = len(s)
	for k, v := range targetMap {
		if v > 0 {
			res = int(math.Min(float64(res), float64(sMap[k])/float64(v)))
			if res == 0 {
				return
			}
		}
	}

	return
}

func main() {
	s := "ilovecodingonleetcode"
	target := "code"
	res := rearrangeCharacters(s, target)
	fmt.Printf("res: %v\n", res)
}
