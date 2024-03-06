package main

func getAfterNumber(number int) int {
	res := 0
	for number > 0 {
		temp := number % 10
		res += (temp * temp)
		number /= 10
	}
	return res
}

func isHappy(n int) bool {
	beforeAndAfter := make(map[int]int)
	for n != 1 {
		after := getAfterNumber(n)
		if _, ok := beforeAndAfter[after]; ok {
			return false
		} else {
			beforeAndAfter[n] = after
			n = after
		}
	}
	return true
}

func main() {
	println(isHappy(19))
}
