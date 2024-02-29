package main

import "fmt"

func isSubsequence(s string, t string) bool {
	var sIdx, tIdx int
	for sIdx, tIdx = 0, 0; sIdx < len(s) && tIdx < len(t); {
		if s[sIdx] == t[tIdx] {
			sIdx++
			tIdx++
		} else {
			tIdx++
		}
	}
	return sIdx == len(s)
}

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Printf("isSubsequence(s, t): %v\n", isSubsequence(s, t))
}
