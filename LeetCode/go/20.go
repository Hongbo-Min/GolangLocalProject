package main

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, b := range s {
		runeB := rune(b)
		switch runeB {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		case ')', ']', '}':
			if len(stack) != 0 && stack[len(stack)-1] == runeB {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	println(isValid("()"))
}
