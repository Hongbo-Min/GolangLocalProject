package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sByte := []byte(s)
	tByte := []byte(t)
	sMap := make(map[byte]int)
	for _, b := range sByte {
		sMap[b]++
	}
	for _, b := range tByte {
		if count, ok := sMap[b]; ok {
			if count < 1 {
				return false
			}
			sMap[b]--
		} else {
			return false
		}
	}
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	println(isAnagram(s, t))
}
