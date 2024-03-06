package main

func isIsomorphic(s string, t string) bool {
	sByte := []byte(s)
	tByte := []byte(t)
	if (s == "" && t == "") || (len(sByte) != len(tByte)) {
		return false
	}
	tMap := make(map[byte]byte)
	sMap := make(map[byte]byte)
	for index, b := range sByte {
		if _, ok := sMap[b]; !ok {
			if _, ok = tMap[tByte[index]]; !ok {
				sMap[b] = tByte[index]
				tMap[tByte[index]] = b
			} else {
				if tMap[tByte[index]] != b {
					return false
				}
			}
		} else {
			if sMap[b] != tByte[index] {
				return false
			}
		}
	}
	return true
}

func main() {
	s := "egg"
	t := "add"
	println(isIsomorphic(s, t))
}
