package main

import "fmt"

const (
	x = iota
	_
	y
	z = "zz"
	// k
	p = iota
)

type Flags int

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

const (
	Apple, Banana = iota + 1, iota + 2
	Cherimoya, Durian
	Elderberry, Fig
)

const (
	i = iota
	j = 3.14
	k = iota
	l
)

func main() {
	fmt.Println(x, y, z, k, p)
	fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)
	fmt.Println(Apple, Banana, Cherimoya, Durian, Elderberry, Fig)
	fmt.Println(i, j, k, l)
}
