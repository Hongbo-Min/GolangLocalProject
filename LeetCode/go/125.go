package main

import "fmt"

func isPalindrome(s string) bool {
	convertStr := ""
	for _, c := range s {
		if (string(c) >= "A" && string(c) <= "Z") ||
			(string(c) >= "a" && string(c) <= "z") ||
			(string(c) >= "0" && string(c) <= "9") {
			if string(c) >= "A" && string(c) <= "Z" {
				convertStr += string(c + 32)
			} else {
				convertStr += string(c)
			}
		}
	}
	for i, j := 0, len(convertStr)-1; i < j; i, j = i+1, j-1 {
		if convertStr[i] != convertStr[j] {
			return false
		}
	}
	return true
}

func main() {
	s := "0P"
	fmt.Printf("isPalindrome(s): %v\n", isPalindrome(s))
}
