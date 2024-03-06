package main

import "strings"

func wordPattern(pattern string, s string) bool {
	sList := strings.Split(s, " ")
	if (pattern == "" && s == "") || (len(pattern) != len(sList)) {
		return false
	}
	patternByte := []byte(pattern)
	patternMap := make(map[byte]string)
	sMap := make(map[string]byte)
	for index, b := range patternByte {
		if _, ok := patternMap[b]; !ok {
			if _, ok := sMap[sList[index]]; !ok {
				patternMap[b] = sList[index]
				sMap[sList[index]] = b
			} else {
				if b != sMap[sList[index]] {
					return false
				}
			}
		} else {
			if patternMap[b] != sList[index] {
				return false
			}
		}
	}
	return true
}

func main() {
	pattern := "abba"
	s := "dog cat cat dog"
	println(wordPattern(pattern, s))
}
