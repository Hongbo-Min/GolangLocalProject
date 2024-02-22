package main

import "fmt"

func deleteWord(strs []string) {
	var slow, fast int
	for slow, fast = 0, 1; slow < len(strs) && fast < len(strs); {
		if strs[slow] == strs[fast] {
			fast = fast + 1
		} else {
			slow = slow + 1
			strs[slow] = strs[fast]
			fast = fast + 1
		}
	}
	for i := slow + 1; i < len(strs); i++ {
		strs[i] = ""
	}
}

func main() {
	strs := []string{"hello", "hello", "hi", "hi"}
	deleteWord(strs)
	fmt.Printf("strs: %v\n", strs)
	fmt.Printf("len(strs): %v\n", len(strs))
}
